package common

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func PrefixEnvVar(prefix, suffix string) []string {
	return []string{prefix + "_" + suffix}
}

func ValidateEnvVars(prefix string, flags []cli.Flag, log log.Logger) {
	for _, envVar := range validateEnvVars(prefix, os.Environ(), cliFlagsToEnvVars(flags)) {
		log.Warn("Unknown env var", "prefix", prefix, "env_var", envVar)
	}
}

func cliFlagsToEnvVars(flags []cli.Flag) map[string]struct{} {
	definedEnvVars := make(map[string]struct{})
	for _, flag := range flags {
		envVars := reflect.ValueOf(flag).Elem().FieldByName("EnvVars")
		for i := 0; i < envVars.Len(); i++ {
			envVarField := envVars.Index(i)
			definedEnvVars[envVarField.String()] = struct{}{}
		}
	}
	return definedEnvVars
}

func validateEnvVars(prefix string, providedEnvVars []string, definedEnvVars map[string]struct{}) []string {
	var out []string
	for _, envVar := range providedEnvVars {
		parts := strings.Split(envVar, "=")
		if len(parts) == 0 {
			continue
		}
		key := parts[0]
		if strings.HasPrefix(key, prefix) {
			if _, ok := definedEnvVars[key]; !ok {
				out = append(out, envVar)
			}
		}
	}
	return out
}

func ParseAddress(address string) (common.Address, error) {
	if common.IsHexAddress(address) {
		return common.HexToAddress(address), nil
	}
	return common.Address{}, fmt.Errorf("invalid address: %v", address)
}

func CloseAction(fn func(ctx context.Context, shutdown <-chan struct{}) error) error {
	stopped := make(chan error, 1)
	shutdown := make(chan struct{}, 1)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		stopped <- fn(ctx, shutdown)
	}()

	doneCh := make(chan os.Signal, 1)
	signal.Notify(doneCh, []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}...)

	select {
	case <-doneCh:
		cancel()
		shutdown <- struct{}{}

		select {
		case err := <-stopped:
			return err
		case <-time.After(time.Second * 10):
			return errors.New("command action is unresponsive for more than 10 seconds... shutting down")
		}
	case err := <-stopped:
		cancel()
		return err
	}
}

package retry

import (
	"math"
	"math/rand"
	"time"
)

type Strategy interface {
	Duration(attempt int) time.Duration
}

type ExponentialStrategy struct {
	Min       time.Duration
	Max       time.Duration
	MaxJitter time.Duration
}

func (e *ExponentialStrategy) Duration(attempt int) time.Duration {
	var jitter time.Duration
	if e.MaxJitter > 0 {
		jitter = time.Duration(rand.Int63n(e.MaxJitter.Nanoseconds()))
	}
	if attempt < 0 {
		return e.Min + jitter
	}
	durFloat := float64(e.Min)
	durFloat += math.Pow(2, float64(attempt)) * float64(time.Second)
	dur := time.Duration(durFloat)
	if durFloat > float64(e.Max) {
		dur = e.Max
	}
	dur += jitter

	return dur
}

func Exponential() Strategy {
	return &ExponentialStrategy{
		Min:       0,
		Max:       10 * time.Second,
		MaxJitter: 250 * time.Millisecond,
	}
}

type FixedStrategy struct {
	Dur time.Duration
}

func (f *FixedStrategy) Duration(attempt int) time.Duration {
	return f.Dur
}

func Fixed(dur time.Duration) Strategy {
	return &FixedStrategy{
		Dur: dur,
	}
}

package routes

import (
	"math/big"
	"strconv"

	"errors"

	"github.com/ethereum/go-ethereum/common"
)

type Validator struct{}

func (v *Validator) ParseValidateLimit(limit string) (int, error) {
	if limit == "" {
		return defaultPageLimit, nil
	}
	val, err := strconv.Atoi(limit)
	if err != nil {
		return 0, errors.New("limit must be an integer value")
	}
	if val <= 0 {
		return 0, errors.New("limit must be greater than 0")
	}
	return val, nil
}

func (v *Validator) ParseValidatePage(page string) (int, error) {
	if page == "" {
		return 0, nil
	}
	val, err := strconv.Atoi(page)
	if err != nil {
		return 0, errors.New("page must be an integer value")
	}
	if val < 0 {
		return 0, errors.New("page must be greater than 0")
	}
	return val, nil
}

func (v *Validator) ParseValidateInt(intValue string) (int, error) {
	if intValue == "" {
		return 0, errors.New("lack of necessary params")
	}
	val, err := strconv.Atoi(intValue)
	if err != nil {
		return 0, errors.New("intValue must be an integer value")
	}
	if val < 0 {
		return 0, errors.New("page must be greater than 0")
	}

	return val, nil
}

func (v *Validator) ParseValidateBigInt(bigIntValue string) (*big.Int, error) {
	if bigIntValue == "" {
		return big.NewInt(0), errors.New("lack of necessary params")
	}
	val, err := new(big.Int).SetString(bigIntValue, 10)
	if err != true {
		return big.NewInt(0), errors.New("bigIntValue must be an integer value")
	}

	if val.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0), errors.New("bigIntValue must be greater than 0")
	}

	return val, nil
}

func (v *Validator) ParseValidateAddress(addr string) (common.Address, error) {
	if !common.IsHexAddress(addr) {
		return common.Address{}, errors.New("address must be represented as a valid hexadecimal string")
	}
	parsedAddr := common.HexToAddress(addr)
	if parsedAddr == common.HexToAddress("0x0") {
		return common.Address{}, errors.New("address cannot be the zero address")
	}
	return parsedAddr, nil
}

func (v *Validator) ValidateCursor(cursor string) error {
	if cursor == "" {
		return nil
	}
	if len(cursor) != 66 {
		return errors.New("cursor must be a 32 byte hex string")
	}
	if cursor[:2] != "0x" {
		return errors.New("cursor must begin with 0x")
	}
	return nil
}

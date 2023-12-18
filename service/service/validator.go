package service

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
)

type Validator struct{}

func (v *Validator) ParseValidateAddress(addr string) (common.Address, error) {
	var parsedAddr common.Address
	if addr != "0x00" {
		if !common.IsHexAddress(addr) {
			return common.Address{}, errors.New("address must be represented as a valid hexadecimal string")
		}
		parsedAddr = common.HexToAddress(addr)
		if parsedAddr == common.HexToAddress("0x0") {
			return common.Address{}, errors.New("address cannot be the zero address")
		}
	}
	return parsedAddr, nil
}

func (v *Validator) ValidatePage(page int) error {
	if page <= 0 {
		return errors.New("page must be more than 1")
	}
	return nil
}

func (v *Validator) ValidatePageSize(pageSize int) error {
	if pageSize <= 0 {
		return errors.New("page size must be more than 1")
	}
	if pageSize > 1000 {
		return errors.New("page size must be less than 1000")
	}
	return nil
}

func (v *Validator) ValidateOrder(order string) error {
	if order == "asc" || order == "ASC" || order == "DESC" || order == "desc" {
		return nil
	} else {
		return errors.New("order is must one of asc, ASC, DESC, desc value")
	}
}

func (v *Validator) ValidateIdOrIndex(idOrIndex uint64) error {
	if idOrIndex <= 0 {
		return errors.New("page size must be more than 0")
	}
	return nil
}

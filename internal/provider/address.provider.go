package provider

import "errors"

type IAddressProvider interface {
	ProviderName() string
	SearchByZipCode(zipcode string) (*SearchAddressByZipCodeResult, error)
}

type SearchAddressByZipCodeResult struct {
	Address string `json:"address"`
}

var (
	ErrSearchAddressByZipCode        = errors.New("fail to search address by zip code")
	ErrSearchAddressByZipCodeTimeout = errors.New("timeout")
)

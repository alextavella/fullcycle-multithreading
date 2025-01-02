package usecase

import (
	"time"

	"github.com/alextavella/multithreading/internal/provider"
	"github.com/alextavella/multithreading/internal/repository"
)

type ISearchAddressUsecase interface {
	SearchByZipCode(zipCode string) (*SearchAddressUsecaseOutput, error)
}

type SearchAddressUsecaseOutput struct {
	Provider string
	Result   provider.SearchAddressByZipCodeResult
}

type SearchAddressUseCase struct {
	primaryProvider   provider.IAddressProvider
	secondaryProvider provider.IAddressProvider
}

func NewSearchAddressUsecase() ISearchAddressUsecase {
	return &SearchAddressUseCase{
		primaryProvider:   repository.NewBrasilAPIRepository(),
		secondaryProvider: repository.NewViaCEPRepository(),
	}
}

func (u *SearchAddressUseCase) SearchByZipCode(zipCode string) (*SearchAddressUsecaseOutput, error) {
	primaryChannel := make(chan SearchAddressUsecaseOutput)
	secondaryChannel := make(chan SearchAddressUsecaseOutput)

	// Primary provider
	go func(ch chan SearchAddressUsecaseOutput, zipCode string) {
		res, err := u.primaryProvider.SearchByZipCode(zipCode)
		if err != nil {
			return
		}
		ch <- SearchAddressUsecaseOutput{
			Provider: u.primaryProvider.ProviderName(),
			Result:   *res,
		}
	}(primaryChannel, zipCode)

	// Secondary provider
	go func(ch chan SearchAddressUsecaseOutput, zipCode string) {
		res, err := u.secondaryProvider.SearchByZipCode(zipCode)
		if err != nil {
			return
		}
		ch <- SearchAddressUsecaseOutput{
			Provider: u.secondaryProvider.ProviderName(),
			Result:   *res,
		}
	}(secondaryChannel, zipCode)

	select {
	case primaryResult := <-primaryChannel:
		return &primaryResult, nil
	case secondaryResult := <-secondaryChannel:
		return &secondaryResult, nil
	case <-time.After(time.Second):
		return nil, provider.ErrSearchAddressByZipCodeTimeout
	}
}

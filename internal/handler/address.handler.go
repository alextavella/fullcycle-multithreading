package handler

import (
	"github.com/alextavella/multithreading/internal/provider"
	"github.com/alextavella/multithreading/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	SearchAddressUseCase usecase.ISearchAddressUsecase
}

func NewAddressHandler() provider.IHandler {
	return &AddressHandler{
		SearchAddressUseCase: usecase.NewSearchAddressUsecase(),
	}
}

func (h *AddressHandler) Handle(c *fiber.Ctx) error {
	zipCode := c.Params("zipcode")
	result, err := h.SearchAddressUseCase.SearchByZipCode(zipCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"provider": result.Provider,
		"address":  result.Result.Address,
	})
}

package gateways

import (
	"fmt"
	"practice/domain/entities"
	// "practice/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h HTTPGateway) GetAllArtistsGateway(ctx *fiber.Ctx) error {
	data, err := h.artistService.GetAllArtistsService()

	if err != nil {
		fmt.Println("cannot get artist gateway")
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get artists"})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h HTTPGateway) GetArtistByNameGateway(ctx *fiber.Ctx) error {
	// tokenData, err := middlewares.DecodeJWTToken(ctx)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: err.Error()})
	// }

	// myuserID := tokenData.UserID
	params := ctx.Queries()
	name:= params["name"]

	data, err := h.artistService.GetArtistByNameService(name)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h HTTPGateway) CreateArtistGateway(ctx *fiber.Ctx) error {
	// tokenData, err := middlewares.DecodeJWTToken(ctx)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: err.Error()})
	// }

	// myuserID := tokenData.UserID

	var data entities.ArtistDataFormat

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: err.Error()})
	}

	err := h.artistService.CreateArtistService(data)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h HTTPGateway) UpdateArtistGateway(ctx *fiber.Ctx) error {
	// tokenData, err := middlewares.DecodeJWTToken(ctx)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: err.Error()})
	// }	

	// myuserID := tokenData.UserID

	params := ctx.Queries()
	name := params["name"]

	var data entities.ArtistDataFormat

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseModel{Message: err.Error()})
	}

	err := h.artistService.UpdateArtistByNameService(name, data)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h HTTPGateway) DeleteArtistGateway(ctx *fiber.Ctx) error {
	// tokenData, err := middlewares.DecodeJWTToken(ctx)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: err.Error()})
	// }

	// myuserID := tokenData.UserID

	params := ctx.Queries()
	name := params["name"]

	err := h.artistService.DeleteArtistByNameService(name)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}	

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}
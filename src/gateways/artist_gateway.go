package gateways

import (
	"fmt"
	"practice/domain/entities"

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
	params := ctx.Queries()
	name:= params["name"]

	data, err := h.artistService.GetArtistByNameService(name)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get single artist"})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}
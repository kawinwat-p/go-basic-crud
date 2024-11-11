package gateways

import (
	"practice/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	artistService services.IArtistService
}

func NewHTTPGateway(app *fiber.App, artistService services.IArtistService) {
	gateway := &HTTPGateway{
		artistService: artistService,
	}

	GatewayArtists(*gateway,app)
}
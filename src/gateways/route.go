package gateways

import "github.com/gofiber/fiber/v2"

func GatewayArtists(gateway HTTPGateway, app *fiber.App) {
	apiArtist := app.Group("/api/v1/artist")

	apiArtist.Get("/get_all_artists", gateway.GetAllArtistsGateway)
	apiArtist.Get("/get_artist", gateway.GetArtistByNameGateway)
	apiArtist.Post("/create_artist", gateway.CreateArtistGateway)
	apiArtist.Put("/update_artist", gateway.UpdateArtistGateway)
	apiArtist.Delete("/delete_artist", gateway.DeleteArtistGateway)
}

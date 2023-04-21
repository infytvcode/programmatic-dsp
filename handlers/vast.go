package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/infytvcode/dsp/database"
	"github.com/infytvcode/dsp/utils"
)

func VastHandler(c *fiber.Ctx) error {
	camp := database.GetCampaigns()
	admVast := utils.GenerateADM(camp.GetRandomCampaign())
	return c.XML(admVast)
}

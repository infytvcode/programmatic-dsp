package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/infytvcode/dsp/database"
	"github.com/infytvcode/dsp/utils"
)

func OpenRTBHandler(c *fiber.Ctx) error {
	campaigns := database.GetCampaigns()
	adId := utils.AdID()
	bResp := utils.GenerateBidResponse(adId, campaigns.GetRandomCampaign())
	return c.JSON(bResp)
}

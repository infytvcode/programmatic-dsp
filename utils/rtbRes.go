package utils

import (
	"encoding/xml"
	"fmt"

	"github.com/infytvcode/dsp/models"
	"github.com/mxmCherry/openrtb/v17/openrtb2"
)

const SERVER_BASE_URL string = "https://programmatic-dsp.infytvcode.repl.co"

func GenerateBidResponse(adID string, campaign *models.Campaign) *openrtb2.BidResponse {
	admVast := GenerateADM(campaign)
	admString, err := xml.Marshal(admVast)
	if err != nil {
		admString = []byte(``)
	}
	bResp := openrtb2.BidResponse{
		ID: adID,
		SeatBid: []openrtb2.SeatBid{
			{
				Seat: "infytv",
				Bid: []openrtb2.Bid{{
					ID:      adID,
					Price:   campaign.Floor,
					AdM:     string(admString),
					ImpID:   campaign.Ad.ID,
					NURL:    fmt.Sprintf("%s/tracking?e=nurl", SERVER_BASE_URL),
					BURL:    fmt.Sprintf("%s/tracking?e=burl", SERVER_BASE_URL),
					LURL:    fmt.Sprintf("%s/tracking?e=lurl", SERVER_BASE_URL),
					CID:     campaign.Ad.ID,
					CrID:    campaign.Creative.ID,
					W:       int64(campaign.Creative.W),
					H:       int64(campaign.Creative.H),
					AdID:    campaign.Ad.ID,
					ADomain: []string{"infy.tv"},
					Dur:     int64(campaign.Ad.Duration),
				}},
			},
		},
	}
	return &bResp
}

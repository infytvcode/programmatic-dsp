package utils

import (
	"fmt"
	"time"

	hvast "github.com/haxqer/vast"
	"github.com/infytvcode/dsp/models"
)

func GenerateADM(campaign *models.Campaign) *hvast.VAST {
	v := hvast.VAST{
		Version: "3.0",
		Ads: []hvast.Ad{
			{
				ID:     campaign.Ad.ID,
				AdType: "video",
				InLine: &hvast.InLine{
					AdSystem: &hvast.AdSystem{
						Version: "1.0",
						Name:    campaign.Ad.Title,
					},
					Description: &hvast.CDATAString{CDATA: "VAST 2.0"},
					AdTitle:     hvast.CDATAString{CDATA: campaign.Ad.Title},
					Errors: []hvast.CDATAString{{
						CDATA: fmt.Sprintf("%s/tracking?e=error", SERVER_BASE_URL),
					}},
					Impressions: []hvast.Impression{{
						ID:  campaign.Ad.ID,
						URI: fmt.Sprintf("%s/tracking?e=imp", SERVER_BASE_URL),
					}},
					Pricing: &hvast.Pricing{
						Value:    fmt.Sprintf("%f", campaign.Floor),
						Currency: "USD",
					},
					Creatives: []hvast.Creative{{
						ID: campaign.Ad.ID,
						Linear: &hvast.Linear{
							Duration: hvast.Duration(campaign.Ad.Duration * int(time.Second)),
							VideoClicks: &hvast.VideoClicks{
								ClickThroughs: []hvast.VideoClick{
									{URI: ""},
								},
							},
							MediaFiles: []hvast.MediaFile{{
								Width:    campaign.Creative.W,
								Height:   campaign.Creative.H,
								Delivery: "progressive",
								Type:     "video/mp4",
								URI:      campaign.Creative.MediaFile,
							}},
							TrackingEvents: []hvast.Tracking{
								{
									Event: "start",
									URI:   fmt.Sprintf("%s/tracking?e=start", SERVER_BASE_URL),
								},
								{
									Event: "firstQuartile",
									URI:   fmt.Sprintf("%s/tracking?e=firstQuartile", SERVER_BASE_URL),
								},
								{
									Event: "midpoint",
									URI:   fmt.Sprintf("%s/tracking?e=midpoint", SERVER_BASE_URL),
								},
								{
									Event: "thirdQuartile",
									URI:   fmt.Sprintf("%s/tracking?e=thirdQuartile", SERVER_BASE_URL),
								},
								{
									Event: "complete",
									URI:   fmt.Sprintf("%s/tracking?e=complete", SERVER_BASE_URL),
								},
							},
						},
					}},
				},
			},
		},
		Errors: []hvast.CDATAString{{
			CDATA: fmt.Sprintf("%s/tracking?e=error", SERVER_BASE_URL),
		}},
	}
	return &v
}

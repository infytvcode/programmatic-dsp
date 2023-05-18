package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/medialab-ai/vmap"
)

func VmapHandler(c *fiber.Ctx) error {
	localVmap := vmap.VMAP{
		AdBreaks: []vmap.AdBreak{},
		Version:  "1.0",
	}
	localVmap.AdBreaks = append(localVmap.AdBreaks, getAdBreak("https://programmatic-dsp.infytvcode.repl.co/vast",
		"linear", "pre-roll", "pre-roll-ad", &vmap.Offset{Position: vmap.OffsetPositionStart}))

	localVmap.AdBreaks = append(localVmap.AdBreaks, getAdBreak("https://programmatic-dsp.infytvcode.repl.co/vast",
		"linear", "mid-roll-10", "mid-roll-ad-10", &vmap.Offset{TimeDur: vmap.Duration(10*time.Second)}))
	localVmap.AdBreaks = append(localVmap.AdBreaks, getAdBreak("https://programmatic-dsp.infytvcode.repl.co/vast",
		"linear", "mid-roll-20", "mid-roll-ad-20", &vmap.Offset{TimeDur: vmap.Duration(20*time.Second)}))
	localVmap.AdBreaks = append(localVmap.AdBreaks, getAdBreak("https://programmatic-dsp.infytvcode.repl.co/vast",
		"linear", "mid-roll-30", "mid-roll-ad-30", &vmap.Offset{TimeDur: vmap.Duration(30*time.Second)}))
	return c.XML(localVmap)
}

func getAdBreak(adUrl string, bType string, bId string, adId string, offset *vmap.Offset) vmap.AdBreak {
	return vmap.AdBreak{
		BreakType:  bType,
		TimeOffset: offset,
		BreakID:    bId,
		AdSource:   getAdSource(adUrl, adId, true),
	}
}

func getAdSource(adUrl string, id string, multiple bool) vmap.AdSource {
	return vmap.AdSource{
		FollowRedirects:  true,
		AllowMultipleAds: multiple,
		ID:               id,
		AdTagURI:         getAdTag(adUrl),
	}
}

func getAdTag(adUrl string) vmap.AdTagURI {
	return vmap.AdTagURI{
		TemplateType: "vast3",
		Text:         adUrl,
	}
}

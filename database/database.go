package database

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	datajson "github.com/infytvcode/dsp/database/data_json"
	"github.com/infytvcode/dsp/models"
)

var (
	db        []*models.Customer
	campaigns models.Campaigns
	mu        sync.Mutex
)

// Connect with database
func Connect() {
	db = make([]*models.Customer, 0)
	cams := []models.Campaign{}
	err := json.Unmarshal(datajson.CampaignsData(), &cams)
	if err != nil {
		log.Println(err.Error())
	}
	campaigns.Campaigns = cams
	fmt.Println("Connected with Database")
}

func Insert(user *models.Customer) {
	mu.Lock()
	db = append(db, user)
	mu.Unlock()
}

func Get() []*models.Customer {
	return db
}

func GetCampaigns() models.Campaigns {
	return campaigns
}

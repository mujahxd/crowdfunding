package transaction

import (
	"time"

	"github.com/mujahxd/crowdfunding/campaign"
	"github.com/mujahxd/crowdfunding/user"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

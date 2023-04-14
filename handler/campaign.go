package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mujahxd/crowdfunding/campaign"
	"github.com/mujahxd/crowdfunding/helper"
)

// tankap param
// handler maping ke service
// service tentukan repo yang dipakai

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of campaigns", http.StatusOK, "sucess", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// handler: mapping id yang di url ke struct input => service
	// service : inputan berupa struct input
	// repository : get campaign By ID
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

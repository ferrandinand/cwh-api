package domain

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/errs"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/logger"
)

type GatusResponse []struct {
	Name    string `json:"name"`
	Group   string `json:"group"`
	Key     string `json:"key"`
	Results []struct {
		Status           int    `json:"status"`
		Hostname         string `json:"hostname"`
		Duration         int    `json:"duration"`
		ConditionResults []struct {
			Condition string `json:"condition"`
			Success   bool   `json:"success"`
		} `json:"conditionResults"`
		Success   bool      `json:"success"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"results"`
}

type GatusAPIAdapter struct {
	client *http.Client
	url    string
}

func (c GatusAPIAdapter) GetStatus(projectId string) (StatusResponseList, *errs.AppError) {

	var gatusResponse GatusResponse
	var response StatusResponseList

	// Make the GET request
	httpResponse, err := c.client.Get(c.url + "/api/v1/endpoints/statuses")
	if err != nil {
		logger.Error("Error while creating new status entry: " + err.Error())

	}
	defer httpResponse.Body.Close()

	err = json.NewDecoder(httpResponse.Body).Decode(&gatusResponse)
	if err != nil {
		logger.Error("Error while creating new status entry: " + err.Error())
	}

	statusResponseItems := make([]StatusResponse, 0)
	for _, resp := range gatusResponse {

		// Get only the data
		if resp.Group == projectId {
			parsedName := strings.Split(resp.Name, " ")
			url := parsedName[len(parsedName)-1]
			statustype := parsedName[0]
			var latestResult []LatestResults

			// Get only 2 results from Gatus API
			for i := len(resp.Results) - 1; i >= 0 && len(latestResult) < 2; i-- {
				status := resp.Results[i]
				newLatestResults := LatestResults{
					Status:    status.Status,
					Success:   status.Success,
					Timestamp: status.Timestamp,
				}
				latestResult = append(latestResult, newLatestResults)
			}
			newStatusResponse := StatusResponse{
				Project:           resp.Group,
				URL:               url,
				Statustype:        statustype,
				BadgeURL:          c.url + "/api/v1/endpoints/" + resp.Key + "/health/badge.svg",
				LatestResultsList: latestResult,
			}
			statusResponseItems = append(statusResponseItems, newStatusResponse)
		}

	}

	response.Items = statusResponseItems

	return response, nil
}

package domain

import (
	"net/http"
	"time"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/dto"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/errs"
)

type StatusResponse struct {
	Project           string `json:"project"`
	URL               string `json:"url"`
	Statustype        string
	BadgeURL          string
	LatestResultsList []LatestResults `json:"results"`
}

type LatestResults struct {
	Status    int
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
}

type StatusResponseList struct {
	Items []StatusResponse `json:"items"`
}

type StatusAdapter interface {
	GetStatus(projectId string) (StatusResponseList, *errs.AppError)
}

func (c StatusResponse) ToDto() dto.StatusResponse {
	var dtoLatestResults []dto.LatestResults

	for _, latest := range c.LatestResultsList {
		newDto := dto.LatestResults{
			Status:    latest.Status,
			Success:   latest.Success,
			Timestamp: latest.Timestamp,
		}
		dtoLatestResults = append(dtoLatestResults, newDto)

	}

	return dto.StatusResponse{
		Project:           c.Project,
		URL:               c.URL,
		Statustype:        c.Statustype,
		BadgeURL:          c.BadgeURL,
		LatestResultsList: dtoLatestResults,
	}
}

func NewStatusAdapter(client *http.Client, url string) GatusAPIAdapter {
	return GatusAPIAdapter{
		client: client,
		url:    url,
	}
}

package dto

import "time"

type StatusResponse struct {
	Project           string          `json:"project"`
	URL               string          `json:"url"`
	Statustype        string          `json:"type"`
	BadgeURL          string          `json:"badge_url"`
	LatestResultsList []LatestResults `json:"results"`
}

type LatestResults struct {
	Status    int       `json:"status"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
}

type StatusResponseList struct {
	Items []StatusResponse `json:"items"`
}

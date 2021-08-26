package service

import (
	"fmt"
)

const dbTSLayout = "2006-01-02 15:04:05"

var statuses = map[string]string{
	"disabled": "0",
	"creating": "1",
	"created":  "2",
	"failed":   "3",
	"deleted":  "4",
	"":         "", //No status defined
}

func statusToNumber(status string) (string, error) {

	if statusInMap, ok := statuses[status]; ok {
		return statusInMap, nil
	}
	return "", fmt.Errorf("Status not found")
}

package domain

import (
	"fmt"
)

func commonStatusAsText(status string) (string, error) {

	if statusInMap, ok := statuses[status]; ok {
		return statusInMap, nil
	}
	return "", fmt.Errorf("Status not found")
}

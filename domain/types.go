package domain

// Set pagination size
const pageSize = 20

type JSONField map[string]interface{}

var statuses = map[string]string{
	"0": "disabled",
	"1": "creating",
	"2": "created",
	"3": "failed",
	"4": "deleted",
}

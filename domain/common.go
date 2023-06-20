package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func (jf *JSONField) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &jf)
		return nil
	case string:
		json.Unmarshal([]byte(v), &jf)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (jf *JSONField) Value() (driver.Value, error) {
	return json.Marshal(jf)
}

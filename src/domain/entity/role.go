package entity

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Role struct {
	RoleId   int    `json:"role_id"`
	RoleCode string `json:"role_code"`
	RoleName string `json:"role_name"`
}

func (r *Role) Scan(src interface{}) error {
	b, ok := src.([]uint8)

	if !ok {
		return nil
	}

	_ = json.Unmarshal(b, &r)

	return nil
}

func (r Role) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	role, _ := json.Marshal(r)
	fmt.Fprintf(buf, "%s", role)
	return buf.Bytes(), nil
}

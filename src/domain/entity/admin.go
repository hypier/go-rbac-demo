package entity

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Admin struct {
	//BaseEntity
	AdminId       int
	AdminName     string
	AdminPassword string
	RoleCode      string
	Role          *Role
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

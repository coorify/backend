package types

import "github.com/gofrs/uuid/v5"

type UUIDBody struct {
	ID string `form:"uuid" json:"uuid" binding:"required"`
}

func (u UUIDBody) ToUUID() uuid.UUID {
	return uuid.Must(uuid.FromString(u.ID))
}

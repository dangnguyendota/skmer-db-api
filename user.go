package casino_database

import "github.com/google/uuid"

type SkmerUser struct {
	Id uuid.UUID
	Username string
	DisplayName string
	Avatar string
	Point int64
}

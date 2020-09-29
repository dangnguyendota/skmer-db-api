package casino_database

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID
	Username string
	DisplayName string
	Avatar string
	Point int64
}

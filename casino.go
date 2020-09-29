package casino_database

import "github.com/google/uuid"

type CasinoDatabase interface {
	GetUserById(id uuid.UUID) User
	UpdatePoint(id uuid.UUID, old, new int64) User
}

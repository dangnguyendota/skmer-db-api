package casino_database

import "github.com/google/uuid"

type SkmerDB interface {
	GetUserById(id uuid.UUID) SkmerUser
	UpdateCoins(id uuid.UUID, oldCoins, newCoins int64) SkmerUser
}
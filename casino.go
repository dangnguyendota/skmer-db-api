package skmerdb

import "github.com/google/uuid"

type Game struct {
	Name string
	Code string
}

type City struct {
	Name string
	Code string
	MinJoin int64
	MaxJoin int64
	MinBet int64
	MaxBet int64
}

type SkmerDB interface {
	// Lấy SkmerUser theo id của tài khoản.
	GetUserById(id uuid.UUID) SkmerUser
	// cập point của tài khoản sau đó trả về tài khoản sau khi cập nhật.
	// o là số point cũ và n là số point mới.
	// id là id của tài khoản.
	UpdatePoint(id uuid.UUID, o, n int64) SkmerUser
	// danh sách các game id phù hợp
	GetAvailableGameIds() []Game
	// danh sách các thành phố
	GetAvailableCities() []City
}

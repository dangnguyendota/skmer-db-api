package skmerdb

import "github.com/google/uuid"

type SkmerDB interface {
	// Lấy SkmerUser theo id của tài khoản.
	GetUserById(id uuid.UUID) SkmerUser
	// cập point của tài khoản sau đó trả về tài khoản sau khi cập nhật.
	// o là số point cũ và n là số point mới.
	// id là id của tài khoản.
	UpdatePoint(id uuid.UUID, o, n int64) SkmerUser
}
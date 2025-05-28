package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  Role   `db:"role"`
}

type Role int32

type UserCreate struct {
	Name         string
	Email        string
	Role         Role
	PasswordHash []byte
}

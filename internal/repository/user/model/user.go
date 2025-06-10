package model

import (
	"database/sql"
	"time"
)

// User represents a user entity with ID, Info, CreatedAt, and UpdatedAt fields.
type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// UserInfo represents user info entity with Name, Email, and Role fields.
type UserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  Role   `db:"role"`
}

// Role represents a user role with an integer value.
type Role int32

// UserCreate represents a user creation with Name, Email, Role, and PasswordHash fields.
type UserCreate struct {
	Name         string
	Email        string
	Role         Role
	PasswordHash []byte
}

// UserUpdate represents a user update with Name, Email, and Role fields.
type UserUpdate struct {
	Name  *string
	Email *string
	Role  *Role
}

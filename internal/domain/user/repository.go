package user

import "context"

type Repository interface {
	Save(ctx context.Context, u *User) error
	FindByID(ctx context.Context, id int64) (*User, error)
	Delete(ctx context.Context, id int64) error
}

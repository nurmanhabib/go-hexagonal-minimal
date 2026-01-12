package mongodb

import (
	"context"
	"database/sql"
	"hexagonal-minimal/internal/domain/user"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, u *user.User) error {
	_, err := r.db.Collection("users").InsertOne(ctx, u)
	return err
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	row := r.db.Collection("users").FindOne(ctx, sql.Named("id", id))
	var u user.User
	if err := row.Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Collection("users").DeleteOne(ctx, sql.Named("id", id))
	return err
}

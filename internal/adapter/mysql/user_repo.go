package mysql

import (
	"context"
	"database/sql"
	"hexagonal-minimal/internal/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, u *user.User) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO users(id, name) VALUES(?, ?)",
		u.ID, u.Name,
	)
	return err
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		"SELECT id, name FROM users WHERE id = ?",
		id,
	)

	var u user.User
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(
		ctx,
		"DELETE FROM users WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, errAffected := res.RowsAffected()
	if errAffected != nil {
		return errAffected
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

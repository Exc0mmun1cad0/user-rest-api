package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"test-api-task/internal/entity"
	userrepo "test-api-task/internal/repository/user"

	"github.com/jmoiron/sqlx"

	"github.com/lib/pq"
)

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetUser(userID string) (*entity.User, error) {
	const op = "repository.user.postgres.GetUser"

	var user entity.User
	err := r.db.Get(&user, `SELECT * FROM users WHERE user_id = $1`, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, userrepo.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (r *repo) CreateUser(user *entity.User) (*entity.User, error) {
	const op = "repository.user.postgres.CreateUser"

	var newUser entity.User
	err := r.db.Get(
		&newUser,
		`INSERT INTO users(first_name, last_name, email, age) VALUES ($1, $2, $3, $4) RETURNING *`,
		user.FirstName, user.LastName, user.Email, user.Age,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return nil, fmt.Errorf("%s: %w", op, userrepo.ErrUserAlreadyExists)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &newUser, nil
}

// newUserInfo consists of new fields for user to be changed and old ones
func (r *repo) UpdateUser(userID string, newUserInfo entity.User) (*entity.User, error) {
	const op = "repository.user.postgres.UpdateUser"

	var newUser entity.User
	err := r.db.Get(
		&newUser,
		`
			UPDATE users SET first_name = $1, last_name = $2, email = $3, age = $4
			WHERE user_id = $5
			RETURNING *
		`,
		newUserInfo.FirstName, newUserInfo.LastName, newUserInfo.Email, newUserInfo.Age,
		userID,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return nil, fmt.Errorf("%s: %w", op, userrepo.ErrUserAlreadyExists)
		}

		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, userrepo.ErrUserNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &newUser, nil
}

func (r *repo) DeleteUser(userID string) error {
	const op = "repository.user.postgres.DeleteUser"

	res, err := r.db.Exec(`DELETE FROM users WHERE user_id = $1`, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if num == 0 {
		return fmt.Errorf("%s: %w", op, userrepo.ErrUserNotFound)
	}

	return nil
}

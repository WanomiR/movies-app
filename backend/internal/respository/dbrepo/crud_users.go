package dbrepo

import (
	"backend/internal/models"
	"context"
)

func (db *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cansel := context.WithTimeout(context.Background(), dbTimeout)
	defer cansel()

	query := `SELECT * FROM users WHERE email = $1`

	var user models.User
	err := db.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

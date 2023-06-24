package repository

import (
	"context"
	"database/sql"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/core/domain/user"
	sqluser "github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/infrastructure/mysql/gen/user"
)

var _ user.Repository = (*userRepository)(nil)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) SaveUser(ctx context.Context, user user.User) error {
	queries := sqluser.New(u.db)
	_, err := queries.InsertUser(ctx, sqluser.InsertUserParams{
		ID:   user.ID,
		Name: user.Name,
	})
	return err
}

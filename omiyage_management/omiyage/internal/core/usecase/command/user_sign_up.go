package command

import (
	"context"
	"errors"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/core/domain/user"
)

type UserSignUpUseCase struct {
	userRepo user.Repository
}

func NewUserSignUpUseCase(userRepo user.Repository) *UserSignUpUseCase {
	return &UserSignUpUseCase{userRepo: userRepo}
}

type UserSignUpRequest struct {
	Name string
}

type UserSignUpResponse struct {
	User user.User
}

var (
	UserSignUpErrorInvalidParameters = errors.New("invalid parameters")
)

func (uc *UserSignUpUseCase) Invoke(
	ctx context.Context,
	req UserSignUpRequest,
) (res UserSignUpResponse, err error) {
	u, err := user.SignUp(req.Name)
	if err != nil {
		return res, errors.Join(UserSignUpErrorInvalidParameters, err)
	}

	if err := uc.userRepo.SaveUser(ctx, *u); err != nil {
		return res, err
	}

	res.User = *u

	return res, err
}

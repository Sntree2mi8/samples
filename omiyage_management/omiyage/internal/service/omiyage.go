package service

import (
	"context"
	"errors"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/core/usecase/command"
	omiyagev1 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/omiyage/v1"
	userv1 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/types/user/v1"
)

var _ omiyagev1.OmiyageServiceServer = (*OmiyageService)(nil)

type OmiyageService struct {
	userSignUpUseCase *command.UserSignUpUseCase
}

func NewOmiyageService(
	userSignUpUseCase *command.UserSignUpUseCase,
) *OmiyageService {
	return &OmiyageService{
		userSignUpUseCase: userSignUpUseCase,
	}
}

func (s *OmiyageService) SignUp(ctx context.Context, request *omiyagev1.SignUpRequest) (*omiyagev1.SignUpResponse, error) {
	u, err := s.userSignUpUseCase.Invoke(ctx, command.UserSignUpRequest{Name: request.GetUserName()})
	if err != nil {
		switch {
		case errors.Is(err, command.UserSignUpErrorInvalidParameters):
			return &omiyagev1.SignUpResponse{
				Error: &omiyagev1.SignUpResponse_DomainError{
					Code:        omiyagev1.SignUpResponse_DomainError_DOMAIN_ERROR_CODE_VALIDATION_ERROR,
					Description: "",
				},
				User: nil,
			}, nil
		}
		return nil, err
	}

	return &omiyagev1.SignUpResponse{
		Error: nil,
		User: &userv1.User{
			Id:   u.User.ID,
			Name: u.User.Name,
		},
	}, nil
}

func (s *OmiyageService) AddRecipientGroups(ctx context.Context, request *omiyagev1.AddRecipientGroupsRequest) (*omiyagev1.AddRecipientGroupsResponse, error) {
	//TODO implement me
	panic("implement me")
}

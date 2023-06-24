package service

import (
	"context"
	omiyagev1 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/omiyage/v1"
	userv1 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/types/user/v1"
)

var _ omiyagev1.OmiyageServiceServer = (*OmiyageService)(nil)

type OmiyageService struct {
}

func NewOmiyageService() *OmiyageService {
	return &OmiyageService{}
}

func (s *OmiyageService) SignUp(ctx context.Context, request *omiyagev1.SignUpRequest) (*omiyagev1.SignUpResponse, error) {
	if request.UserName == "" {
		return &omiyagev1.SignUpResponse{
			Error: &omiyagev1.SignUpResponse_DomainError{
				Code:        omiyagev1.SignUpResponse_DomainError_DOMAIN_ERROR_CODE_VALIDATION_ERROR,
				Description: "",
			},
			User: nil,
		}, nil
	}

	return &omiyagev1.SignUpResponse{
		Error: nil,
		User: &userv1.User{
			Id:   "unique",
			Name: request.UserName,
		},
	}, nil
}

func (s *OmiyageService) AddRecipientGroups(ctx context.Context, request *omiyagev1.AddRecipientGroupsRequest) (*omiyagev1.AddRecipientGroupsResponse, error) {
	//TODO implement me
	panic("implement me")
}

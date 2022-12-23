package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Group10CapstoneProject/Golang/internal/members/dto"
	memberRepoMock "github.com/Group10CapstoneProject/Golang/internal/members/repository/mock"
	notifRepoMock "github.com/Group10CapstoneProject/Golang/internal/notifications/repository/mock"
	userRepoMock "github.com/Group10CapstoneProject/Golang/internal/users/repository/mock"
	"github.com/Group10CapstoneProject/Golang/model"
	imagekitMock "github.com/Group10CapstoneProject/Golang/utils/imgkit/mock"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/stretchr/testify/suite"
)

type suiteMemberService struct {
	suite.Suite
	memberRepositoryMock       *memberRepoMock.MemberRepositoryMock
	imagekitServiceMock        *imagekitMock.ImagekitServiceMock
	userRepositoryMock         *userRepoMock.UserRepositoryMock
	notificationRepositoryMock *notifRepoMock.NotificationRepositoryMock
	memberService              MemberService
}

func (s *suiteMemberService) SetupTest() {
	s.memberRepositoryMock = new(memberRepoMock.MemberRepositoryMock)
	s.imagekitServiceMock = new(imagekitMock.ImagekitServiceMock)
	s.userRepositoryMock = new(userRepoMock.UserRepositoryMock)
	s.notificationRepositoryMock = new(notifRepoMock.NotificationRepositoryMock)
	s.memberService = NewMemberService(s.memberRepositoryMock, s.imagekitServiceMock, s.notificationRepositoryMock, s.userRepositoryMock)
}

func (s *suiteMemberService) TearDownTest() {
	s.memberRepositoryMock = nil
	s.imagekitServiceMock = nil
	s.userRepositoryMock = nil
	s.notificationRepositoryMock = nil
	s.memberService = nil
}

func (s *suiteMemberService) TestCreateMember() {
	paymentId := uint(1)
	testCase := []struct {
		Name            string
		Body            dto.MemberStoreRequest
		ExpectedErr     error
		ExpectedRes     uint
		CreateMemberRes *model.Member
		CreateMemberErr error
	}{
		{
			Name: "CreateMemberSuccess",
			Body: dto.MemberStoreRequest{
				UserID:          1,
				MemberTypeID:    1,
				Duration:        1,
				PaymentMethodID: &paymentId,
				Total:           100000,
			},
			ExpectedErr: nil,
			ExpectedRes: 1,
			CreateMemberRes: &model.Member{
				ID: 1,
			},
			CreateMemberErr: nil,
		},
		{
			Name: "CreateMemberError",
			Body: dto.MemberStoreRequest{
				UserID:          1,
				MemberTypeID:    1,
				Duration:        1,
				PaymentMethodID: &paymentId,
				Total:           100000,
			},
			ExpectedErr:     errors.New("error"),
			ExpectedRes:     0,
			CreateMemberRes: nil,
			CreateMemberErr: errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("CreateMember").Return(tc.CreateMemberRes, tc.CreateMemberErr)

			res, err := s.memberService.CreateMember(&tc.Body, context.Background())

			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedRes, res)

			s.TearDownTest()
		})
	}
}

func (s *suiteMemberService) TestCreateMemberFromAdmin() {
	testCase := []struct {
		Name               string
		Body               dto.MemberAdminStoreRequest
		ExpectedErr        error
		ExpectedRes        uint
		CreateMemberRes    *model.Member
		CreateMemberErr    error
		FindUserByEmailRes *model.User
		FindUserByEmailErr error
		UpdateMemberErr    error
	}{
		{
			Name: "CreateMemberFromAdminSuccess",
			Body: dto.MemberAdminStoreRequest{
				Email:        "test@gmail.com",
				MemberTypeID: 1,
				Duration:     1,
				Total:        100000,
			},
			ExpectedErr: nil,
			ExpectedRes: 1,
			CreateMemberRes: &model.Member{
				ID: 1,
			},
			CreateMemberErr: nil,
			FindUserByEmailRes: &model.User{
				ID: 1,
			},
			FindUserByEmailErr: nil,
			UpdateMemberErr:    nil,
		},
		{
			Name: "CreateMemberFromAdminError",
			Body: dto.MemberAdminStoreRequest{
				Email:        "test@gmail.com",
				MemberTypeID: 1,
				Duration:     1,
				Total:        100000,
			},
			ExpectedErr:     errors.New("error"),
			ExpectedRes:     0,
			CreateMemberRes: nil,
			CreateMemberErr: errors.New("error"),
			FindUserByEmailRes: &model.User{
				ID: 1,
			},
			FindUserByEmailErr: nil,
			UpdateMemberErr:    nil,
		},
		{
			Name: "CreateMemberFromAdminInvalidEmail",
			Body: dto.MemberAdminStoreRequest{
				Email:        "test@gmail.com",
				MemberTypeID: 1,
				Duration:     1,
				Total:        100000,
			},
			ExpectedErr:        errors.New("error"),
			ExpectedRes:        0,
			CreateMemberRes:    nil,
			CreateMemberErr:    nil,
			FindUserByEmailRes: nil,
			FindUserByEmailErr: errors.New("error"),
			UpdateMemberErr:    nil,
		},
		{
			Name: "CreateMemberFromAdminErrorUpdateMember",
			Body: dto.MemberAdminStoreRequest{
				Email:        "test@gmail.com",
				MemberTypeID: 1,
				Duration:     1,
				Total:        100000,
			},
			ExpectedErr: errors.New("error"),
			ExpectedRes: 0,
			CreateMemberRes: &model.Member{
				ID: 1,
			},
			CreateMemberErr: nil,
			FindUserByEmailRes: &model.User{
				ID: 1,
			},
			FindUserByEmailErr: nil,
			UpdateMemberErr:    errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("CreateMember").Return(tc.CreateMemberRes, tc.CreateMemberErr)
			s.memberRepositoryMock.On("UpdateMember").Return(tc.UpdateMemberErr)
			s.userRepositoryMock.On("FindUserByEmail").Return(tc.FindUserByEmailRes, tc.FindUserByEmailErr)

			res, err := s.memberService.CreateMemberFromAdmin(&tc.Body, context.Background())

			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedRes, res)

			s.TearDownTest()
		})
	}
}
func (s *suiteMemberService) TestFindMembers() {
	testCase := []struct {
		Name             string
		Page             model.Pagination
		ExpectedErr      error
		ExpectedRes      *dto.MemberResponses
		FindMembersRes   []model.Member
		FindMembersCount int
		FindMembersErr   error
	}{
		{
			Name: "FindMembersSuccess",
			Page: model.Pagination{
				Page:  1,
				Limit: 10,
				Q:     "",
			},
			ExpectedErr: nil,
			ExpectedRes: &dto.MemberResponses{
				Members: dto.MemberResources{
					{
						ID: 1,
					},
				},
				Page:  1,
				Limit: 10,
				Count: 1,
			},
			FindMembersRes: []model.Member{
				{
					ID: 1,
				},
			},
			FindMembersCount: 1,
			FindMembersErr:   nil,
		},
		{
			Name: "FindMembersError",
			Page: model.Pagination{
				Page:  1,
				Limit: 10,
				Q:     "",
			},
			ExpectedErr:      errors.New("error"),
			ExpectedRes:      nil,
			FindMembersRes:   nil,
			FindMembersCount: 0,
			FindMembersErr:   errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("FindMembers").Return(tc.FindMembersRes, tc.FindMembersCount, tc.FindMembersErr)

			res, err := s.memberService.FindMembers(&tc.Page, context.Background())

			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedRes, res)

			s.TearDownTest()
		})
	}
}

func (s *suiteMemberService) TestFindMemberById() {
	t := time.Now()
	ac := true
	paymentId := uint(1)
	testCase := []struct {
		Name              string
		ID                uint
		ExpectedErr       error
		ExpectedRes       *dto.MemberDetailResource
		FindMemberByIdRes *model.Member
		FindMemberByIdErr error
	}{
		{
			Name:        "FindMemberByIdSuccess",
			ID:          1,
			ExpectedErr: nil,
			ExpectedRes: &dto.MemberDetailResource{
				ID: 1,
				User: dto.UserResource{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				ExpiredAt:    t,
				ActivedAt:    t,
				Duration:     1,
				ProofPayment: "test",
				PaymentMethod: dto.PaymentMethodResource{
					ID:   &paymentId,
					Name: "test",
				},
				MemberType: dto.MemberTypeResource{
					ID:                 1,
					Name:               "test",
					AccessOfflineClass: true,
					AccessOnlineClass:  true,
					AccessGym:          true,
					AccessTrainer:      true,
				},
				Total:  100000,
				Status: model.ACTIVE,
			},
			FindMemberByIdRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				ExpiredAt:       t,
				ActivedAt:       t,
				Duration:        1,
				ProofPayment:    "test",
				PaymentMethodID: &paymentId,
				PaymentMethod: model.PaymentMethod{
					ID:   &paymentId,
					Name: "test",
				},
				MemberTypeID: 1,
				MemberType: model.MemberType{
					ID:                 1,
					Name:               "test",
					AccessOfflineClass: &ac,
					AccessOnlineClass:  &ac,
					AccessGym:          &ac,
					AccessTrainer:      &ac,
				},
				Total:  100000,
				Status: model.ACTIVE,
			},
			FindMemberByIdErr: nil,
		},
		{
			Name:              "FindMemberByIdError",
			ID:                1,
			ExpectedErr:       errors.New("error"),
			ExpectedRes:       nil,
			FindMemberByIdRes: nil,
			FindMemberByIdErr: errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("FindMemberById").Return(tc.FindMemberByIdRes, tc.FindMemberByIdErr)

			res, err := s.memberService.FindMemberById(tc.ID, context.Background())

			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedRes, res)

			s.TearDownTest()
		})
	}
}

func (s *suiteMemberService) TestCancelMember() {
	testCase := []struct {
		Name              string
		ID                uint
		UserID            uint
		ExpectedErr       error
		FindMemberByIdRes *model.Member
		FindMemberByIdErr error
		UpdateMemberErr   error
	}{
		{
			Name:        "CancelMemberSuccess",
			ID:          1,
			UserID:      1,
			ExpectedErr: nil,
			FindMemberByIdRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:   1,
					Name: "test",
				},
				Status: model.WAITING,
			},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   nil,
		},
		{
			Name:        "CancelMemberError",
			ID:          1,
			UserID:      1,
			ExpectedErr: errors.New("error"),
			FindMemberByIdRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:   1,
					Name: "test",
				},
				Status: model.WAITING,
			},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   errors.New("error"),
		},
		{
			Name:              "CancelMemberNotFound",
			ID:                1,
			UserID:            1,
			ExpectedErr:       errors.New("error"),
			FindMemberByIdRes: nil,
			FindMemberByIdErr: errors.New("error"),
			UpdateMemberErr:   nil,
		},
		{
			Name:        "CancelMemberErrorPermission",
			ID:          1,
			UserID:      1,
			ExpectedErr: myerrors.ErrPermission,
			FindMemberByIdRes: &model.Member{
				ID:     1,
				UserID: 2,
				User: model.User{
					ID:   2,
					Name: "test",
				},
				Status: model.WAITING,
			},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   nil,
		},
		{
			Name:        "CancelMemberErrorIsCancelled",
			ID:          1,
			UserID:      1,
			ExpectedErr: myerrors.ErrIsCanceled,
			FindMemberByIdRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:   1,
					Name: "test",
				},
				Status: model.CANCEL,
			},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   nil,
		},
		{
			Name:        "CancelMemberErrorCantCancel",
			ID:          1,
			UserID:      1,
			ExpectedErr: myerrors.ErrCantCanceled,
			FindMemberByIdRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:   1,
					Name: "test",
				},
				Status: model.ACTIVE,
			},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   nil,
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("FindMemberById").Return(tc.FindMemberByIdRes, tc.FindMemberByIdErr)
			s.memberRepositoryMock.On("UpdateMember").Return(tc.UpdateMemberErr)

			err := s.memberService.CancelMember(tc.ID, tc.UserID, context.Background())

			s.Equal(tc.ExpectedErr, err)

			s.TearDownTest()
		})
	}
}
func (s *suiteMemberService) TestFindMemberByUser() {
	t := time.Now()
	ac := true
	paymentId := uint(1)
	testCase := []struct {
		Name                string
		ID                  uint
		ExpectedErr         error
		ExpectedRes         *dto.MemberDetailResource
		FindMemberByUserRes *model.Member
		FindMemberByUserErr error
	}{
		{
			Name:        "FindMemberByUserSuccess",
			ID:          1,
			ExpectedErr: nil,
			ExpectedRes: &dto.MemberDetailResource{
				ID: 1,
				User: dto.UserResource{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				ExpiredAt:    t,
				ActivedAt:    t,
				Duration:     1,
				ProofPayment: "test",
				PaymentMethod: dto.PaymentMethodResource{
					ID:   &paymentId,
					Name: "test",
				},
				MemberType: dto.MemberTypeResource{
					ID:                 1,
					Name:               "test",
					AccessOfflineClass: true,
					AccessOnlineClass:  true,
					AccessGym:          true,
					AccessTrainer:      true,
				},
				Total:  100000,
				Status: model.ACTIVE,
			},
			FindMemberByUserRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				ExpiredAt:       t,
				ActivedAt:       t,
				Duration:        1,
				ProofPayment:    "test",
				PaymentMethodID: &paymentId,
				PaymentMethod: model.PaymentMethod{
					ID:   &paymentId,
					Name: "test",
				},
				MemberTypeID: 1,
				MemberType: model.MemberType{
					ID:                 1,
					Name:               "test",
					AccessOfflineClass: &ac,
					AccessOnlineClass:  &ac,
					AccessGym:          &ac,
					AccessTrainer:      &ac,
				},
				Total:  100000,
				Status: model.ACTIVE,
			},
			FindMemberByUserErr: nil,
		},
		{
			Name:                "FindMemberByUserError",
			ID:                  1,
			ExpectedErr:         errors.New("error"),
			ExpectedRes:         nil,
			FindMemberByUserRes: nil,
			FindMemberByUserErr: errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("FindMemberByUser").Return(tc.FindMemberByUserRes, tc.FindMemberByUserErr)

			res, err := s.memberService.FindMemberByUser(tc.ID, context.Background())

			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedRes, res)

			s.TearDownTest()
		})
	}
}
func (s *suiteMemberService) TestUpdateMember() {
	paymentId := uint(1)
	testCase := []struct {
		Name            string
		Body            *dto.MemberUpdateRequest
		ExpectedErr     error
		UpdateMemberErr error
	}{
		{
			Name: "UpdateMemberSuccess",
			Body: &dto.MemberUpdateRequest{
				Duration:        1,
				PaymentMethodID: &paymentId,
				MemberTypeID:    1,
				Total:           100000,
			},
			ExpectedErr:     nil,
			UpdateMemberErr: nil,
		},
		{
			Name: "UpdateMemberError",
			Body: &dto.MemberUpdateRequest{
				Duration:        1,
				PaymentMethodID: &paymentId,
				MemberTypeID:    1,
				Total:           100000,
			},
			ExpectedErr:     errors.New("error"),
			UpdateMemberErr: errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("UpdateMember").Return(tc.UpdateMemberErr)

			err := s.memberService.UpdateMember(tc.Body, context.Background())

			s.Equal(tc.ExpectedErr, err)

			s.TearDownTest()
		})
	}
}
func (s *suiteMemberService) TestSetStatusMember() {
	t := time.Now().Add(1 * time.Hour)
	paymentId := uint(1)
	ac := true
	testCase := []struct {
		Name              string
		Body              *dto.SetStatusMember
		ExpectedErr       error
		FindMemberByIdRes *model.Member
		FindMemberByIdErr error
		UpdateMemberErr   error
	}{
		{
			Name: "SetStatusMemberSuccess",
			Body: &dto.SetStatusMember{
				ID:     1,
				Status: model.ACTIVE,
			},
			ExpectedErr: nil,
			FindMemberByIdRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				ExpiredAt:       t,
				ActivedAt:       t,
				Duration:        1,
				ProofPayment:    "test",
				PaymentMethodID: &paymentId,
				PaymentMethod: model.PaymentMethod{
					ID:   &paymentId,
					Name: "test",
				},
				MemberTypeID: 1,
				MemberType: model.MemberType{
					ID:                 1,
					Name:               "test",
					AccessOfflineClass: &ac,
					AccessOnlineClass:  &ac,
					AccessGym:          &ac,
					AccessTrainer:      &ac,
				},
				Total:  100000,
				Status: model.PENDING,
			},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   nil,
		},
		{
			Name: "SetStatusMemberErrorNorFound",
			Body: &dto.SetStatusMember{
				ID:     1,
				Status: model.ACTIVE,
			},
			ExpectedErr:       errors.New("error"),
			FindMemberByIdRes: nil,
			FindMemberByIdErr: errors.New("error"),
			UpdateMemberErr:   nil,
		},
		{
			Name: "SetStatusMemberError",
			Body: &dto.SetStatusMember{
				ID:     1,
				Status: model.ACTIVE,
			},
			ExpectedErr: errors.New("error"),
			FindMemberByIdRes: &model.Member{ID: 1,
				UserID: 1,
				User: model.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				ExpiredAt:       t,
				ActivedAt:       t,
				Duration:        1,
				ProofPayment:    "test",
				PaymentMethodID: &paymentId,
				PaymentMethod: model.PaymentMethod{
					ID:   &paymentId,
					Name: "test",
				},
				MemberTypeID: 1,
				MemberType: model.MemberType{
					ID:                 1,
					Name:               "test",
					AccessOfflineClass: &ac,
					AccessOnlineClass:  &ac,
					AccessGym:          &ac,
					AccessTrainer:      &ac,
				},
				Total:  100000,
				Status: model.PENDING},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   errors.New("error"),
		},
		{
			Name: "SetStatusMemberErrorExpired",
			Body: &dto.SetStatusMember{
				ID:     1,
				Status: model.ACTIVE,
			},
			ExpectedErr:       myerrors.ErrOrderExpired,
			FindMemberByIdRes: &model.Member{},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   nil,
		},
		{
			Name: "SetStatusMemberErrorIsCanceled",
			Body: &dto.SetStatusMember{
				ID:     1,
				Status: model.ACTIVE,
			},
			ExpectedErr: myerrors.ErrIsCanceled,
			FindMemberByIdRes: &model.Member{
				Status: model.CANCEL,
			},
			FindMemberByIdErr: nil,
			UpdateMemberErr:   nil,
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("FindMemberById").Return(tc.FindMemberByIdRes, tc.FindMemberByIdErr)
			s.memberRepositoryMock.On("UpdateMember").Return(tc.UpdateMemberErr)

			err := s.memberService.SetStatusMember(tc.Body, context.Background())

			s.Equal(tc.ExpectedErr, err)

			s.TearDownTest()
		})
	}
}

func (s *suiteMemberService) TestDeleteMember() {
	testCase := []struct {
		Name            string
		ID              uint
		ExpectedErr     error
		DeleteMemberErr error
	}{
		{
			Name:            "DeleteMemberSuccess",
			ID:              1,
			ExpectedErr:     nil,
			DeleteMemberErr: nil,
		},
		{
			Name:            "DeleteMemberError",
			ID:              1,
			ExpectedErr:     errors.New("error"),
			DeleteMemberErr: errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("DeleteMember").Return(tc.DeleteMemberErr)

			err := s.memberService.DeleteMember(tc.ID, context.Background())

			s.Equal(tc.ExpectedErr, err)

			s.TearDownTest()
		})
	}
}

func (s *suiteMemberService) TestCreateMemberType() {
	ac := true
	testCase := []struct {
		Name                string
		Body                *dto.MemberTypeStoreRequest
		ExpectedErr         error
		CreateMemberTypeErr error
	}{
		{
			Name: "CreateMemberTypeSuccess",
			Body: &dto.MemberTypeStoreRequest{
				Name:               "test",
				AccessOfflineClass: ac,
				AccessOnlineClass:  ac,
				AccessGym:          ac,
				AccessTrainer:      ac,
			},
			ExpectedErr:         nil,
			CreateMemberTypeErr: nil,
		},
		{
			Name: "CreateMemberTypeError",
			Body: &dto.MemberTypeStoreRequest{
				Name:               "test",
				AccessOfflineClass: ac,
				AccessOnlineClass:  ac,
				AccessGym:          ac,
				AccessTrainer:      ac,
			},
			ExpectedErr:         errors.New("error"),
			CreateMemberTypeErr: errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("CreateMemberType").Return(tc.CreateMemberTypeErr)

			err := s.memberService.CreateMemberType(tc.Body, context.Background())

			s.Equal(tc.ExpectedErr, err)

			s.TearDownTest()
		})
	}
}

func (s *suiteMemberService) TestFindMemberTypes() {
	ac := true
	testCase := []struct {
		Name               string
		ExpectedErr        error
		ExpectedRes        *dto.MemberTypeResources
		FindMemberTypesErr error
		FindMemberTyoesRes []model.MemberType
	}{
		{
			Name:        "FindMemberTypesSuccess",
			ExpectedErr: nil,
			ExpectedRes: &dto.MemberTypeResources{
				dto.MemberTypeResource{
					ID:                 1,
					Name:               "test",
					AccessOfflineClass: true,
					AccessOnlineClass:  true,
					AccessGym:          true,
					AccessTrainer:      true,
				},
			},
			FindMemberTypesErr: nil,
			FindMemberTyoesRes: []model.MemberType{
				{
					ID:                 1,
					Name:               "test",
					AccessOfflineClass: &ac,
					AccessOnlineClass:  &ac,
					AccessGym:          &ac,
					AccessTrainer:      &ac,
				},
			},
		},
		{
			Name:               "FindMemberTypesError",
			ExpectedErr:        errors.New("error"),
			ExpectedRes:        nil,
			FindMemberTypesErr: errors.New("error"),
			FindMemberTyoesRes: nil,
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("FindMemberTypes").Return(tc.FindMemberTyoesRes, tc.FindMemberTypesErr)

			res, err := s.memberService.FindMemberTypes(context.Background())

			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedRes, res)

			s.TearDownTest()
		})
	}

}

func (s *suiteMemberService) TestFindMemberTypeById() {
	ac := true
	testCase := []struct {
		Name                  string
		ID                    uint
		ExpectedErr           error
		ExpectedRes           *dto.MemberTypeResource
		FindMemberTypeByIdErr error
		FindMemberTypeByIdRes *model.MemberType
	}{
		{
			Name:        "FindMemberTypeByIdSuccess",
			ID:          1,
			ExpectedErr: nil,
			ExpectedRes: &dto.MemberTypeResource{
				ID:                 1,
				Name:               "test",
				AccessOfflineClass: true,
				AccessOnlineClass:  true,
				AccessGym:          true,
				AccessTrainer:      true,
			},
			FindMemberTypeByIdErr: nil,
			FindMemberTypeByIdRes: &model.MemberType{
				ID:                 1,
				Name:               "test",
				AccessOfflineClass: &ac,
				AccessOnlineClass:  &ac,
				AccessGym:          &ac,
				AccessTrainer:      &ac,
			},
		},
		{
			Name:                  "FindMemberTypeByIdError",
			ID:                    1,
			ExpectedErr:           errors.New("error"),
			ExpectedRes:           nil,
			FindMemberTypeByIdErr: errors.New("error"),
			FindMemberTypeByIdRes: nil,
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("FindMemberTypeById").Return(tc.FindMemberTypeByIdRes, tc.FindMemberTypeByIdErr)

			res, err := s.memberService.FindMemberTypeById(tc.ID, context.Background())

			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedRes, res)

			s.TearDownTest()
		})
	}
}

func (s *suiteMemberService) TestUpdateMemberType() {
	testCase := []struct {
		Name        string
		Body        *dto.MemberTypeUpdateRequest
		ExpectedErr error
		UpdateErr   error
	}{
		{
			Name: "UpdateMemberTypeSuccess",
			Body: &dto.MemberTypeUpdateRequest{
				ID:   1,
				Name: "test",
			},
			ExpectedErr: nil,
			UpdateErr:   nil,
		},
		{
			Name: "UpdateMemberTypeError",
			Body: &dto.MemberTypeUpdateRequest{
				ID:   1,
				Name: "test",
			},
			ExpectedErr: errors.New("error"),
			UpdateErr:   errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("UpdateMemberType").Return(tc.UpdateErr)

			err := s.memberService.UpdateMemberType(tc.Body, context.Background())

			s.Equal(tc.ExpectedErr, err)

			s.TearDownTest()
		})
	}
}

func (s *suiteMemberService) TestDeleteMemberType() {
	testCase := []struct {
		Name        string
		ID          uint
		ExpectedErr error
		DeleteErr   error
	}{
		{
			Name:        "DeleteMemberTypeSuccess",
			ID:          1,
			ExpectedErr: nil,
			DeleteErr:   nil,
		},
		{
			Name:        "DeleteMemberTypeError",
			ID:          1,
			ExpectedErr: errors.New("error"),
			DeleteErr:   errors.New("error"),
		},
	}
	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.SetupTest()

			s.memberRepositoryMock.On("DeleteMemberType").Return(tc.DeleteErr)

			err := s.memberService.DeleteMemberType(tc.ID, context.Background())

			s.Equal(tc.ExpectedErr, err)

			s.TearDownTest()
		})
	}
}

func TestSuiteMemberService(t *testing.T) {
	suite.Run(t, new(suiteMemberService))
}

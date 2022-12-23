package mock

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/stretchr/testify/mock"
)

type MemberRepositoryMock struct {
	mock.Mock
}

// member
func (m *MemberRepositoryMock) CreateMember(body *model.Member, ctx context.Context) (*model.Member, error) {
	args := m.Called()
	return args.Get(0).(*model.Member), args.Error(1)
}

func (m *MemberRepositoryMock) FindMembers(page *model.Pagination, ctx context.Context) ([]model.Member, int, error) {
	args := m.Called()
	return args.Get(0).([]model.Member), args.Get(1).(int), args.Error(2)
}

func (m *MemberRepositoryMock) FindMemberById(id uint, ctx context.Context) (*model.Member, error) {
	args := m.Called()
	return args.Get(0).(*model.Member), args.Error(1)
}

func (m *MemberRepositoryMock) FindMemberByUser(userId uint, ctx context.Context) (*model.Member, error) {
	args := m.Called()
	return args.Get(0).(*model.Member), args.Error(1)
}

func (m *MemberRepositoryMock) ReadMembers(body *model.Member, ctx context.Context) ([]model.Member, error) {
	args := m.Called()
	return args.Get(0).([]model.Member), args.Error(1)
}

func (m *MemberRepositoryMock) UpdateMember(body *model.Member, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MemberRepositoryMock) DeleteMember(body *model.Member, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MemberRepositoryMock) MemberInactive(body model.Member, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

// member type
func (m *MemberRepositoryMock) CreateMemberType(body *model.MemberType, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MemberRepositoryMock) FindMemberTypes(ctx context.Context) ([]model.MemberType, error) {
	args := m.Called()
	return args.Get(0).([]model.MemberType), args.Error(1)
}

func (m *MemberRepositoryMock) FindMemberTypeById(id uint, ctx context.Context) (*model.MemberType, error) {
	args := m.Called()
	return args.Get(0).(*model.MemberType), args.Error(1)
}

func (m *MemberRepositoryMock) UpdateMemberType(body *model.MemberType, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MemberRepositoryMock) DeleteMemberType(body *model.MemberType, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

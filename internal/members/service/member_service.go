package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/members/dto"
	"github.com/Group10CapstoneProject/Golang/model"
)

type MemberService interface {
	// member
	CreateMember(request *dto.MemberStoreRequest, ctx context.Context) (uint, error)
	FindMembers(page *model.Pagination, ctx context.Context) (*dto.MemberResponses, error)
	FindMemberById(id uint, ctx context.Context) (*dto.MemberDetailResource, error)
	CancelMember(id uint, userId uint, ctx context.Context) error
	FindMemberByUser(userId uint, ctx context.Context) (*dto.MemberDetailResource, error)
	UpdateMember(request *dto.MemberUpdateRequest, ctx context.Context) error
	SetStatusMember(request *dto.SetStatusMember, ctx context.Context) error
	MemberPayment(request *model.PaymentRequest, ctx context.Context) error
	DeleteMember(id uint, ctx context.Context) error

	// member type
	CreateMemberType(request *dto.MemberTypeStoreRequest, ctx context.Context) error
	FindMemberTypes(ctx context.Context) (*dto.MemberTypeResources, error)
	FindMemberTypeById(id uint, ctx context.Context) (*dto.MemberTypeResource, error)
	UpdateMemberType(request *dto.MemberTypeUpdateRequest, ctx context.Context) error
	DeleteMemberType(id uint, ctx context.Context) error
}

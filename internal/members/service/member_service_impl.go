package service

import (
	"context"
	"time"

	"github.com/Group10CapstoneProject/Golang/internal/members/dto"
	memberRepo "github.com/Group10CapstoneProject/Golang/internal/members/repository"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/google/uuid"
)

type memberServiceImpl struct {
	memberRepository memberRepo.MemberRepository
}

// CreateMember implements MemberService
func (s *memberServiceImpl) CreateMember(request *dto.MemberStoreRequest, ctx context.Context) error {
	member := request.ToModel()
	member.ExpiredAt = time.Now().Add(24 * time.Hour)
	member.Status = model.PENDING
	err := s.memberRepository.CreateMember(member, ctx)
	return err
}

// CreateMemberType implements MemberService
func (s *memberServiceImpl) CreateMemberType(request *dto.MemberTypeStoreRequest, ctx context.Context) error {
	memberType := request.ToModel()
	err := s.memberRepository.CreateMemberType(memberType, ctx)
	return err
}

// DeleteMember implements MemberService
func (s *memberServiceImpl) DeleteMember(id uint, ctx context.Context) error {
	member := model.Member{
		ID: id,
	}
	err := s.memberRepository.DeleteMember(&member, ctx)
	return err
}

// DeleteMemberType implements MemberService
func (s *memberServiceImpl) DeleteMemberType(id uint, ctx context.Context) error {
	memberType := model.MemberType{
		ID: id,
	}
	err := s.memberRepository.DeleteMemberType(&memberType, ctx)
	return err
}

// FindMemberById implements MemberService
func (s *memberServiceImpl) FindMemberById(id uint, ctx context.Context) (*dto.MemberDetailResource, error) {
	member, err := s.memberRepository.FindMemberById(id, ctx)
	if err != nil {
		return nil, err
	}
	if member.Status != model.INACTIVE && time.Now().After(member.ExpiredAt) {
		member.Status = model.INACTIVE
		body := model.Member{
			ID:     member.ID,
			Status: model.INACTIVE,
		}
		err := s.memberRepository.UpdateMember(&body, ctx)
		if err != nil {
			return nil, err
		}
	}
	var result dto.MemberDetailResource
	result.FromModel(member)
	return &result, nil
}

// FindMemberByUser implements MemberService
func (s *memberServiceImpl) FindMemberByUser(userId uint, ctx context.Context) (*dto.MemberResources, error) {
	member, err := s.memberRepository.FindMemberByUser(userId, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.MemberResources
	result.FromModel(member)
	return &result, nil
}

// FindMemberTypes implements MemberService
func (s *memberServiceImpl) FindMemberTypes(ctx context.Context) (*dto.MemberTypeResources, error) {
	memberTypes, err := s.memberRepository.FindMemberTypes(ctx)
	if err != nil {
		return nil, err
	}
	var result dto.MemberTypeResources
	result.FromModel(memberTypes)
	return &result, nil
}

// FindMemberTypeById implements MemberService
func (s *memberServiceImpl) FindMemberTypeById(id uint, ctx context.Context) (*dto.MemberTypeResource, error) {
	memberType, err := s.memberRepository.FindMemberTypeById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.MemberTypeResource
	result.FromModel(memberType)
	return &result, nil
}

// FindMembers implements MemberService
func (s *memberServiceImpl) FindMembers(page *model.Pagination, ctx context.Context) (*dto.MemberResponses, error) {
	members, count, err := s.memberRepository.FindMembers(page, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.MemberResources
	result.FromModel(members)

	response := dto.MemberResponses{
		Members: result,
		Count:   uint(count),
	}
	return &response, nil
}

// UpdateMember implements MemberService
func (s *memberServiceImpl) UpdateMember(request *dto.MemberUpdateRequest, ctx context.Context) error {
	member := request.ToModel()
	check, err := s.memberRepository.FindMemberById(request.ID, ctx)
	if err != nil {
		return err
	}

	if member.Status == model.ACTIVE && check.Status != model.ACTIVE && check.Status != model.INACTIVE {
		member.ExpiredAt = time.Now().Add(24 * 30 * time.Duration(check.Duration) * time.Hour)
		member.Code = uuid.New()
	} else if member.Status == model.REJECT && check.Status != model.REJECT {
		member.ExpiredAt = time.Now()
	} else if member.ProofPayment != "" {
		member.ExpiredAt = time.Now().Add(24 * time.Hour)
		member.Status = model.WAITING
	}
	if time.Now().After(check.ExpiredAt) {
		member.Status = model.INACTIVE
	}

	err = s.memberRepository.UpdateMember(member, ctx)
	return err
}

// UpdateMemberType implements MemberService
func (s *memberServiceImpl) UpdateMemberType(request *dto.MemberTypeUpdateRequest, ctx context.Context) error {
	memberType := request.ToModel()
	err := s.memberRepository.UpdateMemberType(memberType, ctx)
	return err
}

func NewMemberService(memberRepository memberRepo.MemberRepository) MemberService {
	return &memberServiceImpl{
		memberRepository: memberRepository,
	}
}

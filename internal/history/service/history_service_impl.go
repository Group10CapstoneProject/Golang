package service

import (
	"context"
	"sort"

	"github.com/Group10CapstoneProject/Golang/internal/history/dto"
	memberRepo "github.com/Group10CapstoneProject/Golang/internal/members/repository"
	offlineClassRepo "github.com/Group10CapstoneProject/Golang/internal/offline_classes/repository"
	onlineClassRepo "github.com/Group10CapstoneProject/Golang/internal/online_classes/repository"
	trainerRepo "github.com/Group10CapstoneProject/Golang/internal/trainers/repository"
	"github.com/Group10CapstoneProject/Golang/model"
)

type historyServiceImpl struct {
	memberRepository       memberRepo.MemberRepository
	onlineClassRepository  onlineClassRepo.OnlineClassRepository
	offlineClassRepository offlineClassRepo.OfflineClassRepository
	trainerRepository      trainerRepo.TrainerRepository
}

// FindHistoryActivity implements HistoryService
func (s *historyServiceImpl) FindHistoryActivity(request *dto.HistoryActivityRequest, ctx context.Context) ([]dto.HistoryResource, error) {
	var historyActivity []dto.HistoryResource
	if request.Type == "member" || request.Type == "" {
		cond := &model.Member{
			UserID: request.UserID,
			Status: request.Status,
		}
		members, err := s.memberRepository.ReadMembers(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, member := range members {
			var resource dto.HistoryResource
			resource.FromModelMembers(&member)
			historyActivity = append(historyActivity, resource)
		}
		if request.Status == "WAITING" {
			cond := &model.Member{
				UserID: request.UserID,
				Status: model.PENDING,
			}
			members, err := s.memberRepository.ReadMembers(cond, ctx)
			if err != nil {
				return nil, err
			}
			for _, member := range members {
				var resource dto.HistoryResource
				resource.FromModelMembers(&member)
				historyActivity = append(historyActivity, resource)
			}
		}
	}
	if request.Type == "trainer" || request.Type == "" {
		cond := &model.TrainerBooking{
			UserID: request.UserID,
			Status: request.Status,
		}
		trainers, err := s.trainerRepository.ReadTrainerBooking(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, trainer := range trainers {
			var resource dto.HistoryResource
			resource.FromModelTrainer(&trainer)
			historyActivity = append(historyActivity, resource)
		}
		if request.Status == "WAITING" {
			cond := &model.TrainerBooking{
				UserID: request.UserID,
				Status: model.PENDING,
			}
			trainers, err := s.trainerRepository.ReadTrainerBooking(cond, ctx)
			if err != nil {
				return nil, err
			}
			for _, trainer := range trainers {
				var resource dto.HistoryResource
				resource.FromModelTrainer(&trainer)
				historyActivity = append(historyActivity, resource)
			}
		}
	}
	if request.Type == "online_class" || request.Type == "" {
		cond := &model.OnlineClassBooking{
			UserID: request.UserID,
			Status: request.Status,
		}
		onlineClasses, err := s.onlineClassRepository.ReadOnlineClassBooking(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, onlineClass := range onlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOnlineClass(&onlineClass)
			historyActivity = append(historyActivity, resource)
		}
		if request.Status == "WAITING" {
			cond := &model.OnlineClassBooking{
				UserID: request.UserID,
				Status: model.PENDING,
			}
			onlineClasses, err := s.onlineClassRepository.ReadOnlineClassBooking(cond, ctx)
			if err != nil {
				return nil, err
			}
			for _, onlineClass := range onlineClasses {
				var resource dto.HistoryResource
				resource.FromModelOnlineClass(&onlineClass)
				historyActivity = append(historyActivity, resource)
			}
		}
	}
	if request.Type == "offline_class" || request.Type == "" {
		cond := &model.OfflineClassBooking{
			UserID: request.UserID,
			Status: request.Status,
		}
		offlineClasses, err := s.offlineClassRepository.ReadOfflineClassBookings(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, offlineClass := range offlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOfflineClass(&offlineClass)
			historyActivity = append(historyActivity, resource)
		}
		if request.Status == "WAITING" {
			cond := &model.OfflineClassBooking{
				UserID: request.UserID,
				Status: model.PENDING,
			}
			offlineClasses, err := s.offlineClassRepository.ReadOfflineClassBookings(cond, ctx)
			if err != nil {
				return nil, err
			}
			for _, offlineClass := range offlineClasses {
				var resource dto.HistoryResource
				resource.FromModelOfflineClass(&offlineClass)
				historyActivity = append(historyActivity, resource)
			}
		}
	}
	sort.Slice(historyActivity, func(i, j int) bool {
		return historyActivity[i].CreatedAt.After(historyActivity[j].CreatedAt)
	})
	return historyActivity, nil
}

// FindHistoryOrder implements HistoryService
func (s *historyServiceImpl) FindHistoryOrder(request *dto.HistoryOrderRequest, ctx context.Context) ([]dto.HistoryResource, error) {
	var historyActivity []dto.HistoryResource
	if request.Type == "member" || request.Type == "" {
		cond := &model.Member{
			UserID: request.UserID,
			Status: model.DONE,
		}
		members, err := s.memberRepository.ReadMembers(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, member := range members {
			var resource dto.HistoryResource
			resource.FromModelMembers(&member)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.Member{
			UserID: request.UserID,
			Status: model.CENCEL,
		}
		members, err = s.memberRepository.ReadMembers(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, member := range members {
			var resource dto.HistoryResource
			resource.FromModelMembers(&member)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.Member{
			UserID: request.UserID,
			Status: model.REJECT,
		}
		members, err = s.memberRepository.ReadMembers(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, member := range members {
			var resource dto.HistoryResource
			resource.FromModelMembers(&member)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.Member{
			UserID: request.UserID,
			Status: model.INACTIVE,
		}
		members, err = s.memberRepository.ReadMembers(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, member := range members {
			var resource dto.HistoryResource
			resource.FromModelMembers(&member)
			historyActivity = append(historyActivity, resource)
		}
	}
	if request.Type == "online_class" || request.Type == "" {
		cond := &model.OnlineClassBooking{
			UserID: request.UserID,
			Status: model.DONE,
		}
		onlineClasses, err := s.onlineClassRepository.ReadOnlineClassBooking(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, onlineClass := range onlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOnlineClass(&onlineClass)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.OnlineClassBooking{
			UserID: request.UserID,
			Status: model.REJECT,
		}
		onlineClasses, err = s.onlineClassRepository.ReadOnlineClassBooking(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, onlineClass := range onlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOnlineClass(&onlineClass)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.OnlineClassBooking{
			UserID: request.UserID,
			Status: model.CENCEL,
		}
		onlineClasses, err = s.onlineClassRepository.ReadOnlineClassBooking(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, onlineClass := range onlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOnlineClass(&onlineClass)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.OnlineClassBooking{
			UserID: request.UserID,
			Status: model.INACTIVE,
		}
		onlineClasses, err = s.onlineClassRepository.ReadOnlineClassBooking(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, onlineClass := range onlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOnlineClass(&onlineClass)
			historyActivity = append(historyActivity, resource)
		}
	}
	if request.Type == "offline_class" || request.Type == "" {
		cond := &model.OfflineClassBooking{
			UserID: request.UserID,
			Status: model.DONE,
		}
		offlineClasses, err := s.offlineClassRepository.ReadOfflineClassBookings(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, offlineClass := range offlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOfflineClass(&offlineClass)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.OfflineClassBooking{
			UserID: request.UserID,
			Status: model.CENCEL,
		}
		offlineClasses, err = s.offlineClassRepository.ReadOfflineClassBookings(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, offlineClass := range offlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOfflineClass(&offlineClass)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.OfflineClassBooking{
			UserID: request.UserID,
			Status: model.REJECT,
		}
		offlineClasses, err = s.offlineClassRepository.ReadOfflineClassBookings(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, offlineClass := range offlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOfflineClass(&offlineClass)
			historyActivity = append(historyActivity, resource)
		}
		cond = &model.OfflineClassBooking{
			UserID: request.UserID,
			Status: model.INACTIVE,
		}
		offlineClasses, err = s.offlineClassRepository.ReadOfflineClassBookings(cond, ctx)
		if err != nil {
			return nil, err
		}
		for _, offlineClass := range offlineClasses {
			var resource dto.HistoryResource
			resource.FromModelOfflineClass(&offlineClass)
			historyActivity = append(historyActivity, resource)
		}
	}
	sort.Slice(historyActivity, func(i, j int) bool {
		return historyActivity[i].CreatedAt.After(historyActivity[j].CreatedAt)
	})
	return historyActivity, nil
}

func NewHistoryService(
	memberRepository memberRepo.MemberRepository,
	onlineClassRepository onlineClassRepo.OnlineClassRepository,
	offlineClassRepository offlineClassRepo.OfflineClassRepository,
	trainerRepository trainerRepo.TrainerRepository,
) HistoryService {
	return &historyServiceImpl{
		memberRepository:       memberRepository,
		onlineClassRepository:  onlineClassRepository,
		offlineClassRepository: offlineClassRepository,
		trainerRepository:      trainerRepository,
	}
}

package service

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/Group10CapstoneProject/Golang/internal/history/dto"
	memberRepo "github.com/Group10CapstoneProject/Golang/internal/members/repository"
	offlineClassRepo "github.com/Group10CapstoneProject/Golang/internal/offline_classes/repository"
	onlineClassRepo "github.com/Group10CapstoneProject/Golang/internal/online_classes/repository"
	"github.com/Group10CapstoneProject/Golang/model"
)

type historyServiceImpl struct {
	memberRepository       memberRepo.MemberRepository
	onlineClassRepository  onlineClassRepo.OnlineClassRepository
	offlineClassRepository offlineClassRepo.OfflineClassRepository
}

// FindHistoryActivity implements HistoryService
func (s *historyServiceImpl) FindHistoryActivity(request *dto.HistoryActivityRequest, ctx context.Context) ([]dto.HistoryResource, error) {
	var historyActivity []dto.HistoryResource
	if request.Type == "member" || request.Type == "" {
		cond := &model.Member{
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
	if request.Type == "online_class" || request.Type == "" {
		cond := &model.OnlineClassBooking{
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
		timeI, err := time.Parse("2006-01-02 15:04:05", historyActivity[i].CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		timeJ, err := time.Parse("2006-01-02 15:04:05", historyActivity[j].CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		return timeI.After(timeJ)
	})
	return historyActivity, nil
}

// FindHistoryOrder implements HistoryService
func (s *historyServiceImpl) FindHistoryOrder(request *dto.HistoryOrderRequest, ctx context.Context) ([]dto.HistoryResource, error) {
	var historyActivity []dto.HistoryResource
	if request.Type == "member" || request.Type == "" {
		cond := &model.Member{
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
		timeI, err := time.Parse("2006-01-02 15:04:05", historyActivity[i].CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		timeJ, err := time.Parse("2006-01-02 15:04:05", historyActivity[j].CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		return timeI.After(timeJ)
	})
	return historyActivity, nil
}

func NewHistoryService(
	memberRepository memberRepo.MemberRepository,
	onlineClassRepository onlineClassRepo.OnlineClassRepository,
	offlineClassRepository offlineClassRepo.OfflineClassRepository,
) HistoryService {
	return &historyServiceImpl{
		memberRepository:       memberRepository,
		onlineClassRepository:  onlineClassRepository,
		offlineClassRepository: offlineClassRepository,
	}
}

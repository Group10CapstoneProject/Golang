package dto

type DashboardDto struct {
	TotalMember       int `json:"total_member"`
	TotalMemberType   int `json:"total_memberType"`
	TotalUser         int `json:"total_user"`
	TotalOfflineClass int `json:"total_offline_class"`
	TotalOnlineClass  int `json:"total_online_class"`
	TotalTrainer      int `json:"total_trainer"`
}

package model

import "time"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}
type AdminToken struct {
	Access  Access  `json:"access"`
	Refresh Refresh `json:"refresh"`
}

type Access struct {
	AccessAccessToken string    `json:"token"`
	ExpiredAt         time.Time `json:"expired_at"`
}
type Refresh struct {
	RefreshToken string    `json:"token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

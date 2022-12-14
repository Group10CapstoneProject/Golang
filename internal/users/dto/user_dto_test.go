package dto

import (
	"testing"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/stretchr/testify/assert"
)

func TestNewUser_ToModel(t *testing.T) {
	testCase := []struct {
		Name     string
		Dto      NewUser
		Expected *model.User
	}{
		{
			Name: "all filled",
			Dto: NewUser{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "test",
			},
			Expected: &model.User{
				Name:  "test",
				Email: "test@gmail.com",
			},
		},
		{
			Name: "some filled",
			Dto: NewUser{
				Email: "test@gmail.com",
			},
			Expected: &model.User{
				Email: "test@gmail.com",
			},
		},
		{
			Name:     "empty",
			Dto:      NewUser{},
			Expected: &model.User{},
		},
	}
	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			result := v.Dto.ToModel()
			assert.Equal(t, v.Expected, result)
		})
	}
}

package mock

import "github.com/stretchr/testify/mock"

type ImagekitServiceMock struct {
	mock.Mock
}

func (m *ImagekitServiceMock) Upload(title string, file interface{}) (url string, err error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

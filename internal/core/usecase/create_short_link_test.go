package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gotiny/internal/core/model"
)

type mockLinksRepository struct {
	mock.Mock
}

func (m *mockLinksRepository) GetNextLinkIndex() (uint, error) {
	args := m.Called()
	return args.Get(0).(uint), args.Error(1)
}

func (m *mockLinksRepository) SaveLink(link model.Link) error {
	args := m.Called()
	return args.Error(0)
}

func TestCreateShortLink(t *testing.T) {
	repo := new(mockLinksRepository)
	repo.On("GetNextLinkIndex").Return(uint(1), nil).Once()
	repo.On("SaveLink").Return(nil).Once()

	usecase := NewCreateShortLink(repo)

	config := model.LinkConfig{
		Host:     "localhost:8080",
		Protocol: "http",
	}
	link, err := usecase.Call("https://www.google.com", config)

	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com", link.OriginalLink)
	assert.Equal(t, "http://localhost:8080/1", link.ShortLink)
	assert.Equal(t, "1", link.Id)
	repo.AssertExpectations(t)
}

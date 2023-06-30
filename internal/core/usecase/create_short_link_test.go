package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"

	"gotiny/internal/core/model"
)

type mockUrlChecker struct {
	mock.Mock
}

func (m *mockUrlChecker) CheckUrl(ctx context.Context, url string) (bool, error) {
	args := m.Called(url)
	return args.Bool(0), args.Error(1)
}

type mockCoreConfig struct {
	mock.Mock
}

func (m *mockCoreConfig) BaseUrl() string {
	args := m.Called()
	return args.String(0)
}

type mockLinksRepository struct {
	mock.Mock
}

func (m *mockLinksRepository) GetNextLinkIndex(_ context.Context) (uint, error) {
	args := m.Called()
	return args.Get(0).(uint), args.Error(1)
}

func (m *mockLinksRepository) SaveLink(_ context.Context, link model.Link) error {
	args := m.Called()
	return args.Error(0)
}

func TestCreateShortLink(t *testing.T) {
	repo := new(mockLinksRepository)
	repo.On("GetNextLinkIndex").Return(uint(1), nil).Once()
	repo.On("SaveLink").Return(nil).Once()

	core_config := new(mockCoreConfig)
	core_config.On("BaseUrl").Return("http://localhost:8080").Once()

	url_checker := new(mockUrlChecker)
	url_checker.On("CheckUrl", "https://www.google.com").Return(true, nil).Once()

	usecase := NewCreateShortLink(repo, core_config, url_checker)

	config := model.LinkConfig{
		Host: "localhost:8080",
	}
	link, err := usecase.Call(context.Background(), "https://www.google.com", config)

	assert.Nil(t, err)
	assert.Equal(t, "https://www.google.com", link.OriginalLink)
	assert.Equal(t, "http://localhost:8080/1", link.ShortLink)
	assert.Equal(t, "1", link.Id)
	repo.AssertExpectations(t)
	core_config.AssertExpectations(t)
}

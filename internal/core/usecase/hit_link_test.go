package usecase

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gotiny/internal/core/model"
)

type MockHitLinkRepository struct {
	mock.Mock
	wg sync.WaitGroup
}

func (m *MockHitLinkRepository) GetLinkById(id string) (*model.Link, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Link), args.Error(1)
}

func (m *MockHitLinkRepository) DeleteLinkById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockHitLinkRepository) IncrementHitCount(id string) error {
	args := m.Called(id)
	m.wg.Done()
	return args.Error(0)
}

func TestHitLinkValid(t *testing.T) {
	repository := new(MockHitLinkRepository)
	repository.On("GetLinkById", "id").Return(&model.Link{
		OriginalLink: "original_link",
	}, nil)
	repository.wg.Add(1)
	repository.On("IncrementHitCount", "id").Return(nil).Once()

	hitLink := NewHitLink(repository)
	originalLink, err := hitLink.Call("id")

	assert.Nil(t, err)
	assert.Equal(t, "original_link", *originalLink)
	repository.wg.Wait()
	repository.AssertExpectations(t)
}

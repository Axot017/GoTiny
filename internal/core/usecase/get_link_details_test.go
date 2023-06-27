package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gotiny/internal/core/model"
)

type MockGetLinkDetailsRepository struct {
	mock.Mock
}

func (m *MockGetLinkDetailsRepository) GetLinkById(id string) (*model.Link, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Link), args.Error(1)
}

func TestGetLinkDetailsInvalidToken(t *testing.T) {
	mockRepository := new(MockGetLinkDetailsRepository)
	mockRepository.On("GetLinkById", "id").Return(&model.Link{
		Token: "token",
	}, nil).Once()

	getLinkDetails := NewGetLinkDetails(mockRepository)
	_, err := getLinkDetails.Call("id", "invalid_token")

	assert.Equal(t, model.UnauthorizedError, err.Error())
	mockRepository.AssertExpectations(t)
}

func TestGetLinkDetailsValid(t *testing.T) {
	mockRepository := new(MockGetLinkDetailsRepository)
	link := model.Link{
		Token: "token",
	}
	mockRepository.On("GetLinkById", "id").Return(&link, nil).Once()

	getLinkDetails := NewGetLinkDetails(mockRepository)
	result, _ := getLinkDetails.Call("id", "token")

	assert.Equal(t, &link, result)
	mockRepository.AssertExpectations(t)
}

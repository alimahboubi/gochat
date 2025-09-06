package services_test

import (
	"testing"

	"github.com/alimahboubi/gochat/internal/user-service/domain/services"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/alimahboubi/gochat/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestUserDomainService_IsEmailUnique_WithUniqueEmail_ShouldReturnTrue(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	email, _ := valueobjects.NewEmail("test@example.com")
	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		ExistsByEmail(email).
		Return(false, nil)

	service := services.NewUserDomainService(mockRepo)

	// Act
	unique, err := service.IsEmailUnique(email)

	// Assert
	assert.NoError(t, err)
	assert.True(t, unique)
}

func TestUserDomainService_IsEmailUnique_WithExistEmail_ShouldReturnFalse(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	email, _ := valueobjects.NewEmail("test@example.com")

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		ExistsByEmail(email).
		Return(true, nil)
	service := services.NewUserDomainService(mockRepo)

	// Act
	unique, err := service.IsEmailUnique(email)

	// Assert
	assert.NoError(t, err)
	assert.False(t, unique)
}

func TestUserDomainService_IsEmailUnique_WithRepositoryError_ShouldReturnError(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("test@ecample.com")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		ExistsByEmail(email).
		Return(false, errors.New("database connection failed"))

	service := services.NewUserDomainService(mockRepo)
	// Act
	unique, err := service.IsEmailUnique(email)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database connection failed")
	assert.False(t, unique)
}

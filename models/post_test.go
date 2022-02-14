package models

import (
	"testing"

	"github.com/ardhihdra/go-clean-arch/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) FindByID(id int64) (entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.Post), args.Error(1)
}
func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var ids int64 = 1
	post := entity.Post{ID: ids, Title: "A", Text: "B"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostModels(mockRepo)
	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, ids, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{Title: "A", Text: "B"}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostModels(mockRepo)
	result, _ := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testModels := NewPostModels(nil)
	err := testModels.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "post cannot be empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "9"}
	testService := NewPostModels(nil)

	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "post title cannot be empty", err.Error())
}

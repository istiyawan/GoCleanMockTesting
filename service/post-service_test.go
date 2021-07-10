package service

import (
	"cleango/entity"
	"testing"

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

func TestSave(t *testing.T) {
	mockRepo := new(MockRepository)
	//var identifier int64 = 1
	//post := entity.Post{ ID:identifier, Title:"save", Text:"test save"}
	//tes2
	post := entity.Post{Title: "save", Text: "test save"}
	mockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)
	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "save", result.Title)
	assert.Equal(t, "test save", result.Text)
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int64 = 1

	post := entity.Post{ID: identifier, Title: "A", Text: "B"}

	//setup expectation
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	//mock Assertion
	mockRepo.AssertExpectations(t)

	//data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "The post is empty")

}

func TestValidateEmptyTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "B"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "The post title is empty")
}

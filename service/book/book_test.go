package book_test

// import (
// 	"errors"
// 	"testing"

// 	"github.com/dimasyudhana/alterra-group-project-2/entities"
// 	"github.com/dimasyudhana/alterra-group-project-2/service/book"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockRepo struct {
// 	mock.Mock
// }

// func (m *MockRepo) InsertBook(book entities.Core) (entities.Core, error) {
// 	args := m.Called(book)
// 	return args.Get(0).(entities.Core), args.Error(1)
// }

// func TestInsertBook_Success(t *testing.T) {
// 	mockRepo := new(MockRepo)
// 	mockRepo.On("InsertBook", mock.Anything).Return(entities.Core{
// 		ID:       uint(1),
// 		Title:    "Test Book",
// 		Year:     "2022",
// 		Author:   "Test Author",
// 		Contents: "Test Contents",
// 		Image:    "Test Image",
// 		User_ID:  uint(1),
// 	}, nil)

// 	bookService := book.New(mockRepo)
// 	bookInput := entities.Core{
// 		ID:       uint(1),
// 		Title:    "Test Book",
// 		Year:     "2022",
// 		Author:   "Test Author",
// 		Contents: "Test Contents",
// 		Image:    "Test Image",
// 	}

// 	result, err := bookService.InsertBook(bookInput)

// 	assert.NoError(t, err)
// 	assert.Equal(t, uint(1), result.ID)
// 	assert.Equal(t, "Test Book", result.Title)
// 	assert.Equal(t, "2022", result.Year)
// 	assert.Equal(t, "Test Author", result.Author)
// 	assert.Equal(t, "Test Contents", result.Contents)
// 	assert.Equal(t, "Test Image", result.Image)
// 	assert.Equal(t, uint(1), result.User_ID)
// 	mockRepo.AssertExpectations(t)
// }

// func TestInsertBook_Failure(t *testing.T) {
// 	mockRepo := new(MockRepo)
// 	mockRepo.On("InsertBook", mock.Anything).Return(entities.Core{}, errors.New("unexpected error"))

// 	bookService := book.New(mockRepo)
// 	bookInput := entities.Core{
// 		Title:    "Test Book",
// 		Year:     "2022",
// 		Author:   "Test Author",
// 		Contents: "Test Contents",
// 		Image:    "Test Image",
// 	}

// 	result, err := bookService.InsertBook(bookInput)

// 	assert.Error(t, err)
// 	assert.Equal(t, entities.Core{}, result)
// 	mockRepo.AssertExpectations(t)
// }

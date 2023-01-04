package data

import (
	"database/sql"
	"errors"
	"math"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
		GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error)
	}
	Users  UserModel
	Tokens TokenModel
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db},
		Tokens: TokenModel{DB: db},
	}
}

// mock models for testing
type MockMovieModel struct{}

func (m MockMovieModel) Insert(movie *Movie) error {
	// mock action
	return nil
}

func (m MockMovieModel) Get(id int64) (*Movie, error) {
	// mock action
	return nil, nil
}

func (m MockMovieModel) Update(movie *Movie) error {
	// mock action
	return nil
}

func (m MockMovieModel) Delete(id int64) error {
	// mock action
	return nil
}

func (m MockMovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error) {
	// mock action
	return nil, Metadata{}, nil
}

// call NewMockModels() whenever you need it in your unit tests in place of the ‘real’ NewModels() function.
func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

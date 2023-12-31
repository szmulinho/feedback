package endpoints

import (
	"gorm.io/gorm"
	"net/http"
)

type Handlers interface {
	AddOpinion(w http.ResponseWriter, r *http.Request)
	GetAllOpinions(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) Handlers {
	return &handlers{
		db: db,
	}
}

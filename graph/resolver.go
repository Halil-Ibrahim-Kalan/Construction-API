package graph

import (
	"Construction-API/graph/model"
	"sync"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB        *gorm.DB
	Subs      map[string]chan []*model.Message
	SubsMutex sync.Mutex
}

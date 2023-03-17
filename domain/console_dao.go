package domain

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB  *gorm.DB
	Log *zap.Logger
}

type IStorage interface {
	Console()
}

func NewStorage(DB *gorm.DB, Log *zap.Logger) IStorage {
	return &DBInstance{DB: DB, Log: Log}
}

func (DB *DBInstance) Console() {

}

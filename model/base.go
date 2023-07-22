package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"github.com/Makinde1034/budget-app/config"
	
)


type Base struct {
	ID        uuid.UUID  `gorm:"type:char(36);primary_key"`         
	DeletedAt *time.Time `gorm:"index;default:null"`
}

// This functions are called before creating Base
func (base *Base) BeforeCreate(scope *gorm.DB) error { 
	base.ID = uuid.NewV4()
	return nil
}

func init(){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Activity{})
	db.AutoMigrate(&Budget{})
	db.AutoMigrate(&User{})
}
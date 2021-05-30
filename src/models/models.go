package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)


type BaseModel struct {
	gorm.Model
	ID        	uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}

// BeforeSave Pre-save hook to set ID automatically
func (baseModel *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	uuidValue := uuid.NewV4()

	tx.Statement.SetColumn("ID", uuidValue)

	err = nil
	return
}

type EndPoint struct{
	BaseModel
	Url			string	`gorm:"not null"`
	KeyWord		string	`gorm:"size:32;not null"`
	ValidUntil 	time.Time
}

func (endPoint *EndPoint) BeforeCreate(tx *gorm.DB) (err error) {
	randomKeyword := RandStringBytes(8)
	tx.Statement.SetColumn("KeyWord", randomKeyword)

	err = nil
	return
}


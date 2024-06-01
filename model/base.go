package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID uuid.UUID `gorm:"primarykey;size:36;type=varchar"`
}

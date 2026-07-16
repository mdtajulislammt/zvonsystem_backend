package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationCode struct {
	BaseModel
	UserID    *uuid.UUID `gorm:"type:uuid;column:user_id;index" json:"user_id"`
	Token     *string    `gorm:"type:text;column:token" json:"token"`
	Email     *string    `gorm:"type:text;column:email" json:"email"`
	ExpiredAt *time.Time `gorm:"type:timestamp;column:expired_at" json:"expired_at"`
	Status    int        `gorm:"type:smallint;column:status;not null;default:1" json:"status"`
}

func (VerificationCode) TableName() string { return "ucodes" }

func (vc *VerificationCode) BeforeCreate(tx *gorm.DB) (err error) {
	vc.ID = uuid.New()
	return
}

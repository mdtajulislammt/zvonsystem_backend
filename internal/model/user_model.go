package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel // Embedded ID, CreatedAt, UpdatedAt, DeletedAt

	// Text Fields (Using Pointers for Nullable DB Fields)
	Name            *string    `gorm:"type:varchar(255);column:name" json:"name"`
	FirstName       *string    `gorm:"type:varchar(255);column:first_name" json:"first_name"`
	LastName        *string    `gorm:"type:varchar(255);column:last_name" json:"last_name"`
	Username        *string    `gorm:"type:varchar(100);column:username;uniqueIndex" json:"username"`
	Email           *string    `gorm:"type:varchar(255);column:email;uniqueIndex" json:"email"`
	PhoneNumber     *string    `gorm:"type:varchar(50);column:phone_number" json:"phone_number"`
	Password        *string    `gorm:"type:text;column:password" json:"-"` // Omitted from API responses
	TwoFactorSecret *string    `gorm:"type:text;column:two_factor_secret" json:"-"`
	Avatar          *string    `gorm:"type:text;column:avatar" json:"avatar"`
	Gender          *string    `gorm:"type:varchar(20);column:gender" json:"gender"`
	Availability    *string    `gorm:"type:varchar(100);column:availability" json:"availability"`
	BillingID       *string    `gorm:"type:varchar(255);column:billing_id" json:"billing_id"`

	// Address & Region Info
	Country            string  `gorm:"type:varchar(100);column:country_region;not null;default:'Bangladesh'" json:"country_region"` // Kept required non-pointer from previous target state
	UserCountry        *string `gorm:"type:varchar(100);column:country" json:"country"` // Merged from your input
	State              *string `gorm:"type:varchar(100);column:state" json:"state"`
	City               *string `gorm:"type:varchar(100);column:city" json:"city"`
	Address            *string `gorm:"type:text;column:address" json:"address"`
	ZipCode            *string `gorm:"type:varchar(20);column:zip_code" json:"zip_code"`
	LanguagePreference string  `gorm:"type:varchar(50);column:language_preference;not null;default:'English'" json:"language_preference"`

	// DateTime Attributes (Nullable Timestamps)
	DateOfBirth     *time.Time `gorm:"type:date;column:date_of_birth" json:"date_of_birth"`
	EmailVerifiedAt *time.Time `gorm:"type:timestamp;column:email_verified_at" json:"email_verified_at"`
	ApprovedAt      *time.Time `gorm:"type:timestamp;column:approved_at" json:"approved_at"`

	// Core Control & Flag Fields (Non-Pointer, Managed by System Defaults)
	Type               string `gorm:"type:varchar(50);column:type;not null;default:'user'" json:"type"`
	IsTwoFactorEnabled int    `gorm:"type:integer;column:is_two_factor_enabled;not null;default:0" json:"is_two_factor_enabled"`
	Status             int    `gorm:"type:integer;column:status;not null;default:1" json:"status"`
	IsVerified         bool   `gorm:"type:boolean;column:is_verified;not null;default:false" json:"is_verified"` // Architecture core tracker

	// --- Domain Relations Matrix ---
	Releases          []Release          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"releases,omitempty"`
	VerificationCodes []VerificationCode `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"verification_codes,omitempty"`
	Subscription      *UserSubscription  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"subscription,omitempty"`
	Transactions      []Transaction      `gorm:"foreignKey:UserID" json:"transactions,omitempty"`
}

// TableName explicitly overrides table naming to prevent plural mismatch bugs
func (User) TableName() string {
	return "users"
}

// BeforeCreate GORM Hook ensures safe Application-Level v4 UUID injection
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
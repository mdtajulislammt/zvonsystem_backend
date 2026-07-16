package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionPlan struct {
	BaseModel
	Name         string  `gorm:"type:varchar(50);column:name;uniqueIndex;not null" json:"name"` // Free, Standard, Fly High
	Price        float64 `gorm:"type:numeric(10,2);column:price;not null;default:0.00" json:"price"`
	BillingCycle string  `gorm:"type:varchar(20);column:billing_cycle;not null;default:'monthly'" json:"billing_cycle"`
	FeaturesList string  `gorm:"type:jsonb;column:features_list;not null" json:"features_list"` // PostgreSQL JSONB performance optimization
}

func (SubscriptionPlan) TableName() string { return "subscription_plans" }

func (sp *SubscriptionPlan) BeforeCreate(tx *gorm.DB) (err error) {
	sp.ID = uuid.New()
	return
}

type UserSubscription struct {
	BaseModel
	UserID             uuid.UUID         `gorm:"type:uuid;column:user_id;uniqueIndex;not null" json:"user_id"`
	PlanID             uuid.UUID         `gorm:"type:uuid;column:plan_id;not null" json:"plan_id"`
	Status             string            `gorm:"type:varchar(50);column:status;not null;default:'active'" json:"status"` // active, canceled, past_due
	CurrentPeriodStart time.Time         `gorm:"type:timestamp;column:current_period_start;not null" json:"current_period_start"`
	CurrentPeriodEnd   time.Time         `gorm:"type:timestamp;column:current_period_end;not null" json:"current_period_end"`

	Plan *SubscriptionPlan `gorm:"foreignKey:PlanID" json:"plan,omitempty"`
}

func (UserSubscription) TableName() string { return "user_subscriptions" }

func (us *UserSubscription) BeforeCreate(tx *gorm.DB) (err error) {
	us.ID = uuid.New()
	return
}

type Transaction struct {
	BaseModel
	UserID        uuid.UUID `gorm:"type:uuid;column:user_id;index;not null" json:"user_id"`
	Amount        float64   `gorm:"type:numeric(10,2);column:amount;not null" json:"amount"`
	PaymentMethod string    `gorm:"type:varchar(50);column:payment_method;not null" json:"payment_method"` // Visa, Mastercard
	Status        string    `gorm:"type:varchar(20);column:status;not null" json:"status"`                 // success, failed
}

func (Transaction) TableName() string { return "transactions" }

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

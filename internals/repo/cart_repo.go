package repo

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type ICartRepo interface {
	CreateCart(ctx context.Context, userID uint) (uint, error)
	GetCartByUserID(userID uint) (*Cart, error)
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint `gorm:"index;not null"`
	ProductID uint `gorm:"index;not null"`
	Quantity  int  `gorm:"not null;default:1"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Cart struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"index;not null"`
	Items     []CartItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type cartRepo struct {
	db *gorm.DB
}

func CreateRepo(db *gorm.DB) ICartRepo {
	return &cartRepo{db: db}
}

func (c *cartRepo) CreateCart(ctx context.Context, userID uint) (uint, error) {
	cart := Cart{UserID: userID}

	if err := c.db.WithContext(ctx).Create(&cart).Error; err != nil {
		return 0, nil
	}
	return cart.ID, nil
}

func (c *cartRepo) GetCartByUserID(userID uint) (*Cart, error) {
	return &Cart{}, nil
}

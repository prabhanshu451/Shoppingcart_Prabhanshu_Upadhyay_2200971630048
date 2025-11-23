package main

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	Token     string    `gorm:"index" json:"token"`
	CartID    *uint     `json:"cart_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Item struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"index" json:"user_id"`
	Name      string     `json:"name"`
	Status    string     `json:"status"` // e.g., "active" or "ordered"
	CreatedAt time.Time  `json:"created_at"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	CartID uint `gorm:"index" json:"cart_id"`
	ItemID uint `json:"item_id"`
}

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CartID    uint      `json:"cart_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

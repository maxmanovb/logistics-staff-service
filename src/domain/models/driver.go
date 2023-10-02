package models

import "time"

type Driver struct {
    ID           string `json:"id" gorm:"primaryKey"`
    PhoneNumber  string `json:"phone_number" gorm:"index:idx_driver_by_phone"`
    PasswordHash string `json:"password_hash"`
    PasswordSalt string `json:"password_salt"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type DriverProfile struct {
    DriverID     string `json:"driver_id" gorm:"primaryKey"`
    Name         string `json:"name"`
    PhoneNumber  string `json:"phone_number"`
    ImageUrl     string `json:"image_url"`
    Active       bool   `json:"active"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type DriverSmsCode struct {
    ID         int       `json:"id" gorm:"primaryKey"`
    DriverID   string    `json:"driver_id" gorm:"index:idx_sms_code_by_driver"`
    Code       string    `json:"code"`
    ExpiresIn  int       `json:"expires_in"`
    CreatedAt  time.Time `json:"created_at"`
} 
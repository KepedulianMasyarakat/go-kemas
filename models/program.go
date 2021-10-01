package models

import (
    "time"
)

type Event_instansi struct{
    
}

type Program struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    AssingedTo string    `json:"assignedTo"`
    Task       string    `json:"task"`
    Deadline   time.Time `json:"deadline"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Transaksi struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    AssingedTo string    `json:"assignedTo"`
    Task       string    `json:"task"`
    Deadline   time.Time `json:"deadline"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type User struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    AssingedTo string    `json:"assignedTo"`
    Task       string    `json:"task"`
    Deadline   time.Time `json:"deadline"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Admin struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    AssingedTo string    `json:"assignedTo"`
    Task       string    `json:"task"`
    Deadline   time.Time `json:"deadline"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Task struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    AssingedTo string    `json:"assignedTo"`
    Task       string    `json:"task"`
    Deadline   time.Time `json:"deadline"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
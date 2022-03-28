package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	UserId    int       `gorm:"unique_index;not null;"`
	Username  string    `gorm:";not null"`
	CreatedAt time.Time `gorm:";not null"`
	UpdatedAt time.Time `gorm:";not null"`
}

func main() {
	db := createDb()
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	db.Create(&User{UserId: 11, Username: "Alex"})
	db.Create(&User{UserId: 22, Username: "Bob"})
	db.Create(&User{UserId: 33, Username: "Jimmy"})

	u := User{}
	_ = db.Where("user_id = ?", 11).First(&u)
	fmt.Println(u.ID, u.UserId, u.Username, u.CreatedAt, u.UpdatedAt)

	var users []User
	_ = db.Where("user_id >= ?", 22).Find(&users)
	for _, u := range users {
		fmt.Println(u.ID, u.UserId, u.Username, u.CreatedAt, u.UpdatedAt)
	}
}

func createDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:memdb?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

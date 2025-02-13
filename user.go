package main

import "github.com/jinzhu/gorm"

//определяем Userструктуру с полями для имени пользователя, электронной почты и пароля.
// gorm.ModelПоле предоставляет общие поля, такие как ID, CreatedAtи , UpdatedAtкоторые полезны для управления базой данных
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string
}

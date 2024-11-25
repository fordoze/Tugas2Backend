package database

import (
	"fmt"
	"perpus/app/entity"
	"perpus/app/tools"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connet() {

	dsn := "root@tcp(127.0.0.1:3306)/perpus_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("Database Connected")

	//Auto Mitgrate Dari Entity yang di buat

	db.AutoMigrate(
		&entity.SuperAdmin{},
		&entity.Perpustakaan{},
		&entity.Admin{},
		&entity.Buku{},
		&entity.Users{},
		&entity.UsersBook{},
	)

	//Kita Buat Sebuah Create

	User := entity.SuperAdmin{
		Name:     "Super Admin",
		Email:    "SuperAdmin",
		Password: tools.GenaratePassword("123"),
	}

	//Untuk Membuat data di database
	//Cek Kalau Sudah Ada Data Si Super Admin

	var SuperAdmin entity.SuperAdmin
	db.Where("id_super_admin =?", 1).First(&SuperAdmin)

	if SuperAdmin.IdSuperAdmin == 0 {
		//Buat Data SI SuperAdmin
		db.Create(&User)
	}
	fmt.Println("Data Super Admin Created")

	//

	DB = db

}

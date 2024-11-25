package controllers

import (
	"perpus/app/database"
	"perpus/app/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UsersCreate(c *fiber.Ctx) error {
	var request entity.Users

	// Parsing request body untuk data pengguna
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Membuat entri pengguna baru
	newUser := entity.Users{
		Name:      request.Name,
		Jurusan:   request.Jurusan,
		Fakultas:  request.Fakultas,
		Kampus:    request.Kampus,
		Email:     request.Email,
		Password:  request.Password, // Pastikan password di-hash di dalam aplikasi nyata
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Simpan data pengguna ke database
	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User successfully created",
		"data":    newUser,
	})
}

// UsersGetAll - Endpoint untuk mendapatkan semua data pengguna
func UsersGetAll(c *fiber.Ctx) error {
	var dataUsers []entity.Users

	// Mengambil query parameter untuk filter (misal id_users)
	id := c.Query("id")
	baseQuery := database.DB

	if id != "" {
		baseQuery = baseQuery.Where("id_users = ?", id)
	}

	// Query untuk mengambil data pengguna
	if err := baseQuery.Find(&dataUsers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Retrieved all users data",
		"data":    dataUsers,
	})
}

// UsersUpdate - Endpoint untuk memperbarui data pengguna berdasarkan ID
func UsersUpdate(c *fiber.Ctx) error {
	id := c.Query("id")
	var dataUser entity.Users

	// Mencari data pengguna berdasarkan id
	if err := database.DB.Where("id_users =?", id).First(&dataUser).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Parsing request body untuk pembaruan
	var request entity.Users
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update data pengguna
	updates := entity.Users{
		Name:      request.Name,
		Jurusan:   request.Jurusan,
		Fakultas:  request.Fakultas,
		Kampus:    request.Kampus,
		Email:     request.Email,
		Password:  request.Password, // Harus di-hash jika menyimpan password
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Melakukan update ke database
	if err := database.DB.Model(&dataUser).Updates(&updates).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User data updated successfully",
		"data":    dataUser,
	})
}

// UsersDelete - Endpoint untuk menghapus data pengguna berdasarkan ID
func UsersDelete(c *fiber.Ctx) error {
	id := c.Query("id")
	var dataUser entity.Users

	// Mencari data pengguna berdasarkan id
	if err := database.DB.Where("id_users =?", id).First(&dataUser).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Menghapus data pengguna dari database
	if err := database.DB.Delete(&dataUser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User data deleted successfully",
	})
}

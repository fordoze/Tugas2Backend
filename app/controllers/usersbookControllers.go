package controllers

import (
	"perpus/app/database"
	"perpus/app/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UsersBookCreate(c *fiber.Ctx) error {
	var request entity.UsersBook

	// Parsing request body untuk data peminjaman buku
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Membuat entri peminjaman buku baru
	newUsersBook := entity.UsersBook{
		IdBuku:        request.IdBuku,
		IdUsers:       request.IdUsers,
		Durasi:        request.Durasi,
		TanggalPinjam: request.TanggalPinjam,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	// Simpan data peminjaman buku ke database
	if err := database.DB.Create(&newUsersBook).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "UsersBook successfully created",
		"data":    newUsersBook,
	})
}

// UsersBookGetAll - Endpoint untuk mendapatkan semua data peminjaman buku
func UsersBookGetAll(c *fiber.Ctx) error {
	var dataUsersBook []entity.UsersBook

	// Mengambil query parameter untuk filter (misal id_users_book)
	id := c.Query("id")
	baseQuery := database.DB

	if id != "" {
		baseQuery = baseQuery.Where("id_users_book = ?", id)
	}

	// Query untuk mengambil data peminjaman buku
	if err := baseQuery.Find(&dataUsersBook).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Retrieved all UsersBook data",
		"data":    dataUsersBook,
	})
}

// UsersBookUpdate - Endpoint untuk memperbarui data peminjaman buku berdasarkan ID
func UsersBookUpdate(c *fiber.Ctx) error {
	id := c.Query("id")
	var dataUsersBook entity.UsersBook

	// Mencari data peminjaman buku berdasarkan id
	if err := database.DB.Where("id_users_book =?", id).First(&dataUsersBook).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "UsersBook not found",
		})
	}

	// Parsing request body untuk pembaruan
	var request entity.UsersBook
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update data peminjaman buku
	updates := entity.UsersBook{
		IdBuku:        request.IdBuku,
		IdUsers:       request.IdUsers,
		Durasi:        request.Durasi,
		TanggalPinjam: request.TanggalPinjam,
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	// Melakukan update ke database
	if err := database.DB.Model(&dataUsersBook).Updates(&updates).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "UsersBook data updated successfully",
		"data":    dataUsersBook,
	})
}

// UsersBookDelete - Endpoint untuk menghapus data peminjaman buku berdasarkan ID
func UsersBookDelete(c *fiber.Ctx) error {
	id := c.Query("id")
	var dataUsersBook entity.UsersBook

	// Mencari data peminjaman buku berdasarkan id
	if err := database.DB.Where("id_users_book =?", id).First(&dataUsersBook).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "UsersBook not found",
		})
	}

	// Menghapus data peminjaman buku dari database
	if err := database.DB.Delete(&dataUsersBook).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "UsersBook data deleted successfully",
	})
}

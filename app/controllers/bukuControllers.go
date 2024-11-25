package controllers

import (
	"perpus/app/database"
	"perpus/app/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

// BukuCreate - Endpoint untuk membuat data buku
func BukuCreate(c *fiber.Ctx) error {
	var request entity.Buku

	// Parsing request body untuk data buku
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Membuat entri buku baru
	newBuku := entity.Buku{
		IdPerpustakaan: request.IdPerpustakaan,
		Judul:          request.Judul,
		Halaman:        request.Halaman,
		Penerbit:       request.Penerbit,
		Pengarang:      request.Pengarang,
		Tahun:          request.Tahun,
		Kategori:       request.Kategori,
		CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
	}

	// Simpan data buku ke database
	if err := database.DB.Create(&newBuku).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Buku successfully created",
		"data":    newBuku,
	})
}

// BukuGetAll - Endpoint untuk mendapatkan semua data buku
func BukuGetAll(c *fiber.Ctx) error {
	var dataBuku []entity.Buku

	// Mengambil query parameter untuk filter (misal id_buku)
	id := c.Query("id")
	baseQuery := database.DB

	if id != "" {
		baseQuery = baseQuery.Where("id_buku = ?", id)
	}

	// Query untuk mengambil data buku
	if err := baseQuery.Find(&dataBuku).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Retrieved all buku data",
		"data":    dataBuku,
	})
}

// BukuUpdate - Endpoint untuk memperbarui data buku berdasarkan ID
func BukuUpdate(c *fiber.Ctx) error {
	id := c.Query("id")
	var dataBuku entity.Buku

	// Mencari data buku berdasarkan id
	if err := database.DB.Where("id_buku =?", id).First(&dataBuku).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Buku not found",
		})
	}

	// Parsing request body untuk pembaruan
	var request entity.Buku
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update data buku
	updates := entity.Buku{
		IdPerpustakaan: request.IdPerpustakaan,
		Judul:          request.Judul,
		Halaman:        request.Halaman,
		Penerbit:       request.Penerbit,
		Pengarang:      request.Pengarang,
		Tahun:          request.Tahun,
		Kategori:       request.Kategori,
	}

	// Melakukan update ke database
	if err := database.DB.Model(&dataBuku).Updates(&updates).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Buku data updated successfully",
		"data":    dataBuku,
	})
}

// BukuDelete - Endpoint untuk menghapus data buku berdasarkan ID
func BukuDelete(c *fiber.Ctx) error {
	id := c.Query("id")
	var dataBuku entity.Buku

	// Mencari data buku berdasarkan id
	if err := database.DB.Where("id_buku =?", id).First(&dataBuku).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Buku not found",
		})
	}

	// Menghapus data buku dari database
	if err := database.DB.Delete(&dataBuku).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Buku data deleted successfully",
	})
}

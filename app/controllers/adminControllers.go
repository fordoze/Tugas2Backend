package controllers

import (
	"perpus/app/database"
	"perpus/app/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

// AdminCreate - Endpoint untuk membuat data admin
func AdminCreate(c *fiber.Ctx) error {
	var request entity.Admin

	// Parsing request body untuk data admin
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Membuat entri admin baru
	newAdmin := entity.Admin{
		IdPerpustakaan:   request.IdPerpustakaan,
		Nama:             request.Nama,
		NamaPerpustakaan: request.NamaPerpustakaan,
		Email:            request.Email,
		Password:         request.Password,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
	}

	// Simpan data admin ke database
	if err := database.DB.Create(&newAdmin).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Admin successfully created",
		"data":    newAdmin,
	})
}

// AdminGetAll
func AdminGetAll(c *fiber.Ctx) error {
	var dataAdmin []entity.Admin

	// Mengambil query parameter untuk filter (misal id_admin)
	id := c.Query("id")
	baseQuery := database.DB

	if id != "" {
		baseQuery = baseQuery.Where("id_admin = ?", id)
	}

	// Query untuk mengambil data admin
	if err := baseQuery.Find(&dataAdmin).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Retrieved all admin data",
		"data":    dataAdmin,
	})
}

// AdminUpdate - Endpoint untuk memperbarui data admin berdasarkan ID
func AdminUpdate(c *fiber.Ctx) error {
	id := c.Query("id")
	var dataAdmin entity.Admin

	// Mencari data admin berdasarkan id
	if err := database.DB.Where("id_admin =?", id).First(&dataAdmin).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Admin not found",
		})
	}

	// Parsing request body untuk pembaruan
	var request entity.Admin
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update data admin
	updates := entity.Admin{
		IdPerpustakaan:   request.IdPerpustakaan,
		Nama:             request.Nama,
		NamaPerpustakaan: request.NamaPerpustakaan,
		Email:            request.Email,
		Password:         request.Password, // Harus di-hash dalam aplikasi nyata
	}

	// Melakukan update ke database
	if err := database.DB.Model(&dataAdmin).Updates(&updates).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Admin data updated successfully",
		"data":    dataAdmin,
	})
}

// AdminDelete - Endpoint untuk menghapus data admin berdasarkan ID
func AdminDelete(c *fiber.Ctx) error {
	id := c.Query("id")
	var dataAdmin entity.Admin

	// Mencari data admin berdasarkan id
	if err := database.DB.Where("id_admin =?", id).First(&dataAdmin).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Admin not found",
		})
	}

	// Menghapus data admin dari database
	if err := database.DB.Delete(&dataAdmin).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Admin data deleted successfully",
	})
}

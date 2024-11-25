package controllers

import (
	"perpus/app/database"
	"perpus/app/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreatePerpustakaan(c *fiber.Ctx) error {

	var Request entity.Perpustakaan

	if err := c.BodyParser(&Request); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newPerpustakaan := entity.Perpustakaan{
		Nama:      Request.Nama,
		Lokasi:    Request.Lokasi,
		Provinsi:  Request.Provinsi,
		Kota:      Request.Kota,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	//Simpan Data New Perpustakaan

	if err := database.DB.Create(&newPerpustakaan).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//Kita Bisa Isi Data data dari perpustakaannya

	return c.JSON(fiber.Map{
		"Pesan": "Create Book",
		"data":  newPerpustakaan,
	})
}

func GetAllDataPerpustakaan(c *fiber.Ctx) error {

	var DataPerpus []entity.Perpustakaan

	//memakai pameter query Id
	id := c.Query("id")

	baseQuery := database.DB

	if id != "" {
		baseQuery = baseQuery.Where("id_perpustakaan = ?", id)
	}

	//Cari Semua Data Base n
	baseQuery.Find(&DataPerpus)
	//Kita Bisa Isi Data data dari perpustakaannya

	return c.JSON(fiber.Map{
		"Pesan": "Get All Data Perpustakaan",
		"data":  DataPerpus,
	})

	return c.JSON(fiber.Map{
		"Pesan": "Get All Data Perpustakaan",
		"data":  nil,
	})
}

func UpdatedAtPerpustakaan(c *fiber.Ctx) error {

	//Ambil get Data By Id

	id := c.Query("id")
	var DataPerpus entity.Perpustakaan

	database.DB.Where("id_perpustakaan =?", id).First(&DataPerpus)

	if DataPerpus.IdPerpustakaan == 0 {
		return c.Status(404).JSON(fiber.Map{
			"Pesan": "Data Tidak Ditemukan",
		})
	}

	var Request entity.Perpustakaan
	if err := c.BodyParser(&Request); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//Jika Pake Model
	updates := entity.Perpustakaan{
		Nama:      Request.Nama,
		Lokasi:    Request.Lokasi,
		Provinsi:  Request.Provinsi,
		Kota:      Request.Kota,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	database.DB.Model(&DataPerpus).Updates(&updates)

	return c.JSON(fiber.Map{
		"Pesan": "Update Data Perpustakaan",
		"data":  DataPerpus,
	})
}

func DeleteDataPerpustakaan(c *fiber.Ctx) error {

	//Ambil get Data By Id

	id := c.Query("id")

	var DataPerpus entity.Perpustakaan
	database.DB.Where("id_perpustakaan =?", id).First(&DataPerpus)

	if DataPerpus.IdPerpustakaan == 0 {
		return c.Status(404).JSON(fiber.Map{
			"Pesan": "Data Tidak Ditemukan",
		})
	}

	database.DB.Delete(&DataPerpus)

	return c.JSON(fiber.Map{
		"Pesan": "Delete Data Perpustakaan",
	})
}

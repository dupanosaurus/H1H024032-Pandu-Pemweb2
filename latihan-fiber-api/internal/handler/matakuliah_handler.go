package handler

import (
	"errors"
	"strconv"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func DaftarMatakuliah(c fiber.Ctx) error {
	halaman, _ := strconv.Atoi(c.Query("halaman", "1"))
	perHalaman, _ := strconv.Atoi(c.Query("per_halaman", "10"))

	if halaman < 1 {
		halaman = 1
	}

	if perHalaman < 1 {
		perHalaman = 10
	}

	if perHalaman > 100 {
		perHalaman = 100
	}

	data, total, err := repository.AmbilMatakuliah(halaman, perHalaman)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	totalHalaman := int((total + int64(perHalaman) - 1) / int64(perHalaman))

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   data,
		"meta": fiber.Map{
			"halaman":       halaman,
			"per_halaman":   perHalaman,
			"total":         total,
			"total_halaman": totalHalaman,
		},
	})
}

func DetailMatakuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mata kuliah tidak valid",
		})
	}

	data, err := repository.AmbilMatakuliahPerId(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Mata kuliah tidak ditemukan",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   data,
	})
}

func BuatMatakuliah(c fiber.Ctx) error {
	var permintaan PermintaanMatakuliah

	if err := c.Bind().Body(&permintaan); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format permintaan tidak valid",
		})
	}

	if permintaan.Kode == "" ||
		permintaan.Nama == "" ||
		permintaan.Sks < 1 ||
		permintaan.Semester < 1 ||
		permintaan.Semester > 8 {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data mata kuliah tidak lengkap atau tidak valid",
		})
	}

	data := model.Matakuliah{
		Kode:     permintaan.Kode,
		Nama:     permintaan.Nama,
		Sks:      permintaan.Sks,
		Semester: permintaan.Semester,
	}

	if err := repository.SimpanMatakuliah(&data); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menyimpan data mata kuliah",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil ditambahkan",
		"data":   data,
	})
}

func PerbaruiMatakuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mata kuliah tidak valid",
		})
	}

	var permintaan PermintaanMatakuliah

	if err := c.Bind().Body(&permintaan); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format permintaan tidak valid",
		})
	}

	if permintaan.Kode == "" ||
		permintaan.Nama == "" ||
		permintaan.Sks < 1 ||
		permintaan.Semester < 1 ||
		permintaan.Semester > 8 {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data mata kuliah tidak lengkap atau tidak valid",
		})
	}

	data := model.Matakuliah{
		ID:       uint(id),
		Kode:     permintaan.Kode,
		Nama:     permintaan.Nama,
		Sks:      permintaan.Sks,
		Semester: permintaan.Semester,
	}

	if err := repository.PerbaruiMatakuliah(&data); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Mata kuliah tidak ditemukan",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal memperbarui mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil diperbarui",
		"data":   data,
	})
}

func HapusMatakuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mata kuliah tidak valid",
		})
	}

	err = repository.HapusMatakuliah(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Mata kuliah tidak ditemukan",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menghapus mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil dihapus",
	})
}

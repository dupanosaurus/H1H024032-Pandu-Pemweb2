package handler

import (
	"errors"
	"strconv"

	"latihan-fiber-api/internal/repository"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func PulihkanMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mahasiswa tidak valid",
		})
	}

	err = repository.Pulihkan(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Mahasiswa tidak ditemukan atau belum dihapus",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal memulihkan mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mahasiswa berhasil dipulihkan",
	})
}

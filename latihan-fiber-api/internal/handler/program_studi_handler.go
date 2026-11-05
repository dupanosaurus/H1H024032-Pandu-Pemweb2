package handler

import (
	"latihan-fiber-api/internal/repository"

	"github.com/gofiber/fiber/v3"
)

func DaftarProgramStudi(c fiber.Ctx) error {
	data, err := repository.AmbilProgramStudi()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data program studi",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   data,
	})
}

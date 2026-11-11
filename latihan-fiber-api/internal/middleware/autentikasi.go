package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func Autentikasi() fiber.Handler {
	return func(c fiber.Ctx) error {
		header := c.Get("Authorization")

		if header == "" {
			return fiber.NewError(401, "Token autentikasi diperlukan")
		}

		bagian := strings.SplitN(header, " ", 2)

		if len(bagian) != 2 || bagian[0] != "Bearer" {
			return fiber.NewError(401, "Format token tidak valid")
		}

		tokenAsli := bagian[1]

		konf := config.Muat()

		klaim, err := auth.PeriksaToken(
			tokenAsli,
			konf.JwtSecret,
		)
		if err != nil {
			return fiber.NewError(401, "Token tidak valid atau kedaluwarsa")
		}

		// Menghitung hash SHA-256 dari token.
		hash := sha256.Sum256([]byte(tokenAsli))
		tokenHash := hex.EncodeToString(hash[:])

		// Memeriksa apakah token terdapat dalam blacklist.
		var blacklist model.TokenBlacklist

		err = config.DB.Where("token_hash = ?", tokenHash).
			First(&blacklist).Error

		if err == nil {
			return fiber.NewError(401, "Token sudah tidak berlaku")
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(500, "Gagal memeriksa blacklist token")
		}

		idPengguna, err := strconv.ParseUint(klaim.Subject, 10, 64)
		if err != nil {
			return fiber.NewError(401, "Identitas pengguna tidak valid")
		}

		c.Locals("id_pengguna", uint(idPengguna))
		c.Locals("peran", klaim.Peran)
		c.Locals("token", tokenAsli)

		if klaim.ExpiresAt != nil {
			c.Locals("token_exp", klaim.ExpiresAt.Time)
		}

		return c.Next()
	}
}

func HanyaPeran(peranDiizinkan ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		peran, cocok := c.Locals("peran").(string)

		if !cocok {
			return fiber.NewError(403, "Akses ditolak")
		}

		for _, peranDiizinkan := range peranDiizinkan {
			if peran == peranDiizinkan {
				return c.Next()
			}
		}

		return fiber.NewError(403, "Anda tidak memiliki izin untuk mengakses sumber daya ini")
	}
}

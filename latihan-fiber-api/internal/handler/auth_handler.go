package handler

import (
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/validasi"

	"crypto/sha256"
	"encoding/hex"
	"time"
)

func Registrasi(c fiber.Ctx) error {
	var permintaan PermintaanRegistrasi

	err := c.Bind().Body(&permintaan)
	if err != nil {
		return fiber.NewError(400, "Format body tidak valid")
	}

	galat := validasi.Periksa(permintaan)
	if galat != nil {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data registrasi tidak valid",
			"galat":  galat,
		})
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(permintaan.KataSandi),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fiber.NewError(500, "Gagal memproses kata sandi")
	}

	pengguna := model.Pengguna{
		Nama:      permintaan.Nama,
		Email:     permintaan.Email,
		KataSandi: string(hash),
		Peran:     "mahasiswa",
	}

	err = config.DB.Create(&pengguna).Error
	if err != nil {
		return fiber.NewError(422, "Email sudah terdaftar")
	}

	return c.Status(201).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Registrasi berhasil",
		"data":   pengguna,
	})
}

func Login(c fiber.Ctx) error {
	var permintaan PermintaanLogin

	err := c.Bind().Body(&permintaan)
	if err != nil {
		return fiber.NewError(400, "Format body tidak valid")
	}

	galat := validasi.Periksa(permintaan)
	if galat != nil {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data login tidak valid",
			"galat":  galat,
		})
	}

	var pengguna model.Pengguna

	err = config.DB.Where("email = ?", permintaan.Email).
		First(&pengguna).Error
	if err != nil {
		return fiber.NewError(401, "Email atau kata sandi tidak sesuai")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(pengguna.KataSandi),
		[]byte(permintaan.KataSandi),
	)
	if err != nil {
		return fiber.NewError(401, "Email atau kata sandi tidak sesuai")
	}

	konf := config.Muat()

	token, err := auth.TerbitkanToken(
		pengguna.ID,
		pengguna.Peran,
		konf.JwtSecret,
		konf.JwtTtlMenit,
	)
	if err != nil {
		return fiber.NewError(500, "Gagal menerbitkan token")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Login berhasil",
		"data": fiber.Map{
			"pengguna": pengguna,
			"token":    token,
			"tipe":     "Bearer",
		},
	})
}

func Profil(c fiber.Ctx) error {
	idPengguna := c.Locals("id_pengguna")

	var pengguna model.Pengguna

	err := config.DB.First(&pengguna, idPengguna).Error
	if err != nil {
		return fiber.NewError(404, "Pengguna tidak ditemukan")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   pengguna,
	})
}

func SegarkanToken(c fiber.Ctx) error {
	idPengguna, cocok := c.Locals("id_pengguna").(uint)
	if !cocok {
		return fiber.NewError(401, "Identitas pengguna tidak valid")
	}

	peran, cocok := c.Locals("peran").(string)
	if !cocok {
		return fiber.NewError(401, "Peran pengguna tidak valid")
	}

	konf := config.Muat()

	token, err := auth.TerbitkanToken(
		idPengguna,
		peran,
		konf.JwtSecret,
		konf.JwtTtlMenit,
	)
	if err != nil {
		return fiber.NewError(500, "Gagal menerbitkan token baru")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Token berhasil diperbarui",
		"data": fiber.Map{
			"token": token,
			"tipe":  "Bearer",
		},
	})
}

func Logout(c fiber.Ctx) error {
	// Mengambil token dari middleware.
	tokenAsli, cocok := c.Locals("token").(string)
	if !cocok {
		return fiber.NewError(401, "Token tidak ditemukan")
	}

	// Mengambil waktu kedaluwarsa token.
	waktuKedaluwarsa, cocok := c.Locals("token_exp").(time.Time)
	if !cocok {
		return fiber.NewError(401, "Waktu kedaluwarsa token tidak valid")
	}

	// Menghitung hash SHA-256 token.
	hash := sha256.Sum256([]byte(tokenAsli))
	tokenHash := hex.EncodeToString(hash[:])

	// Menyimpan token ke blacklist.
	blacklist := model.TokenBlacklist{
		TokenHash: tokenHash,
		ExpiresAt: waktuKedaluwarsa,
	}

	err := config.DB.Create(&blacklist).Error
	if err != nil {
		return fiber.NewError(500, "Gagal memproses logout")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Logout berhasil. Token tidak dapat digunakan kembali",
	})
}

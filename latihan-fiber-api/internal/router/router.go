package router

import (
	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/middleware"
)

func Daftarkan(app *fiber.App) {
	api := app.Group("/api")

	// Endpoint status
	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	// Endpoint autentikasi (tanpa login)
	api.Post("/auth/registrasi", handler.Registrasi)
	api.Post("/auth/login", handler.Login)

	// Grup endpoint yang membutuhkan JWT
	terlindungi := api.Group("", middleware.Autentikasi())

	// Profil pengguna yang sedang login
	terlindungi.Get("/auth/profil", handler.Profil)
	terlindungi.Post("/auth/logout", handler.Logout)

	// Penyegaran token JWT
	terlindungi.Post("/auth/segarkan", handler.SegarkanToken)

	// CRUD mahasiswa membutuhkan autentikasi JWT
	mahasiswa := terlindungi.Group("/mahasiswa")

	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)

	// Membuat dan memperbarui data hanya untuk admin dan dosen
	mahasiswa.Post("/", middleware.HanyaPeran("admin", "dosen"), handler.BuatMahasiswa)
	mahasiswa.Put("/:id", middleware.HanyaPeran("admin", "dosen"), handler.PerbaruiMahasiswa)

	// Menghapus dan memulihkan data hanya untuk admin
	mahasiswa.Delete("/:id", middleware.HanyaPeran("admin"), handler.HapusMahasiswa)
	mahasiswa.Post("/:id/pulihkan", middleware.HanyaPeran("admin"), handler.PulihkanMahasiswa)

	// CRUD matakuliah tetap menggunakan API Key
	matakuliah := api.Group("/matakuliah")
	matakuliah.Use(middleware.KunciApi())

	matakuliah.Get("/", handler.DaftarMatakuliah)
	matakuliah.Post("/", handler.BuatMatakuliah)
	matakuliah.Get("/:id", handler.DetailMatakuliah)
	matakuliah.Put("/:id", handler.PerbaruiMatakuliah)
	matakuliah.Delete("/:id", handler.HapusMatakuliah)

	// Program studi tetap menggunakan API Key
	programStudi := api.Group("/program_studi")
	programStudi.Use(middleware.KunciApi())

	programStudi.Get("/", handler.DaftarProgramStudi)
}

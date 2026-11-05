package repository

import (
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"

	"gorm.io/gorm"
)

func AmbilMatakuliah(halaman, perHalaman int) ([]model.Matakuliah, int64, error) {
	var daftar []model.Matakuliah
	var total int64

	err := config.DB.Model(&model.Matakuliah{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (halaman - 1) * perHalaman

	err = config.DB.
		Order("nama ASC").
		Limit(perHalaman).
		Offset(offset).
		Find(&daftar).Error

	return daftar, total, err
}

func AmbilMatakuliahPerId(id uint) (model.Matakuliah, error) {
	var matakuliah model.Matakuliah

	err := config.DB.First(&matakuliah, id).Error

	return matakuliah, err
}

func SimpanMatakuliah(data *model.Matakuliah) error {
	return config.DB.Create(data).Error
}

func PerbaruiMatakuliah(data *model.Matakuliah) error {
	hasil := config.DB.Model(&model.Matakuliah{}).
		Where("id = ?", data.ID).
		Updates(map[string]interface{}{
			"kode":     data.Kode,
			"nama":     data.Nama,
			"sks":      data.Sks,
			"semester": data.Semester,
		})

	if hasil.Error != nil {
		return hasil.Error
	}

	if hasil.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func HapusMatakuliah(id uint) error {
	hasil := config.DB.Delete(&model.Matakuliah{}, id)

	if hasil.Error != nil {
		return hasil.Error
	}

	if hasil.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

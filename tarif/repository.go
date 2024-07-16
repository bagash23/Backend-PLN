package tarif

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateTarif(tarif Tarif)(Tarif, error)
	FindAll()([]Tarif, error)
	FindIDTarif(IDTarif string)(Tarif, error)
	Update(IDTarif string, tarif ReqUpdateTarif)(Tarif, error)
	Delete(IDTarif string)(error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) CreateTarif(tarif Tarif)(Tarif, error) {
	err := r.db.Create(&tarif).Error
	if err != nil {
		return tarif, err
	}
	return tarif, nil
}

func (r *repository) FindAll()([]Tarif, error) {
	var tarif []Tarif
	err := r.db.Find(&tarif).Error

	if err != nil {
		return nil, err
	}

	return tarif, nil
}

func (r *repository) Update(IDTarif string, tarif ReqUpdateTarif)(Tarif, error) {
	var resp Tarif
	result := r.db.Model(&resp).Where("id_tarif = ?", IDTarif).Updates(&tarif).Error
	if result != nil {
		return resp, result
	}
	return resp, nil
}

func (r *repository) Delete(IDTarif string)(error) {
	var tarif Tarif
	err := r.db.Model(&tarif).Where("id_tarif = ?", IDTarif).Delete(&tarif).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindIDTarif(IDTarif string)(Tarif, error) {
	var tarif Tarif
	err := r.db.First(&tarif, "id_tarif = ?", IDTarif).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return tarif, nil
		}
		return tarif, err
	}
	return tarif, nil
}
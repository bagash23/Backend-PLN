package pelanggan

import "gorm.io/gorm"

type Repository interface {
	CreatePelanggan(pelanggan Pelanggan)(Pelanggan, error)
	FindIDPelanggan(IDPelanggan string)(Pelanggan, error)
	FindAll()([]Pelanggan, error)
	FindUsernamePelanggan(username string)(Pelanggan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository)CreatePelanggan(pelanggan Pelanggan)(Pelanggan, error) {
	err := r.db.Create(&pelanggan).Error
	if err != nil {
		return pelanggan, err
	}
	return pelanggan, nil
}

func (r *repository) FindIDPelanggan(IDPelanggan string)(Pelanggan, error) {
	var pelanggan Pelanggan		
	err := r.db.First(&pelanggan, "id_pelanggan = ?", IDPelanggan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return pelanggan, nil
		}
		return pelanggan, err
	}
	return pelanggan, nil
}

func (r *repository) FindUsernamePelanggan(username string)(Pelanggan, error) {
	var user Pelanggan
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindAll()([]Pelanggan, error) {
	var pelanggan []Pelanggan
	err := r.db.Find(&pelanggan).Error
	if err != nil {
		return nil, err
	}
	return pelanggan, nil
}
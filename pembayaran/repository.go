package pembayaran

import "gorm.io/gorm"

type Repository interface {
	CreatePembayaran(pembayaran Pembayaran)(Pembayaran, error)
	FindIDPembayaran(IDPembayaran string)(Pembayaran, error)
	FindID(ID string)([]Pembayaran, error) 
	FindAll()([]Pembayaran, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePembayaran(pembayaran Pembayaran)(Pembayaran, error) {
	err := r.db.Create(&pembayaran).Error
	if err != nil {
		return pembayaran,err
	}
	return pembayaran, nil
}

func (r *repository) FindIDPembayaran(IDPembayaran string)(Pembayaran, error) {
	var pembayaran Pembayaran		
	err := r.db.First(&pembayaran, "id_pembayaran = ?", IDPembayaran).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return pembayaran, nil
		}
		return pembayaran, err
	}
	return pembayaran, nil
}

func (r *repository) FindAll()([]Pembayaran, error) {
	var pembayaran []Pembayaran
	err := r.db.Find(&pembayaran).Error
	if err != nil {
		return nil, err
	}
	return pembayaran, nil
}

func (r *repository) FindID(ID string)([]Pembayaran, error) {
	var pembayaran []Pembayaran	
	err := r.db.Where("id_pembayaran = ?", ID).Find(&pembayaran).Error
	if err != nil {
		return nil, err
	}
	return pembayaran, nil
}
package penggunaan

import "gorm.io/gorm"

type Repository interface {
	CreatePenggunaan(penggunaan Penggunaan)(Penggunaan, error)
	FindByPenggunaan(IDPenggunaan string)([]Penggunaan, error)
	FindIDPenggunaan(IDPenggunaan string)(Penggunaan, error)
	FindAll()([]Penggunaan, error)
	Update(IDPenggunaan string, penggunaan ReqUpdatePenggunaan)(Penggunaan, error)
	Delete(IDPenggunaan string)(error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
} 

func (r *repository) CreatePenggunaan(penggunaan Penggunaan)(Penggunaan, error) {
	err := r.db.Create(&penggunaan).Error
	if err != nil {
		return penggunaan,err
	}
	return penggunaan, nil
}

func (r *repository) FindByPenggunaan(IDPenggunaan string)([]Penggunaan, error) {
	var penggunaan []Penggunaan	
	err := r.db.Where("id_penggunaan = ?", IDPenggunaan).Find(&penggunaan).Error
	if err != nil {
		return nil, err
	}
	return penggunaan, nil
}

func (r *repository) FindAll()([]Penggunaan, error) {
	var penggunaan []Penggunaan
	err := r.db.Find(&penggunaan).Error
	if err != nil {
		return nil, err
	}
	return penggunaan, nil
}

func (r *repository) Update(IDPenggunaan string, penggunaan ReqUpdatePenggunaan)(Penggunaan, error) {
	var resp Penggunaan
	result := r.db.Model(&resp).Where("id_penggunaan = ?", IDPenggunaan).Updates(&penggunaan).Error
	if result != nil {
		return resp, result
	}
	return resp, nil

}

func (r *repository) Delete(IDPenggunaan string) (error) {
	var penggunaan Penggunaan
	err := r.db.Model(&penggunaan).Where("id_penggunaan = ?", IDPenggunaan).Delete(&penggunaan).Error
	if err != nil {
	  return err 
	}
	return nil
}


func (r *repository) FindIDPenggunaan(IDPenggunaan string)(Penggunaan, error) {
	var penggunaan Penggunaan		
	err := r.db.First(&penggunaan, "id_penggunaan = ?", IDPenggunaan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return penggunaan, nil
		}
		return penggunaan, err
	}
	return penggunaan, nil
}
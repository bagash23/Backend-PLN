package tagihan

import "gorm.io/gorm"

type Repository interface {
	CreateTagihan(tagihan Tagihan)(Tagihan, error)
	FindAll()([]Tagihan, error)
	FindIDTagihan(IDTagihan string)(Tagihan, error)
	UpdateTagihan(IDTagihan string, tagihan ReqUpdateTagihan)(Tagihan, error)
	DeleteTagihan(IDTagihan string)(error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTagihan(tagihan Tagihan)(Tagihan, error) {
	err := r.db.Create(&tagihan).Error
	if err != nil {
		return tagihan, err
	}
	return tagihan, nil
}

func (r *repository) FindAll()([]Tagihan, error) {
	var tagihan []Tagihan
	err := r.db.Find(&tagihan).Error
	if err != nil {
		return nil, err
	}
	return tagihan, nil
}

func (r *repository) FindIDTagihan(IDTagihan string)(Tagihan, error) {
	var tagihan Tagihan		
	err := r.db.First(&tagihan, "id_tagihan = ?", IDTagihan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return tagihan, nil
		}
		return tagihan, err
	}
	return tagihan, nil
}

func (r *repository) UpdateTagihan(IDTagihan string, tagihan ReqUpdateTagihan)(Tagihan, error) {
	var resp Tagihan
	result := r.db.Model(&resp).Where("id_tagihan = ?", IDTagihan).Updates(&tagihan).Error
	if result != nil {
		return resp, result
	}
	return resp, nil
}

func (r *repository) DeleteTagihan(IDTagihan string)(error) {
	var tagihan Tagihan
	err := r.db.Model(&tagihan).Where("id_tagihan = ?", IDTagihan).Delete(&tagihan).Error
	if err != nil {
	  return err 
	}
	return nil
}
package penggunaan

import "errors"


type Service interface {
	InputPenggunaan(input PenggunaanInput)(Penggunaan, error)
	GetPenggunaan(ID string)([]Penggunaan, error)
	UpdatePenggunaan(ID string, penggunaInput ReqUpdatePenggunaan)(Penggunaan, error)
	DeletePenggunaan(ID string)(error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InputPenggunaan(input PenggunaanInput)(Penggunaan, error) {
	penggunaan, err := s.repository.FindIDPenggunaan(input.IDPenggunaan)
	
	if err != nil {
		return penggunaan, err
	}

	if penggunaan.IDPenggunaan == input.IDPenggunaan {
		return penggunaan, errors.New("ID Penggunaan already in use")
	}

	penggunaan.IDPenggunaan = input.IDPenggunaan
	penggunaan.IDPelanggan = input.IDPelanggan
	penggunaan.Bulan = input.Bulan
	penggunaan.Tahun = input.Tahun
	penggunaan.MeterAwal = input.MeterAwal
	penggunaan.MeterAkhir = input.MeterAkhir

	newPenggunaan, err := s.repository.CreatePenggunaan(penggunaan)	
	if err != nil {
		return newPenggunaan, err
	}

	return newPenggunaan, nil
}

func (s *service) GetPenggunaan(ID string)([]Penggunaan, error) {
	if ID != "" {
		penggunaans, err := s.repository.FindByPenggunaan(ID)
		if err != nil {
			return penggunaans, err
		}
		return penggunaans, nil
	}
	penggunaans, err := s.repository.FindAll()
	if err != nil {
		return penggunaans, err
	}
	return penggunaans, nil
}

func (s *service) UpdatePenggunaan(ID string, penggunaInput ReqUpdatePenggunaan)(Penggunaan, error) {	 
	penggunaan, err := s.repository.FindIDPenggunaan(ID)
	if err != nil {
		return penggunaan, err
	}

	updatePenggunaan, err := s.repository.Update(penggunaan.IDPenggunaan, penggunaInput)
	if err != nil {
		return penggunaan, err
	}
	return updatePenggunaan, nil
}

func (s *service) DeletePenggunaan(ID string)(error) {	
	err := s.repository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
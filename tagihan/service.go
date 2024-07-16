package tagihan

import "errors"

type Service interface {
	InputService(input TagihanInput)(Tagihan, error)
	GetTagihan()([]Tagihan, error)
	UpdateTagihan(ID string, penggunaInput ReqUpdateTagihan)(Tagihan, error)
	DeleteTagihan(ID string)(error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InputService(input TagihanInput)(Tagihan, error) {
	tagihan, err := s.repository.FindIDTagihan(input.IDTagihan)
	if err != nil {
		return tagihan, err
	}
	if tagihan.IDTagihan == input.IDTagihan {
		return tagihan, errors.New("ID Tagihan already in use")
	}

	tagihan.IDTagihan = input.IDTagihan
	tagihan.IDPenggunaan = input.IDPenggunaan
	tagihan.IDPelanggan = input.IDPelanggan
	tagihan.Bulan = input.Bulan
	tagihan.Tahun = input.Tahun
	tagihan.JumlahMeter = input.JumlahMeter
	tagihan.Status = input.Status

	newTagihan, err := s.repository.CreateTagihan(tagihan)
	if err != nil {
		return newTagihan, err
	}
	return newTagihan, nil
}


func (s *service) GetTagihan()([]Tagihan, error) {
	tagihan, err := s.repository.FindAll()
	if err != nil {
		return tagihan, err
	}
	return tagihan, nil
}

func (s *service) UpdateTagihan(ID string, penggunaInput ReqUpdateTagihan)(Tagihan, error) {
	tagihan, err := s.repository.FindIDTagihan(ID)
	if err != nil {
		return tagihan, err
	}
	updateTagihan, err := s.repository.UpdateTagihan(tagihan.IDTagihan, penggunaInput)

	if err != nil {
		return tagihan, err
	}
	return updateTagihan, nil
}

func (s *service) DeleteTagihan(ID string)(error) {
	err := s.repository.DeleteTagihan(ID)
	if err != nil {
		return err
	}
	return nil
}
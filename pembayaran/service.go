package pembayaran

import "errors"

type Service interface {
	InputPembayaran(input PembayaranInput)(Pembayaran, error)
	GetPembayaran(ID string)([]Pembayaran, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InputPembayaran(input PembayaranInput)(Pembayaran, error) {
	pembayaran, err := s.repository.FindIDPembayaran(input.IDPembayaran)

	if err != nil {
		return pembayaran, err
	}

	if pembayaran.IDPembayaran == input.IDPembayaran {
		return pembayaran, errors.New("ID Pembayaran already in use")
	}

	pembayaran.IDPembayaran = input.IDPembayaran
	pembayaran.IDTagihan = input.IDTagihan
	pembayaran.IDPelanggan = input.IDPelanggan
	pembayaran.TglBayar = input.TglBayar
	pembayaran.BiayaAdmin = input.BiayaAdmin
	pembayaran.TotalBayar = input.TotalBayar
	pembayaran.IDUser = input.IDUser

	newPembayaran, err := s.repository.CreatePembayaran(pembayaran)
	if err != nil {
		return newPembayaran, err
	}

	return newPembayaran, nil
}

func (s *service)GetPembayaran(ID string)([]Pembayaran, error) {

	if ID != "" {
		pembayarans, err := s.repository.FindID(ID)
		if err != nil {
			return pembayarans, err
		}
		return pembayarans, nil
	}

	pembayaran, err := s.repository.FindAll()
	if err != nil {
		return pembayaran, err
	}
	return pembayaran, nil
}
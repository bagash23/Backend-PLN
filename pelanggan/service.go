package pelanggan

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	InputPelanggan(input PelangganInput)(Pelanggan, error)	
	GetPelanggan()([]Pelanggan, error)
	GetFindID(ID string)(Pelanggan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InputPelanggan(input PelangganInput)(Pelanggan, error) {
	pelanggan, err := s.repository.FindIDPelanggan(input.IDPelanggan)
	if err != nil {
		return pelanggan, err
	}
	if pelanggan.IDPelanggan == input.IDPelanggan {
		return pelanggan, errors.New("ID pelanggan already in use")
	}

	pelanggan.IDPelanggan = input.IDPelanggan
	pelanggan.NamaPelanggan = input.NamaPelanggan
	pelanggan.Username = input.Username	
	pelanggan.NomorKWH = input.NomorKWH
	pelanggan.Alamat = input.Alamat
	pelanggan.IDTarif = input.IDTarif

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	pelanggan.Password = string(passwordHash)
	if err != nil {
		return pelanggan, err
	}



	newPelanggan, err := s.repository.CreatePelanggan(pelanggan)
	if err != nil {
		return newPelanggan, err
	}
	return newPelanggan, nil
}

func (s *service) GetPelanggan()([]Pelanggan, error) {
	pelanggan, err := s.repository.FindAll()
	if err != nil {
		return pelanggan, err
	}
	return pelanggan, nil
}

func (s *service) GetFindID(ID string)(Pelanggan, error) {
	pelanggan, err := s.repository.FindIDPelanggan(ID)
	if err != nil {
		return pelanggan, err
	}

	if pelanggan.IDPelanggan == "" {
		return pelanggan, errors.New("No pelanggan found on with that ID")
	}

	return pelanggan, nil
}
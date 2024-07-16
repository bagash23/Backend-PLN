package tarif

import (
	"errors"
)

type Service interface {
	TarifInputUser(input TarifInput)(Tarif, error)
	GetTarif()([]Tarif, error)
	UpdateTarif(ID string, penggunaInput ReqUpdateTarif)(Tarif, error)
	DeleteTarif(ID string)(error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) TarifInputUser(input TarifInput)(Tarif, error) {	
	existingTarif, err := s.repository.FindIDTarif(input.IDTarif)
	if err != nil {		
		return existingTarif, err
	}
	if existingTarif.IDTarif == input.IDTarif {
		return existingTarif, errors.New("ID Tarif already in use")
	}
	
	existingTarif.IDTarif = input.IDTarif
	existingTarif.Daya = input.Daya
	existingTarif.TarifPerkwh = input.TarifPerkwh
	
	newTarif, err := s.repository.CreateTarif(existingTarif)
	if err != nil {
		return newTarif, err
	}

	return newTarif, nil
}


func (s *service) GetTarif()([]Tarif, error) {
	tarif, err := s.repository.FindAll()
	if err != nil {
		return tarif, err
	}
	return tarif, nil
}

func (s *service) UpdateTarif(ID string, penggunaInput ReqUpdateTarif)(Tarif, error) {
	tarif, err := s.repository.FindIDTarif(ID)
	if err != nil {
		return tarif, err
	}
	updateTarif, err := s.repository.Update(tarif.IDTarif, penggunaInput)

	if err != nil {
		return tarif, err
	}
	return updateTarif, nil
}

func (s *service) DeleteTarif(ID string)(error) {
	err := s.repository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
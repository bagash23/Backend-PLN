package user

import (
	"errors"
	"pln/pelanggan"

	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterUserInput)(User, error)
	Login(input LoginUserInput)(User, error)
	GetUserByID(id string)(User, error)
}

type service struct {
	repository Repository
	repositoryPelanggan pelanggan.Repository
}

func NewService(repository Repository, repositoryPelanggan pelanggan.Repository) *service {
	return &service{repository, repositoryPelanggan}
}

func (s *service) RegisterUser(input RegisterUserInput)(User, error) {
	user := User{}
	user.IDUser = input.IDUser
	user.Username = input.Username	
	user.NamaAdmin = input.NamaAdmin
	user.IDLevel = input.IDLevel

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	newUser, err :=  s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginUserInput)(User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.Username != username {
		
		pelanggan, err := s.repositoryPelanggan.FindUsernamePelanggan(username)
		if err != nil {
			return user, errors.New("no user found with that username")
		}
		
		err = bcrypt.CompareHashAndPassword([]byte(pelanggan.Password), []byte(password))
		if err != nil {
			return user, errors.New("password no match")
		}

		user.IDUser = pelanggan.IDPelanggan
		user.Username = pelanggan.Username
		user.Password = pelanggan.Password
		user.NamaAdmin = pelanggan.NamaPelanggan
		
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return user, errors.New("password no match")
		}
	}
	return user, nil
}

func (s *service) GetUserByID(id string)(User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	if user.IDUser == "" {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}
package services

import (
	"errors"
	"sportify_services/models"
	"sportify_services/repositories"
	"sportify_services/requests"
)

type Player interface {
	FindAll() ([]models.Player, error)
	FindById(ID int) (models.Player, error)
	Create(request requests.Player) (models.Player, error) 
	Update(ID int, request requests.Player) (models.Player, error) 
	Delete(ID int) (models.Player, error) 
}

type player struct {
	repository repositories.Player 
}

func NewPlayerService(r repositories.Player) *player {
	return &player{r}
}

func (s *player) Create(r requests.Player) (models.Player, error) {
	data := models.Player{ 
		UserId   : r.UserId,
		Name     : r.Name,
		Phone    : r.Phone,
		Email    : r.Email,
		Password : r.Password,
		Address  : r.Address,
		Latitude : r.Latitude,
		Longitude: r.Longitude,
	}

	result, err := s.repository.Create(data)
	return result, err
}

func (s *player) FindAll() ([]models.Player, error) {
	players, err := s.repository.FindAll()
	return players, err
}

func (s *player) FindById(ID int) (models.Player, error) {
	player, err := s.repository.FindById(ID)

	if player.ID == 0 {
		return player, errors.New("404")
	}

	return player, err
}

func (s *player) Update(ID int, r requests.Player) (models.Player, error) {
	player, err := s.repository.FindById(ID)
	if err != nil {
		return player, err
	}

	if player.ID == 0 {
		return player, errors.New("404")
	}

	player.Name      = r.Name
	player.Phone     = r.Phone
	player.Email     = r.Email
	player.Password  = r.Password
	player.Address   = r.Address
	player.Latitude  = r.Latitude
	player.Longitude = r.Longitude

	result, err := s.repository.Update(player)
	return result, err
}

func (s *player) Delete(ID int) (models.Player, error) {
	player, err := s.repository.FindById(ID)
	if err != nil {
		return player, err
	}

	if player.ID == 0 {
		return player, errors.New("404")
	}

	result, err := s.repository.Delete(player)
	return result, err
}

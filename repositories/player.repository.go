package repositories

import (
	"sportify_services/models"
	"time"

	"gorm.io/gorm"
)

type Player interface {
	FindAll() ([]models.Player, error)
	FindById(ID int) (models.Player, error)
	Create(m models.Player) (models.Player, error) 
	Update(m models.Player) (models.Player, error) 
	Delete(m models.Player) (models.Player, error) 
}

type player struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *player {
	return &player{db}
}

func (r *player) Create(m models.Player) (models.Player, error) {
	err := r.db.Create(&m).Error
	return m , err
}

func (r *player) FindAll() ([]models.Player, error) {
	var players []models.Player
	err := r.db.Find(&players).Error

	time.Sleep(2 * time.Second)

	return players, err
}

func (r *player) FindById(ID int) (models.Player, error) {
	var player models.Player
	err := r.db.Find(&player, ID).Error
	return player, err
}

func (r *player) Update(m models.Player) (models.Player, error) {
	err := r.db.Save(&m).Error
	return m , err
}

func (r *player) Delete(m models.Player) (models.Player, error) {
	err := r.db.Delete(&m).Error
	return m , err
}
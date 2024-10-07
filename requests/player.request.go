package requests

type Player struct {
	ID int `json:"id"`
	UserId string `json:"user_id"`
	Name string `json:"name" binding:"required,max=20"`
	Phone string `json:"phone" binding:"required,e164"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Address string `json:"address" binding:"required,max=100"`
	Latitude string `json:"latitude" binding:"required"`
	Longitude string `json:"logitude" binding:"required"`
}
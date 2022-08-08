package entity

import "github.com/google/uuid"

type User struct {
	Id           uint32
	PublicUserId uuid.UUID
	LineUserId   string
	Image        string
	Name         string
}

func NewUser(lineUserId string, image string, name string) (user *User) {
	user = new(User)
	user.PublicUserId = uuid.New()
	user.LineUserId = lineUserId
	user.Image = image
	user.Name = name
	return
}

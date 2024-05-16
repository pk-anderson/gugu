package interfaces

import "time"

type User struct {
	UserId     string    `json:"userId"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Bio        string    `json:"bio"`
	ProfilePic []byte    `json:"profilePic"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Status     string    `json:"status"`
}

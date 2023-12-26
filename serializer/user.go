package serializer

import "blog/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	Username string `json:"username" form:"username" example:"zrq"`
	Status   string `json:"status" form:"status"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		CreateAt: user.CreatedAt.Unix(),
	}
}

package userServiceDto

import "webapi/internal/entity"

type Session struct {
	Jwt  string       `json:"jwt"`
	User *entity.User `json:"user"`
}

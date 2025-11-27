package profile

import (
	"github.com/google/uuid"
)

type profile struct {
	login   string
	isBlock bool
	passwd  string
	uuid    uuid.UUID
}

func NewProfile(login string, passwd string, isBlock bool) *profile {
	obj := profile{login: login,
		passwd:  passwd,
		isBlock: isBlock,
		uuid:    uuid.New(),
	}
	return &obj
}

func (p *profile) ChangeStatus(status bool) {
	p.isBlock = status
}

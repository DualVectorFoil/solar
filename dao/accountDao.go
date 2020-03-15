package dao

import "github.com/DualVectorFoil/solar/model"

type AccountDao interface {
	Login(id string, userName string, phoneNum string, email string, pwd string) (*model.Profile, error)
	Register(id string, userName string, phoneNum string, email string, pwd string, avatarUrl string) (*model.Profile, error)
	// TODO update account info
}

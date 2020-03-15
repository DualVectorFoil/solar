package account

import (
	"errors"
	"github.com/DualVectorFoil/solar/app/conf"
	"github.com/DualVectorFoil/solar/db"
	"github.com/DualVectorFoil/solar/model"
	"time"
)

type AccountDaoImpl struct {
	DB *db.DB
}

func NewAccountDao() *AccountDaoImpl {
	return &AccountDaoImpl{
		DB: db.DBInstance(),
	}
}

func (a *AccountDaoImpl) Login(id string, userName string, phoneNum string, email string, pwd string) (*model.Profile, error) {
	if id == "" && phoneNum == "" && userName == "" && email == "" || pwd == "" {
		return nil, errors.New("Uncorrected login info, login failed")
	}

	a.DB.Lock.Lock()
	defer a.DB.Lock.Unlock()

	var profile model.Profile
	err := a.DB.Mysql.Table(conf.PROFILE_TABLE_NAME).Where("id = ? OR user_name = ? OR phone_num = ?", id, userName, phoneNum).Find(&profile).Error
	if err != nil {
		return nil, err
	}
	if profile.AccountID == "" || profile.UserName == "" {
		return nil, errors.New("Uncorrected login info, login failed")
	}

	return &profile, nil
}

func (a *AccountDaoImpl) Register(id string, userName string, phoneNum string, email string, pwd string, avatarUrl string) (*model.Profile, error) {
	if id == "" || userName == "" || pwd == "" {
		return nil, errors.New("Uncorrected register info, register failed")
	}

	a.DB.Lock.Lock()
	defer a.DB.Lock.Unlock()

	var profile model.Profile
	a.DB.Mysql.Table(conf.PROFILE_TABLE_NAME).Where("id = ? OR user_name = ? OR phone_num = ?", id, userName, phoneNum).Find(&profile)
	if profile.AccountID != "" || profile.UserName != "" {
		return nil, errors.New("userName has registered")
	}

	profile = model.Profile{
		AccountID:           id,
		PhoneNum:     phoneNum,
		UserName:     userName,
		Email:        email,
		AvatarUrl:    avatarUrl,
		Pwd:          pwd,
		Locale:       "",
		Bio:          "",
		Followers:    0,
		Following:    0,
		ArtworkCount: 0,
		RegisterAt:   time.Now().Unix(),
		LastLoginAt:  0,
	}

	errs := a.DB.Mysql.Create(&profile).GetErrors()
	if len(errs) > 0 {
		return nil, errors.New("DB error, Register failed")
	}

	return &profile, nil
}

package serviceServer

import (
	"context"
	"github.com/DualVectorFoil/solar/dao"
	"github.com/DualVectorFoil/solar/dao/account"
	"github.com/DualVectorFoil/solar/pb"
	"github.com/DualVectorFoil/solar/util/ptr"
	"github.com/sirupsen/logrus"
	"sync"
)

type AccountServer struct {
	Dao dao.AccountDao
}

var AS *AccountServer
var ASOnce sync.Once

func NewAccountServer() *AccountServer {
	ASOnce.Do(func() {
		AS = &AccountServer{
			Dao: account.NewAccountDao(),
		}
	})
	return AS
}

func (a *AccountServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResp, error) {
	profile, err := a.Dao.Register(request.GetId(), request.GetUserName(), request.GetPhoneNum(), request.GetEmail(), request.GetPwd(), request.GetAvatarUrl())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id": request.GetId(),
			"userName": request.GetUserName(),
			"phoneNum": request.GetPhoneNum(),
			"pwd": request.GetPwd(),
			"avatarUrl": request.GetAvatarUrl(),
		}).Error("register failed, err: " + err.Error())
		return &pb.RegisterResp{
			IsSuccess:            ptr.BoolPtr(false),
			ErrMsg:               ptr.StringPtr(err.Error()),
			Profile:              nil,
		}, err
	}

	return &pb.RegisterResp{
		IsSuccess:            ptr.BoolPtr(true),
		ErrMsg:               ptr.StringPtr(""),
		Profile:              &pb.Profile{
			Id:                   ptr.StringPtr(profile.AccountID),
			PhoneNum:             ptr.StringPtr(profile.PhoneNum),
			AvatarUrl:            ptr.StringPtr(profile.AvatarUrl),
			UserName:             ptr.StringPtr(profile.UserName),
			Locale:               ptr.StringPtr(profile.Locale),
			Bio:                  ptr.StringPtr(profile.Bio),
			Followers:            ptr.Int32Ptr(int32(profile.Followers)),
			Following:            ptr.Int32Ptr(int32(profile.Following)),
			ArtworkCount:         ptr.Int32Ptr(int32(profile.ArtworkCount)),
			Pwd:                  ptr.StringPtr(profile.Pwd),
			RegisteredAt:         ptr.Int64Ptr(int64(profile.RegisterAt)),
			LastLoginAt:          ptr.Int64Ptr(int64(profile.LastLoginAt)),
			Email:                ptr.StringPtr(profile.Email),
		},
	}, nil
}

func (a *AccountServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResp, error) {
	profile, err := a.Dao.Login(request.GetId(), request.GetUserName(), request.GetPhoneNum(), request.GetEmail(), request.GetPwd())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id": request.GetId(),
			"userName": request.GetUserName(),
			"phoneNum": request.GetPhoneNum(),
			"pwd": request.GetPwd(),
		}).Error("login failed, err: " + err.Error())
		return &pb.LoginResp{
			IsSuccess:            ptr.BoolPtr(false),
			ErrMsg:               ptr.StringPtr(err.Error()),
			Profile:              nil,
		}, err
	}

	return &pb.LoginResp{
		IsSuccess:            ptr.BoolPtr(true),
		ErrMsg:               ptr.StringPtr(""),
		Profile:              &pb.Profile{
			Id:                   ptr.StringPtr(profile.AccountID),
			PhoneNum:             ptr.StringPtr(profile.PhoneNum),
			AvatarUrl:            ptr.StringPtr(profile.AvatarUrl),
			UserName:             ptr.StringPtr(profile.UserName),
			Locale:               ptr.StringPtr(profile.Locale),
			Bio:                  ptr.StringPtr(profile.Bio),
			Followers:            ptr.Int32Ptr(int32(profile.Followers)),
			Following:            ptr.Int32Ptr(int32(profile.Following)),
			ArtworkCount:         ptr.Int32Ptr(int32(profile.ArtworkCount)),
			Pwd:                  ptr.StringPtr(profile.Pwd),
			RegisteredAt:         ptr.Int64Ptr(int64(profile.RegisterAt)),
			LastLoginAt:          ptr.Int64Ptr(int64(profile.LastLoginAt)),
			Email:                ptr.StringPtr(profile.Email),
		},
	}, nil
}

package main

import (
	"fmt"
	"github.com/DualVectorFoil/solar/app/conf"
	"github.com/DualVectorFoil/solar/db"
	"github.com/DualVectorFoil/solar/manager"
	"github.com/DualVectorFoil/solar/model"
	"github.com/DualVectorFoil/solar/pb"
	"github.com/DualVectorFoil/solar/router"
	"github.com/DualVectorFoil/solar/serviceServer"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	initDB()
	defer db.CloseDB()
	initGrpcService()
	router.Init()
}

func initDB() {
	db.DBInstance().Mysql.AutoMigrate(&model.Profile{})
}

func initGrpcService() {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	pb.RegisterAccountServiceServer(grpcServer, serviceServer.NewAccountServer())
	err := manager.ServiceMangerInstance(conf.ETCD_ADDRESS).Register(conf.SERVICE_NAME, conf.LISTEN_IP, conf.SERVICE_IP, conf.SERVICE_PORT, grpcServer, conf.GRPC_TTL)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service_name": conf.SERVICE_NAME,
			"listen_ip":    conf.LISTEN_IP,
			"service_port": conf.SERVICE_PORT,
			"grpc_ttl":     conf.GRPC_TTL,
		}).Error(fmt.Sprintf("Register service %s to etcd failed, err: %s.", conf.SERVICE_NAME, err.Error()))
	}
	panic(err)
}

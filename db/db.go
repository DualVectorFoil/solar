package db

import (
	"fmt"
	"github.com/DualVectorFoil/solar/app/conf"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var dbInstance *DB
var dbInstanceOnce sync.Once

type DB struct {
	Lock  sync.Locker
	Mysql *gorm.DB
	Redis *redis.Client
}

func InitDB() {
	dbInstanceOnce.Do(func() {
		mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", conf.MYSQL_USERNAME, conf.MYSQL_PASSWORD, conf.MYSQL_IP, conf.MYSQL_PORT, conf.MYSQL_DBNAME)
		mysqlDB, err := gorm.Open("mysql", mysqlInfo)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"mySqlInfo": mysqlInfo,
				"err":       err.Error(),
			}).Fatal("Mysql init failed.")
		}
		redisCli := redis.NewClient(&redis.Options{
			Addr:        conf.REDIS_ADDR,
			Password:    conf.REDIS_PASSWORD,
			DB:          conf.REDIS_DB_NUM,
			DialTimeout: conf.REDIS_TIMEOUT,
		})
		dbInstance = &DB{
			Lock:  &sync.Mutex{},
			Mysql: mysqlDB,
			Redis: redisCli,
		}
	})
}

func CloseDB() {
	dbInstance.Lock.Lock()
	defer dbInstance.Lock.Unlock()
	if dbInstance.Mysql != nil {
		dbInstance.Mysql.Close()
	}
	if dbInstance.Redis != nil {
		dbInstance.Redis.Close()
	}
}

func DBInstance() *DB {
	if dbInstance == nil {
		InitDB()
	}
	return dbInstance
}

func (instance *DB) GetCacheKV(key string) (string, error) {
	instance.Lock.Lock()
	defer instance.Lock.Unlock()
	value, err := instance.Redis.Get(key).Result()
	return value, err
}

func (instance *DB) SetCacheKV(key string, value interface{}, expiration time.Duration) error {
	instance.Lock.Lock()
	defer instance.Lock.Unlock()
	err := instance.Redis.Set(key, value, expiration).Err()
	return err
}

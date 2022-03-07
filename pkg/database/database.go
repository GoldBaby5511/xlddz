package database

import (
	"mango/pkg/database/mchelper"
	"mango/pkg/database/mgohelper"
	"mango/pkg/database/redishelper"
	"mango/pkg/database/sqlhelper"
	"mango/pkg/log"
	"mango/third_party/go-simplejson"
	"strings"
)

type DBCollection struct {
	SqlHelperMap map[string]*sqlhelper.SqlHelper
	MCHelper     *mchelper.MCHelper
	MongoHelper  *mgohelper.MgoHelper
	RedisHelper  *redishelper.RedisHelper
}

var (
	DBC *DBCollection = nil
)

func InitDBHelper(dbConfig string) {
	if DBC != nil {
		return
	}
	DBC = &DBCollection{}
	dataBase, err := simplejson.NewJson([]byte(dbConfig))
	if err != nil {
		log.Warning("", "数据库配置异常,dbConfig=%v,err=%v", dbConfig, err)
		return
	}

	noSql, ok := dataBase.CheckGet("nosql")
	if ok {
		memcached, err := noSql.Get("memcached").String()
		if err == nil {
			DBC.MCHelper = &mchelper.MCHelper{}
			DBC.MCHelper.Init(strings.Split(memcached, ",")...)
		}

		mongodb, ok := noSql.CheckGet("mongodb")
		if ok {
			mongoHost, err := mongodb.Get("server").String()
			if err == nil {
				mongoDatabase, _ := mongodb.Get("database").String()
				mongoUserid := mongodb.Get("userid").MustString("")
				mongoPassword := mongodb.Get("password").MustString("")
				DBC.MongoHelper = &mgohelper.MgoHelper{}
				DBC.MongoHelper.Init(mongoHost, mongoDatabase, mongoUserid, mongoPassword)
			}
		}

		redis, ok := noSql.CheckGet("redis")
		if ok {
			redisServer, err := redis.Get("server").String()
			if err == nil {
				redisPassword := redis.Get("password").MustString("")
				DBC.RedisHelper = &redishelper.RedisHelper{}
				DBC.RedisHelper.Init(redisServer, redisPassword)
			}
		}
	}

	DBC.SqlHelperMap = make(map[string]*sqlhelper.SqlHelper)
	for key, _ := range dataBase.MustMap() {
		if key == "nosql" {
			continue
		}
		helper := &sqlhelper.SqlHelper{}
		helper.Init(dataBase.Get(key))
		DBC.SqlHelperMap[key] = helper
	}

	log.Info("", "数据库初始化完成")
}

func GetSqlDB(name string) *sqlhelper.SqlHelper {
	if DBC == nil {
		return nil
	}
	return DBC.SqlHelperMap[name]
}

func GetMCHelper() *mchelper.MCHelper {
	if DBC == nil {
		return nil
	}
	return DBC.MCHelper
}

func GetRedisHelper() *redishelper.RedisHelper {
	if DBC == nil {
		return nil
	}
	return DBC.RedisHelper
}

func GetMongoHelper() *mgohelper.MgoHelper {
	if DBC == nil {
		return nil
	}
	return DBC.MongoHelper
}

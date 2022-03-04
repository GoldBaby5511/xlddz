package mgohelper

import (
	"gopkg.in/mgo.v2"
	"mango/pkg/log"
	"sync"
)

type MgoHelper struct {
	server *mgo.Session
	db     *mgo.Database
}

var (
	instance *MgoHelper
	locker   sync.Mutex
)

func GetMgoInstance(server, database, username, pwd string) *MgoHelper {
	defer locker.Unlock()
	locker.Lock()

	if instance != nil {
		return instance
	}
	instance = &MgoHelper{}
	err := instance.Init(server)
	if err != nil {
		log.Error("", "MgoHelper GetMgoInstance:err=%v", err)
		return nil
	}
	instance.InitDB(database, username, pwd)
	return instance
}

func (mh *MgoHelper) Init(server string) error {
	var err error
	mh.server, err = mgo.Dial(server)
	if err != nil {
		log.Error("", "MgoHelper Init:err=%v", err)
		return err
	}
	mh.server.SetMode(mgo.Monotonic, true)
	return nil
}

func (mh *MgoHelper) InitDB(database, username, pwd string) {
	mh.db = mh.server.DB(database)
	if len(username) > 0 && len(pwd) > 0 {
		err := mh.db.Login(username, pwd)
		if err != nil {
			log.Error("", "MgoHelper InitDB:err=%v", err)
		}
	}
}

func (mh *MgoHelper) GetDB() *mgo.Database {
	return mh.db
}

func (mh *MgoHelper) Ping() error {
	return mh.server.Ping()
}

func (mh *MgoHelper) Close(server string) {
	if mh.server == nil {
		return
	}
	mh.server.Close()
}

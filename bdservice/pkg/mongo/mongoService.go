package mongo

import (
	"time"

	"gopkg.in/mgo.v2"

	"../utils/log"
	"../utils/setting"
)

type IMongoService interface {
	MongoSession() *mgo.Session
	AllUsers(s *mgo.Session) ([]User, error)
	FindUser(s *mgo.Session, id string) (User, error)
	FindUserByUsername(s *mgo.Session, user User) (User, error)
	NewUser(s *mgo.Session, user User) error
	FindWallet(s *mgo.Session, id string) (Wallet, error)
	DepositWallet(s *mgo.Session, id string, amount float64) (Wallet, error)
}

type MongoService struct {
	settings   setting.Setting
	logger     log.ILogger
	mgoSession *mgo.Session
}

func (u *MongoService) MongoSession() *mgo.Session {
	u.inject()

	if u.mgoSession == nil {
		info := &mgo.DialInfo{
			Addrs:    []string{u.settings.Mongo.Hosts},
			Timeout:  60 * time.Second,
			Database: u.settings.Mongo.Db,
			Username: u.settings.Mongo.User,
			Password: u.settings.Mongo.Pass,
		}

		var err error
		u.mgoSession, err = mgo.DialWithInfo(info)

		if err != nil {
			u.logger.WriteLog(err.Error(), log.Fatal)
		}

		u.mgoSession.SetMode(mgo.Monotonic, true)
	}

	return u.mgoSession.Clone()
}

//Injections
func (u *MongoService) inject() {
	var injLogger log.Logger

	u.logger = &injLogger
	u.settings = setting.GetSettings()
}

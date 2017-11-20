package mongo

import (
	"../utils/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Wallet   *Wallet       `json:"wallet"`
}

func (u *MongoService) AllUsers(s *mgo.Session) ([]User, error) {
	u.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(u.settings.Mongo.Db).C("users")

	var users []User
	err := c.Find(bson.M{}).All(&users)

	if err != nil {
		u.logger.WriteLog(err.Error(), log.Info)
	}

	return users, err
}

func (u *MongoService) FindUser(s *mgo.Session, id string) (User, error) {
	u.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(u.settings.Mongo.Db).C("users")

	var user User
	err := c.FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		u.logger.WriteLog(err.Error(), log.Info)
	}

	return user, err
}

func (u *MongoService) FindUserByUsername(s *mgo.Session, user User) (User, error) {
	u.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(u.settings.Mongo.Db).C("users")

	err := c.Find(bson.M{"username": user.Username}).One(&user)
	if err != nil {
		u.logger.WriteLog(err.Error(), log.Info)
	}

	return user, err
}

func (u *MongoService) NewUser(s *mgo.Session, user User) error {
	u.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(u.settings.Mongo.Db).C("users")

	_, errf := u.FindUserByUsername(s, user)
	if errf == nil {
		return errf
	}

	hash, _ :=
		func() ([]byte, error) {
			return bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		}()
	user.Password = string(hash)

	err := c.Insert(user)
	if err != nil {
		u.logger.WriteLog(err.Error(), log.Info)
	}

	return err
}

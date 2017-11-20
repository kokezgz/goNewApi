package mongo

import (
	"../utils/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Wallet struct {
	Amount float64 `json:"amount"`
	Rate   string  `json:"rate"`
}

func (u *MongoService) FindWallet(s *mgo.Session, id string) (Wallet, error) {
	u.inject()

	session := s.Copy()
	defer session.Close()

	user, err := u.FindUser(session, id)
	if err != nil {
		u.logger.WriteLog(err.Error(), log.Info)
	}

	wallet := *user.Wallet
	return wallet, err
}

func (u *MongoService) DepositWallet(s *mgo.Session, id string, amount float64) (Wallet, error) {
	u.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(u.settings.Mongo.Db).C("users")

	err := c.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"wallet.amount": amount}})
	wallet, _ := u.FindWallet(session, id)
	return wallet, err
}

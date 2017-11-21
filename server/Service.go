package mongo

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
)

func (u *MongoService) Login(s *mgo.Session, user User) error {
	u.inject()

	finduser, err := u.FindUserByUsername(s, user)
	if err != nil {
		return err
	}

	return func() error {
		return bcrypt.CompareHashAndPassword([]byte(finduser.Password), []byte(user.Password))
	}()
}

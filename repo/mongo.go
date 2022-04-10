package repo

import (
	"atem-stock/configs"
	"fmt"

	"gopkg.in/mgo.v2"
)

type Mg struct {
	Maddr string
}

func NewMongo() *Mg {
	config := configs.AppConfigInstance.MongoConfig
	//"mongodb://root:system@192.168.1.156:27017"
	addr := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.User, config.Pwd, config.Server, config.Port)
	return &Mg{addr}
}

func (mg *Mg) SetAddr(addr string) {
	mg.Maddr = addr
}

func (mg *Mg) Insert(db string, collection string, docs ...interface{}) error {
	session, e := mgo.Dial(mg.Maddr)
	if e != nil {
		return e
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collection)
	err := c.Insert(docs...)
	if err != nil {
		return err
	}
	return nil
}

func (mg *Mg) FindOne(db string, collection string, json map[string]interface{}, result interface{}) error {
	session, e := mgo.Dial(mg.Maddr)
	if e != nil {
		return e
	}
	defer session.Close()
	c := session.DB(db).C(collection)
	c.Find(json).One(result)
	return nil
}

func (mg *Mg) FindSortLimit(db string, collection string, find map[string]interface{}, sort string, limit int, result interface{}) error {
	session, e := mgo.Dial(mg.Maddr)
	if e != nil {
		return e
	}
	defer session.Close()
	c := session.DB(db).C(collection)
	e = c.Find(find).Sort(sort).Limit(limit).One(result)
	if e != nil {
		return e
	}
	return nil
}

func (mg *Mg) FindAll(db string, collection string, json map[string]interface{}, result interface{}) error {
	session, e := mgo.Dial(mg.Maddr)
	if e != nil {
		return e
	}
	defer session.Close()
	c := session.DB(db).C(collection)
	c.Find(json).All(result)
	return nil
}

func (mg *Mg) RemoveAll(db string, collection string) error {
	session, e := mgo.Dial(mg.Maddr)
	if e != nil {
		return e
	}
	defer session.Close()
	c := session.DB(db).C(collection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		return err
	}
	return nil
}

package env

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type serviceAddr struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type Database struct {
	Name       string `json:"name"`
	User       string `json:"user"`
	Password   string `json:"pass"`
	DBName     string `json:"dbname"`
	WithTables bool   `json:"with_tables"`
}

func GetFullAddr(name string) string {
	addr, err := getAddr(name)
	if err != nil {
		panic(err)
	}
	return addr.Addr + ":" + addr.Port
}

func GetPort(name string) string {
	addr, err := getAddr(name)
	if err != nil {
		panic(err)
	}
	return addr.Port
}

func GetAddr(name string) string {
	addr, err := getAddr(name)
	if err != nil {
		panic(err)
	}
	return addr.Addr
}

func getAddrs() ([]serviceAddr, error) {
	raw, err := ioutil.ReadFile("env/addresses.json")
	if err != nil {
		return nil, err
	}
	var c []serviceAddr
	err = json.Unmarshal(raw, &c)
	return c, err
}

func getAddr(name string) (serviceAddr, error) {
	addrs, err := getAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if addr.Name == name {
			return addr, nil
		}
	}
	panic(errors.New("Address not found"))
}

func getDBs() ([]Database, error) {
	raw, err := ioutil.ReadFile("env/databases.json")
	if err != nil {
		return nil, err
	}
	var d []Database
	err = json.Unmarshal(raw, &d)
	return d, err
}

func GetDB(name string) Database {
	dbs, err := getDBs()
	if err != nil {
		panic(err)
	}
	for _, db := range dbs {
		if db.Name == name {
			return db
		}
	}
	panic("No DB exists")
}

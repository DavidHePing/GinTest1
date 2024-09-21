package domain

import "fmt"

type Config struct {
	IsDebug bool
	Server  ServerInfo
	TestDb  DbConnect
}

type ServerInfo struct {
	Port string
}

type DbConnect struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (dbConnect DbConnect) GetGormPostgreConnectString() string {
	return fmt.Sprint("host=", dbConnect.Host, " port=", dbConnect.Port, " user=", dbConnect.User, " password=", dbConnect.Password, " dbname=", dbConnect.Name)
}

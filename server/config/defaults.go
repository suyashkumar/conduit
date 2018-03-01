package config

const LogFile = "LogFile"
const DBConnString = "DBConnString"
const Port = "Port"
const CertKey = "CertKey"
const PrivKey = "PrivKey"
const UseSSL = "UseSSL"

var defaults = map[string]string{
	LogFile:      "logs.txt",
	DBConnString: "",
	Port:         "8000",
	UseSSL:       "false",
}

package config

const APP_PORT = "4000"
const IS_TEST = false    // false:운영, true:테스트
const IS_LOGGING = false // true:로그남김, false:로그없음

var START_ACTIVE = map[string]bool{
	"SSMON": true,
}

var SSMON_DB_CONFIG = map[string]string{
	"CONFIG_NAME":    "SSMON_DB",
	"DBTYPE":         "mysql",
	"HOST":           "localhost",
	"PORT":           "3306",
	"DBNAME":         "ssmon_db",
	"USERNAME":       "ssmon",
	"PASSWORD":       "ssmon123",
	"MAX_IDLE_CONNS": "10",
	"MAX_OPEN_CONNS": "10",
}

var TOKEN_DB_CONFIG = map[string]string{
	"CONFIG_NAME":    "TOKEN_DB",
	"DBTYPE":         "mysql",
	"HOST":           "localhost",
	"PORT":           "3306",
	"DBNAME":         "token_db",
	"USERNAME":       "ssmon",
	"PASSWORD":       "ssmon123",
	"MAX_IDLE_CONNS": "10",
	"MAX_OPEN_CONNS": "10",
}

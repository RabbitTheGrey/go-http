package conf

import "os"

var HttpHost = os.Getenv("HTTP_HOST")
var HttpPort = os.Getenv("HTTP_PORT")

var DbDriver = os.Getenv("DB_DRIVER")
var DbConnectionString = os.Getenv("DB_CONNECTION_STRING")

var JWTSecret = os.Getenv("JWT_SECRET")

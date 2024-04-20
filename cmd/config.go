package main

import "github.com/aadi-1024/ConnectAbility/pkg/database"

type Config struct {
	Db        *database.Database
	JwtSecret []byte
}

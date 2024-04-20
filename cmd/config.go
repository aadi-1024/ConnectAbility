package main

import (
	"github.com/aadi-1024/ConnectAbility/pkg/database"
	"time"
)

type Config struct {
	Db        *database.Database
	JwtSecret []byte
	JwtExpiry time.Duration
}

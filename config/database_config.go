package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	Pguser     string `name:"PGUSER"`
	Pghost     string `name:"PGHOST"`
	Pgport     int    `name:"PGPORT"`
	Pgdatabase string `name:"PGDATABASE"`
	Pgpassword string `name:"PGPASSWORD"`
	Pgsslmode  bool   `name:"PGPREFIX"`
}

func DatabaseConfig() *Database {
	return &Database{
		Pguser:     viper.GetString("PGUSER"),
		Pghost:     viper.GetString("PGHOST"),
		Pgport:     viper.GetInt("PGPORT"),
		Pgdatabase: viper.GetString("PGDATABASE"),
		Pgpassword: viper.GetString("PGPASSWORD"),
		Pgsslmode:  viper.GetBool("PGSSLMODE"),
	}
}

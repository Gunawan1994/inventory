package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DbHost string
	DbUser string
	DbPass string
	DbName string
	DbPort string
}

type Database struct {
	db *gorm.DB
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

func NewDatabase(driver string, cfg *Config) *Database {
	var db *gorm.DB
	var err error
	var dialect gorm.Dialector

	switch driver {
	case "postgres", "pgsql":
		dialect = postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbPort))
	default:
		slog.Warn("unknown database driver")
		os.Exit(1)
	}

	for {
		configGorm := &gorm.Config{
			Logger:          logger.Default.LogMode(logger.Info),
			NowFunc:         time.Now().UTC,
			CreateBatchSize: 1000,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
		db, err = gorm.Open(dialect, configGorm)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to connect to %s database", driver), "error", err.Error())
			slog.Info(fmt.Sprintf("retrying to connect to %s database in 5 seconds...", driver))
			time.Sleep(5 * time.Second)
			continue
		}
		slog.Info(fmt.Sprintf("successfully connected to %s database", driver))
		break
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("failed to configure connection pool", "error", err.Error())
		os.Exit(1)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Hour)

	restore(sqlDB)

	return &Database{db: db}
}

func restore(db *sql.DB) {
	sqlDump := `-- public.article definition

-- Drop table

-- DROP TABLE public.article;

CREATE TABLE public.article (
	id serial4 NOT NULL,
	title varchar NULL,
	"content" varchar NULL,
	author_id int4 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	tag jsonb NULL,
	CONSTRAINT article_pk PRIMARY KEY (id)
);


-- public."user" definition

-- Drop table

-- DROP TABLE public."user";

CREATE TABLE public."user" (
	id serial4 NOT NULL,
	"password" varchar NULL,
	email varchar NULL,
	username varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)
);


-- public."role" definition

-- Drop table

-- DROP TABLE public."role";

CREATE TABLE public."role" (

);
`
	data := strings.Split(sqlDump, ";")
	for _, stmt := range data {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Schema SQL executed successfully")
}

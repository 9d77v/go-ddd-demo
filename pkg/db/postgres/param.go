package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Param struct {
	Host         string `yaml:"host"`
	Port         uint   `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Dbname       string `yaml:"dbname"`
	TablePrefix  string `yaml:"table_prefix"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	Debug        bool   `yaml:"debug"`
}

func (c *Param) New(pgdb *PgDB) (*PgDB, error) {
	fmt.Println("PgDB初始化", c)
	c.createDatabaseIfNotExist()
	client, err := c.newClient()
	if err != nil {
		log.Panicf("Could not initialize gorm: %s\n", err.Error())
	}
	pgdb.db = client
	return pgdb, err
}

func (c *Param) createDatabaseIfNotExist() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable password=%s",
		c.Host, c.Port, c.User, c.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("connect to postgres failed:", err)
	}
	if c.databaseNotExist(db) {
		c.createDatabase(db)
	}
	sqlDBInit, _ := db.DB()
	sqlDBInit.Close()
}

func (c *Param) databaseNotExist(db *gorm.DB) bool {
	var total int64
	err := db.Raw("SELECT 1 FROM pg_database WHERE datname = ?", c.Dbname).Scan(&total).Error
	if err != nil {
		log.Println("check db failed", err)
	}
	return total == 0
}

func (c *Param) createDatabase(db *gorm.DB) {
	initSQL := fmt.Sprintf("CREATE DATABASE \"%s\" WITH  OWNER =%s ENCODING = 'UTF8' CONNECTION LIMIT=-1;",
		c.Dbname, c.User)
	err := db.Exec(initSQL).Error
	if err != nil {
		log.Println("create db failed:", err)
	} else {
		log.Printf("create db '%s' succeed\n", c.Dbname)
	}
}

func (c *Param) newClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		c.Host, c.Port, c.User, c.Dbname, c.Password)
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}
	if c.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.DisableForeignKeyConstraintWhenMigrating = true
	}
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	return db, err
}

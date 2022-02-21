package stores

import (
	"github.com/xuexiangyou/code-art/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := config.Mysql.DataSource
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

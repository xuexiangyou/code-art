package stores

import (
	"context"
	"github.com/xuexiangyou/code-art/config"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(lc fx.Lifecycle, config *config.Config) (*gorm.DB, error) {
	dsn := config.Mysql.DataSource
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			db, err := conn.DB()
			if err != nil {
				return err
			}
			db.Close()
			return nil
		},
	})

	return conn, err
}

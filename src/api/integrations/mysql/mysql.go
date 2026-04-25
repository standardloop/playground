package mysql

import (
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLPoolDestroy() {
	MySQLPool.Close()
}

func MySQLPoolInit() {
	// FIXME, makes variables
	pool, err := sql.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/playground")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// connection pooling
	pool.SetMaxOpenConns(25)
	pool.SetMaxIdleConns(25)
	pool.SetConnMaxLifetime(time.Minute * 5)

	if err := pool.Ping(); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("MySQLPool has been successfully initialized")
	MySQLPool = pool
}

func MySQLPoolPing() error {
	if err := MySQLPool.Ping(); err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}

var MySQLPool *sql.DB = nil

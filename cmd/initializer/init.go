package initializer

import (
	"github.com/go-redis/redis/v8"
	"github.com/linchengzhi/go-clean-backend/Infra/config"
	"github.com/linchengzhi/go-clean-backend/Infra/database"
	"github.com/linchengzhi/go-clean-backend/Infra/logger"
	"github.com/linchengzhi/go-clean-backend/domain/dto"
	"github.com/linchengzhi/go-clean-backend/repository/mysql"
	"github.com/linchengzhi/go-clean-backend/usecase"
	"go.uber.org/zap"
)

type App struct {
	Conf *dto.Config

	Log      *zap.Logger
	MysqlLog *zap.Logger

	MysqlDb *database.Gorm
	RedisDb *redis.Client

	RepoMysql mysql.RepoMysql

	UcAll usecase.UcAll
}

func NewApp(configPath string) (*App, error) {
	app := &App{}
	err := WithConfig(app, configPath)
	if err != nil {
		return nil, err
	}

	err = WithLogger(app)
	if err != nil {
		return nil, err
	}

	err = WithMysql(app)
	if err != nil {
		return nil, err
	}

	err = WithRedis(app)
	if err != nil {
		return nil, err
	}

	WithRepoMysql(app)

	WithUc(app)

	return app, nil
}

func WithConfig(app *App, configPath string) error {
	conf, err := config.NewConfig(configPath)
	if err != nil {
		return err
	}
	app.Conf = conf
	return nil
}

func WithLogger(app *App) error {
	log, err := logger.New(app.Conf.Log)
	if err != nil {
		return err
	}
	app.Log = log

	logConf := app.Conf.Log
	logConf.Level = "info"
	logConf.Filename = "logs/mysql.log"
	mysqlLog, err := logger.New(logConf)
	if err != nil {
		return err
	}
	app.MysqlLog = mysqlLog
	return nil
}

func WithMysql(app *App) error {
	db, err := database.NewGorm(app.Conf.Mysql, app.MysqlLog)
	if err != nil {
		return err
	}
	app.MysqlDb = db
	return nil
}

func WithRedis(app *App) error {
	db, err := database.NewRedis(app.Conf.Redis.Addr, app.Conf.Redis.Password, app.Conf.Redis.Db)
	if err != nil {
		return err
	}
	app.RedisDb = db
	return nil
}

func WithRepoMysql(app *App) {
	repo := mysql.NewRepoMysql(app.MysqlDb.DB)
	app.RepoMysql = repo
}

func WithUc(app *App) {
	us := usecase.NewUcAll(app.MysqlDb, &app.RepoMysql)
	app.UcAll = us
}

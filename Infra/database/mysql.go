package database

import (
	"time"

	"github.com/linchengzhi/go-clean-backend/domain/dto"
	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Gorm struct {
	config dto.Mysql
	log    *zap.Logger
	*gorm.DB
}

func Dsn(cfg dto.Mysql) string {
	return cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Dbname + "?" + cfg.Config
}

func NewGorm(cfg dto.Mysql, logs ...*zap.Logger) (*Gorm, error) {
	mysqlConfig := mysql.Config{
		DSN:                       Dsn(cfg), // DSN data source name
		DefaultStringSize:         255,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，Mysql 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，Mysql 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，Mysql 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据当前 Mysql 版本自动配置
	}

	g := &Gorm{config: cfg}
	if len(logs) > 0 {
		g.log = logs[0]
	}

	//连接数据库
	db, err := gorm.Open(mysql.New(mysqlConfig), getGormConfig(g, cfg.LogLevel))
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}
	// 设置默认值
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifeTime) * time.Second)
	g.DB = db
	return g, nil
}

func getGormConfig(g *Gorm, logMode string) *gorm.Config {
	//禁用外键约束
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if g.log == nil {
		return config
	}

	newLogger := logger.New(
		g, // io writer
		logger.Config{
			SlowThreshold: 10 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Warn,           // log level
			Colorful:      true,                  // 禁用彩色打印
		},
	)

	//设置logger的日志输出等级
	switch logMode {
	case "silent", "Silent", "debug", "Debug":
		config.Logger = newLogger.LogMode(logger.Silent)
	case "cerror", "Error":
		config.Logger = newLogger.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = newLogger.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = newLogger.LogMode(logger.Info)
	default:
		config.Logger = newLogger.LogMode(logger.Info)
	}
	return config
}

//func (g *Gorm) Printf(msg string, args ...interface{}) {
//	if len(args) == 4 {
//		//d := "%s [%.3fms] [row:%d] %s"
//		//g.log.Sugar().Infof(d, args[1:]...)
//		g.log.Info(fmt.Sprintf(msg+"\n", args...))
//	}
//
//	if len(args) == 5 { //出现错误
//		d := "[%.3fms] [row:%d] %s [err:%s]"
//		g.log.Sugar().Infof(d, args[2], args[3], args[4], args[1])
//	}
//	//return
//}

func (g *Gorm) Printf(msg string, args ...interface{}) {
	if len(args) == 4 {
		d := "[%.3fms] [row:%d] %s"
		g.log.Sugar().Infof(d, args[1:]...)
	}

	if len(args) == 5 { //出现错误
		d := "[%.3fms] [row:%v] %s [err:%v]"
		g.log.Sugar().Infof(d, args[2], args[3], args[4], args[1])
	}
	//return
}

// 自动生成表
func AutoMigrate(db *gorm.DB) error {
	db.AutoMigrate(entity.Account{})
	db.AutoMigrate(entity.User{})
	db.AutoMigrate(entity.Article{})
	return nil
}

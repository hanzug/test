package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// compared with use viper each time we need configurations.
// save it in structure is more intuitive.

var Conf = new(AppConfig)

type AppConfig struct {
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

var filePath = "./config.yaml"

func init() {

	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("viper.ReadinConfig failed, err: %v\n", err)
		return
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("configurations was changed")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		}
	})

	Init(Conf.LogConfig, Conf.Mode)
	return
}

var lg *zap.Logger

// InitLogger 初始化Logger
// Init 函数用于初始化日志记录器。
// 参数：
//   - cfg: 日志配置参数，包括文件名、最大大小、最大备份数量、最大保留天数等。
//   - mode: 日志模式，可以是 "dev"（开发模式）或其他（生产模式）。
//
// 返回值：
//   - err: 初始化过程中的错误，如果为 nil 表示初始化成功。
func Init(cfg *LogConfig, mode string) (err error) {
	//fmt.Println("logger init begin") // 打印初始化开始的日志信息

	// 获取日志写入器，根据配置创建一个具有滚动、备份等特性的写入器
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)

	// 获取日志编码器，用于将日志消息编码为字节流
	encoder := getEncoder()

	// 解析配置中的日志级别
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		fmt.Println("logger init failed") // 如果解析日志级别失败，打印初始化失败的日志信息
		return
	}

	var core zapcore.Core

	// 根据模式选择是否在控制台输出日志
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),                                     // 主要的日志记录核心
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 控制台输出核心
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l) // 在生产模式下，只使用主要的日志记录核心
	}

	// 创建最终的 Logger，添加了调用者信息
	lg = zap.New(core, zap.AddCaller())

	// 替换全局日志记录器为当前创建的 Logger
	zap.ReplaceGlobals(lg)

	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

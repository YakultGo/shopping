package config

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	DB struct {
		Mysql string `mapstructure:"mysql"`
		Redis string `mapstructure:"redis"`
	} `mapstructure:"db"`
	User struct {
		Http struct {
			Host string `mapstructure:"host"`
			Port int    `mapstructure:"port"`
		} `mapstructure:"http"`
		Grpc struct {
			Host        string   `mapstructure:"host"`
			ServiceName string   `mapstructure:"service_name"` // 带下划线的字段要是用mapstructure标签
			Tags        []string `mapstructure:"tags"`
		} `mapstructure:"grpc"`
	}
	Consul struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"consul"`
}

var Conf = new(Config)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}
	return nil
}

func InitLogger(name string) {
	writeSyncer := getLogWriter(name)
	encoder := getEncoder()
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.Lock(os.Stdout)), zapcore.DebugLevel,
	)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(name string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./log/%s.log", name),
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

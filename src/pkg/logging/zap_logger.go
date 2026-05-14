package logging

import (
	"golang-clean-arch/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type ZapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func newZapLogger(cfg *config.Config) *ZapLogger {
	logger := &ZapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *ZapLogger) getLogLevel() zapcore.Level {
	level, exists := logLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	} else {
		return level
	}
}

func (l *ZapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.cfg.Logger.FilePath,
		MaxSize:    1,
		MaxAge:     5,
		MaxBackups: 10,
		Compress:   true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	l.logger = logger
}

func prepareLogKeys(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}
	extra["category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}

func (l *ZapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Debugw(msg, params...)
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *ZapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Debugw(msg, params...)
}

func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *ZapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Debugw(msg, params...)
}

func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *ZapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Errorw(msg, params...)
}

func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *ZapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Fatalw(msg, params...)
}

func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

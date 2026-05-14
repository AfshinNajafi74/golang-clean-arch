package logging

import (
	"golang-clean-arch/config"

	"go.uber.org/zap"
)

type ZapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

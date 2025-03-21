package provider

import (
	"context"
	"io"
	"voo.su/internal/config"
	clickhouseModel "voo.su/internal/infrastructure/clickhouse/model"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
	"voo.su/pkg/logger"
)

type LoggerWriter struct {
	Conf             *config.Config
	Writer           io.Writer
	LoggerRepository *clickhouseRepo.LoggerRepository
}

func NewLoggerWriter(
	conf *config.Config,
	writer io.Writer,
	loggerRepository *clickhouseRepo.LoggerRepository,
) *LoggerWriter {
	return &LoggerWriter{
		Conf:             conf,
		Writer:           writer,
		LoggerRepository: loggerRepository,
	}
}

func (c *LoggerWriter) Write(p []byte) (n int, err error) {
	if err := c.LoggerRepository.Create(context.Background(), &clickhouseModel.Logger{
		LogMessage: string(p),
	}); err != nil {
		logger.Errorf("Err logger writer:%s", err)
		return 0, err
	}

	if c.Writer != nil {
		n, err = c.Writer.Write(p)
	}

	return n, err
}

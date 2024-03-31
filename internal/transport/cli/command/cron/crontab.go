package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	crontab "voo.su/internal/transport/cli/handle/cron"
)

type Command *cli.Command

type ICrontab interface {
	Spec() string
	Enable() bool
	Handle(ctx context.Context) error
}

type Subcommands struct {
	ClearWsCache      *crontab.ClearWsCache
	ClearTmpFile      *crontab.ClearTmpFile
	ClearExpireServer *crontab.ClearExpireServer
}

func NewCrontabCommand(handles *Subcommands) Command {
	return &cli.Command{
		Name:  "crontab",
		Usage: "Команда Crontab - постоянные периодические задачи",
		Action: func(ctx *cli.Context) error {
			c := cron.New()
			for _, exec := range toCrontab(handles) {
				job := exec
				_, _ = c.AddFunc(job.Spec(), func() {
					defer func() {
						if err := recover(); err != nil {
							log.Printf("Ошибка Crontab: %v \n", err)
						}
					}()
					_ = job.Handle(ctx.Context)
				})
				log.Printf("Запущена постоянная задача %T => Планировщик задач %s \n", job, job.Spec())
			}
			log.Println("Crontab периодические задачи запущены...")
			return run(c, ctx.Context)
		},
	}
}

func run(cron *cron.Cron, ctx context.Context) error {
	cron.Start()
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	select {
	case <-s:
		cron.Stop()
	case <-ctx.Done():
		cron.Stop()
	}

	log.Println("Crontab периодические задачи остановлены")

	return nil
}

func toCrontab(value any) []ICrontab {
	var jobs []ICrontab
	elem := reflect.ValueOf(value).Elem()
	for i := 0; i < elem.NumField(); i++ {
		if v, ok := elem.Field(i).Interface().(ICrontab); ok {
			if v.Enable() {
				jobs = append(jobs, v)
			}
		}
	}

	return jobs
}

package cli

import (
	"context"
	"github.com/jedib0t/go-pretty/v6/table"
	crontab "github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"voo.su/internal/config"
	"voo.su/internal/transport/cli/handle/cron"
)

type ICrontab interface {
	Name() string
	Spec() string
	Enable() bool
	Handle(ctx context.Context) error
}

type Crontab struct {
	ClearWsCache      *cron.ClearWsCache
	ClearTmpFile      *cron.ClearTmpFile
	ClearExpireServer *cron.ClearExpireServer
}

type CronProvider struct {
	Conf    *config.Config
	Crontab *Crontab
}

func Cron(ctx *cli.Context, app *CronProvider) error {
	c := crontab.New()

	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{"#", "Name", "Time"})

	for i, exec := range toCrontab(app.Crontab) {
		job := exec

		_, _ = c.AddFunc(job.Spec(), func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Ошибка Crontab: %v \n", err)
				}
			}()

			_ = job.Handle(ctx.Context)
		})

		tbl.AppendRow([]any{i + 1, job.Name(), job.Spec()})
	}

	tbl.Render()

	return run(c, ctx.Context)
}

func run(cron *crontab.Cron, ctx context.Context) error {
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

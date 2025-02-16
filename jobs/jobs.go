package jobs

import (
	"github.com/openuniland/good-guy/configs"
	"github.com/openuniland/good-guy/external/facebook"
	"github.com/openuniland/good-guy/internal/articles"
	"github.com/openuniland/good-guy/internal/users"
	"github.com/robfig/cron/v3"
)

type Jobs struct {
	cfg        *configs.Configs
	articleUC  articles.UseCase
	userUC     users.UseCase
	facebookUC facebook.UseCase
}

func NewJobs(cfg *configs.Configs, articleUC articles.UseCase, userUC users.UseCase, facebookUC facebook.UseCase) *Jobs {
	return &Jobs{cfg: cfg, articleUC: articleUC, userUC: userUC, facebookUC: facebookUC}
}

func (j *Jobs) Run() {

	c := cron.New(cron.WithSeconds())

	//every 25 minutes
	c.AddFunc("*/25 * * * *", func() {
		go j.SyncArticles()
	})

	c.Start()

}

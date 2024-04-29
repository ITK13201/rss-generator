package jobs

import (
	"context"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/testfeed"
	"github.com/sirupsen/logrus"
	"time"
)

type DeleteExpiredTestFeedsJob interface {
	Run(ctx context.Context)
}

type deleteExpiredTestFeedsJob struct {
	cfg       *domain.Config
	logger    *logrus.Logger
	sqlClient *ent.Client
}

func NewDeleteExpiredTestFeedsJob(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) DeleteExpiredTestFeedsJob {
	return &deleteExpiredTestFeedsJob{
		cfg:       cfg,
		logger:    logger,
		sqlClient: sqlClient,
	}
}

func (dj *deleteExpiredTestFeedsJob) Run(ctx context.Context) {
	dj.logger.WithFields(logrus.Fields{
		"job": "delete_expired_test_feeds",
	}).Info("job started")

	now := time.Now()
	expired_date := now.Add(-7 * 24 * time.Hour)
	_, err := dj.sqlClient.TestFeed.Delete().Where(testfeed.CreatedAtLT(expired_date)).Exec(ctx)
	if err != nil {
		dj.logger.WithFields(logrus.Fields{
			"job":   "delete_expired_test_feeds",
			"error": err,
		}).Fatal("job failed")
	}

	dj.logger.WithFields(logrus.Fields{
		"job": "delete_expired_test_feeds",
	}).Info("job finished")
}

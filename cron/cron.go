package cron

import (
	"github.com/cenkalti/backoff"
	"github.com/prolifera/kit/log"
	"github.com/robfig/cron"
	"reflect"
	"time"
)

const (
	initialInterval = time.Second * 3
	maxRetryTime    = 3
)

const (
	JobDurationHourly   = "@hourly"
	JobDurationDaily    = "@daily"
	JobDurationMinutely = "0 */1 * * *"
)

func AddCronJob(cronJob func() error, jobDuration string, preHandle bool) error {

	if preHandle {
		if err := cronJob(); err != nil {
			log.Error().Msgf("cronjob activity %s err %s", reflect.TypeOf(cronJob).Name(), err.Error())
			return err
		}
	}

	c := cron.New()
	err := c.AddFunc(jobDuration, func() {
		if err := cronJob(); err != nil {
			log.Error().Msgf("cronjob activity %s error %s", reflect.TypeOf(cronJob), err.Error())
		}
	})
	if err != nil {
		log.Error().Msgf("add cronjob activity error %s, check your input ", err.Error())
		return err
	}

	c.Start()
	return nil
}

func AddCronJobWithBackoff(cronJob func() error, jobDuration string) error {

	if err := cronJob(); err != nil {
		log.Error().Msgf("cronjob activity %s err %s", reflect.TypeOf(cronJob()).Name(), err.Error())
		return err
	}

	c := cron.New()
	err := c.AddFunc(jobDuration, func() {
		backOffHandler(cronJob)
	})
	if err != nil {
		return err
	}

	c.Start()
	return nil
}

func backOffHandler(cronJob func() error) {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = initialInterval
	err := backoff.Retry(cronJob, backoff.WithMaxRetries(b, maxRetryTime))
	if err != nil {
		log.Error().Msgf("cronjob activity %s  backoff handler after retry %d times, got error %s", reflect.TypeOf(cronJob), maxRetryTime, err.Error())
	} else {
		log.Info().Msgf("cronjob activity %s backoff handler success", reflect.TypeOf(cronJob))
	}
}

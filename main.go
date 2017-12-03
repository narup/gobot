package main

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
)

var jobsQueue chan string

var (
	consumerKey       = getenv("BD_TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("BD_TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("BD_TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("BD_TWITTER_ACCESS_TOKEN_SECRET")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func main() {

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#uiux", "#ui", "#ux", "#materialdesign", "#googledesign", "#dribbble"},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Warningf("received unexpected value of type %T", v)
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}

		_, err := api.Retweet(t.Id, false)
		if err != nil {
			log.Errorf("could not retweet %d: %v", t.Id, err)
			continue
		}
	}

	defer close(jobsQueue)
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }

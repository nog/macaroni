package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fujiwara/macaroni"
)

func main() {
	conf := buildConf()
	err := macaroni.Run(conf, os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func buildConf() *macaroni.Config {
	// ignore error because default false
	muteOnNormal, _ := strconv.ParseBool(os.Getenv("SLACK_MUTE_ON_NORMAL"))

	var prefix string
	if prefix = os.Getenv("MACKEREL_METRIC_NAME_PREFIX"); prefix == "" {
		prefix = macaroni.DefaultMetricNamePrefix
	}

	conf := &macaroni.Config{
		Mackerel: &macaroni.MackerelConfig{
			ApiKey:           os.Getenv("MACKEREL_APIKEY"),
			Service:          os.Getenv("MACKEREL_SERVICE"),
			MetricNamePrefix: prefix,
			MetricName:       os.Getenv("MACKEREL_METRIC_NAME"),
		},
		Slack: &macaroni.SlackConfig{
			Endpoint:     os.Getenv("SLACK_ENDPOINT"),
			Channel:      os.Getenv("SLACK_CHANNEL"),
			Username:     os.Getenv("SLACK_USERNAME"),
			IconEmoji:    os.Getenv("SLACK_ICON_EMOJI"),
			Mention:      os.Getenv("SLACK_MENTION"),
			PasteBinCmd:  os.Getenv("SLACK_PASTEBIN_CMD"),
			MuteOnNormal: muteOnNormal,
		},
	}
	return conf
}

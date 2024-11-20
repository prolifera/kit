package xslack

import (
	"fmt"
	"github.com/prolifera/kit/log"
	"net/http"
	"net/url"
	"os"
)

var slackWebhookURL string
var defaultUsername string
var defaultChannel string

func init() {
	slackWebhookURL = os.Getenv("SLACK_WEBHOOK")
	defaultUsername = os.Getenv("SLACK_USERNAME")
	defaultChannel = os.Getenv("SLACK_CHANNEL")
}

func SendSlackDefault(message string) {
	go func() {
		err := SendSlackNotification("", "", message)
		if err != nil {
			log.Error().Fields(map[string]interface{}{"action": "send slack error", "error": err.Error()}).Send()
		}
	}()
}

// SendSlackNotification 发送 Slack 消息
func SendSlackNotification(channel, username, message string) error {

	message = "----------\n" + message
	message += "\n"

	if slackWebhookURL == "" {
		return nil
	}

	if username == "" {
		username = defaultUsername
	}
	if channel == "" {
		channel = defaultChannel
	}

	payload := fmt.Sprintf(
		`{"channel": "%s", "username": "%s", "text": "%s", "icon_emoji": ""}`,
		channel, username, message,
	)
	data := url.Values{}
	data.Set("payload", payload)

	resp, err := http.PostForm(slackWebhookURL, data)
	if err != nil {
		return fmt.Errorf("error sending slack notification: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

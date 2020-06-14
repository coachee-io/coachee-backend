package slack_test

import (
	"coachee-backend/pkg/slack"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessenger_Post(t *testing.T) {
	messenger := slack.NewMessenger("https://hooks.slack.com/services/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")

	simpleMessage := slack.SimpleMessage("This is a test message")

	err := messenger.Post(simpleMessage)
	require.Nil(t, err)
}

package plugin

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/Brightscout/mattermost-plugin-servicenow-virtual-agent/server/serializer"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

func Test_MessageHasBeenPosted(t *testing.T) {
	defer monkey.UnpatchAll()

	for _, testCase := range []struct {
		description                       string
		Message                           string
		getChannelError                   *model.AppError
		getUserError                      error
		parseAuthTOkenError               error
		sendMessageToVirtualAgentAPIError error
	}{
		{
			description: "Message is posted and successfully sent to Virtual Agent",
			Message:     "mockMessage",
		},
		{
			description:     "Message is posted and failed to get current channel",
			getChannelError: &model.AppError{},
			Message:         "mockMessage",
		},
		{
			description:  "Message is posted and failed to get user from KV store",
			getUserError: errors.New("mockError"),
			Message:      "mockMessage",
		},
		{
			description:  "Message is posted and user is not connected to ServiceNow",
			getUserError: ErrNotFound,
			Message:      "mockMessage",
		},
		{
			description: "Message is posted and user is not connected to ServiceNow",
			Message:     "disconnect",
		},
		{
			description:         "Message is posted and failed to parse auth token",
			parseAuthTOkenError: errors.New("mockError"),
			Message:             "mockMessage",
		},
		{
			description:                       "Message is posted and failed to parse auth token",
			sendMessageToVirtualAgentAPIError: errors.New("mockError"),
			Message:                           "mockMessage",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			p := Plugin{}

			p.botUserID = "mock-botID"
			mockAPI := &plugintest.API{}

			mockAPI.On("LogError", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("LogError error")

			mockAPI.On("GetChannel", "mockChannelID").Return(&model.Channel{
				Type: "D",
				Name: "mock-botID__mock",
			}, testCase.getChannelError)

			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "Ephemeral", func(_ *Plugin, _, _, _ string, _ ...interface{}) {})

			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "GetUser", func(_ *Plugin, _ string) (*serializer.User, error) {
				return &serializer.User{}, testCase.getUserError
			})

			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "DM", func(_ *Plugin, _, _ string, _ ...interface{}) (string, error) {
				return "mockPostID", nil
			})

			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "DMWithAttachments", func(_ *Plugin, _ string, _ ...*model.SlackAttachment) (string, error) {
				return "mockPostID", nil
			})

			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "ParseAuthToken", func(_ *Plugin, _ string) (*oauth2.Token, error) {
				return &oauth2.Token{}, testCase.parseAuthTOkenError
			})

			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "MakeClient", func(_ *Plugin, _ context.Context, _ *oauth2.Token) Client {
				return &client{}
			})

			monkey.PatchInstanceMethod(reflect.TypeOf(&client{}), "SendMessageToVirtualAgentAPI", func(_ *client, _, _ string, _ bool) error {
				return testCase.sendMessageToVirtualAgentAPIError
			})

			p.SetAPI(mockAPI)

			post := &model.Post{
				ChannelId: "mockChannelID",
				UserId:    "mock-userID",
				Message:   testCase.Message,
			}

			p.MessageHasBeenPosted(&plugin.Context{}, post)
		})
	}
}
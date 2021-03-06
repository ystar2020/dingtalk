package cmd

import (
	"github.com/ystar2020/dingtalk/client"
	"github.com/ystar2020/dingtalk/message"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var feedCardCmd = &cobra.Command{
	Use:   "feedCard",
	Short: "send feedCard message with DingTalk robot",
	Long:  `send feedCard message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		if !CheckToken() {
			log.L(log.Red, "access_token can not be empty")
			return
		}

		if len(feedCardVars.titles) < 1 || len(feedCardVars.picURLs) < 1 || len(feedCardVars.messageURLs) < 1 {
			log.L(log.Red, "titles & picURLs & messageURLs can not be empty")
			return
		}

		if len(feedCardVars.titles) == len(feedCardVars.picURLs) && len(feedCardVars.picURLs) == len(feedCardVars.messageURLs) {
			dingTalk := client.DingTalk{
				AccessToken: rootVars.accessToken,
				Secret:      rootVars.secret,
			}

			msg := message.NewFeedCardMessage()
			for i := 0; i < len(feedCardVars.titles); i++ {
				msg.AppendLink(feedCardVars.titles[i], feedCardVars.messageURLs[i], feedCardVars.picURLs[i])
			}
			if _, err := dingTalk.Send(msg); err != nil {
				log.L(log.Red, err.Error())
			}
		} else {
			log.L(log.Red, "titles & picURLs & messageURLs count must be equal")
			return
		}
	},
}

// FeedCardVars struct
type FeedCardVars struct {
	titles      []string
	picURLs     []string
	messageURLs []string
}

var feedCardVars FeedCardVars

func init() {
	rootCmd.AddCommand(feedCardCmd)

	feedCardCmd.Flags().StringSliceVarP(&feedCardVars.titles, "titles", "i", []string{}, "titles")
	feedCardCmd.Flags().StringSliceVarP(&feedCardVars.picURLs, "picURLs", "p", []string{}, "picURLs")
	feedCardCmd.Flags().StringSliceVarP(&feedCardVars.messageURLs, "messageURLs", "u", []string{}, "messageURLs")
}

package cmd

import (
	"github.com/ystar2020/dingtalk/client"
	"github.com/ystar2020/dingtalk/message"
	"github.com/CatchZeng/gutils/log"
	"github.com/spf13/cobra"
)

var actionCardCmd = &cobra.Command{
	Use:   "actionCard",
	Short: "send actionCard message with DingTalk robot",
	Long:  `send actionCard message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		if !CheckToken() {
			log.L(log.Red, "access_token can not be empty")
			return
		}

		if len(actionCardVars.Title) < 1 {
			log.L(log.Red, "title can not be empty")
			return
		}

		if len(actionCardVars.Text) < 1 {
			log.L(log.Red, "text can not be empty")
			return
		}

		var isOverallJump = false
		if len(actionCardVars.SingleTitle) < 1 {
			if len(btnTitles) < 1 {
				log.L(log.Red, "btns can not be empty when singleTitle is empty")
				return
			}
		} else {
			isOverallJump = true
			if len(actionCardVars.SingleURL) < 1 {
				log.L(log.Red, "singleURL can not be empty")
				return
			}
		}

		dingTalk := client.DingTalk{
			AccessToken: rootVars.accessToken,
			Secret:      rootVars.secret,
		}
		msg := message.NewActionCardMessage()
		if isOverallJump {
			msg.SetOverallJump(
				actionCardVars.Title,
				actionCardVars.Text,
				actionCardVars.SingleTitle,
				actionCardVars.SingleURL,
				actionCardVars.BtnOrientation,
				actionCardVars.HideAvatar)
		} else {
			if len(btnTitles) != len(btnActionURLs) {
				log.L(log.Red, "btnTitles & btnActionURLs count must be equal")
				return
			}
			for i := 0; i < len(btnTitles); i++ {
				actionCardVars.Btns = append(actionCardVars.Btns, message.Btn{
					Title:     btnTitles[i],
					ActionURL: btnActionURLs[i],
				})
			}
			msg.SetIndependentJump(
				actionCardVars.Title,
				actionCardVars.Text,
				actionCardVars.Btns,
				actionCardVars.BtnOrientation,
				actionCardVars.HideAvatar)
		}
		if _, err := dingTalk.Send(msg); err != nil {
			log.L(log.Red, err.Error())
		}
	},
}

var actionCardVars message.ActionCard
var btnTitles, btnActionURLs []string

func init() {
	rootCmd.AddCommand(actionCardCmd)
	actionCardCmd.Flags().StringVarP(&actionCardVars.Title, "title", "i", "", "title")
	actionCardCmd.Flags().StringVarP(&actionCardVars.Text, "text", "e", "", "text")
	actionCardCmd.Flags().StringVarP(&actionCardVars.SingleTitle, "singleTitle", "n", "", "singleTitle")
	actionCardCmd.Flags().StringVarP(&actionCardVars.SingleURL, "singleURL", "u", "", "singleURL")
	actionCardCmd.Flags().StringSliceVarP(&btnTitles, "btnTitles", "b", []string{}, "btnTitles")
	actionCardCmd.Flags().StringSliceVarP(&btnActionURLs, "btnActionURLs", "c", []string{}, "btnActionURLs")
	actionCardCmd.Flags().StringVarP(&actionCardVars.BtnOrientation, "btnOrientation", "o", "", "btnOrientation")
	actionCardCmd.Flags().StringVarP(&actionCardVars.HideAvatar, "hideAvatar", "d", "", "hideAvatar")
}

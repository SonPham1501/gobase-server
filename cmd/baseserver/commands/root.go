package commands

import (
	"github.com/spf13/cobra"
)

type Command = cobra.Command

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "gobase",
	Short: "Open source, self-hosted Slack-alternative",
	Long:  `Gobase offers workplace messaging across web, PC and phones with archiving, search and integration with your existing systems. Documentation available at https://docs.gobase.com`,
}

func init() {
	RootCmd.PersistentFlags().StringP("config", "c", "", "Configution file to use")
}

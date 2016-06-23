package cmd

import (
	"github.com/spf13/cobra"
)

func addSubCommands(cmd *cobra.Command) {
	cmd.AddCommand(NewCmdDescription())
	cmd.AddCommand(NewCmdLocation())
	cmd.AddCommand(NewCmdPrint())
	cmd.AddCommand(NewCmdAddingCity())
}

func NewCmd() *cobra.Command  {
	cmd := &cobra.Command{
		Use: "sun",
		Short: "Md. Abdulla-Al-Sun",
		Long: "Go to https://sites.google.com/site/sunkuet02/",
		Example: "sun description",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	addSubCommands(cmd)
	return cmd
}
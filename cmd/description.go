package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdDescription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe (school|college)",
		Short: "Details about Sun",
		Long:  `Read Details from his website`,
	}
	cmd.AddCommand(NewCmdSchool())
	cmd.AddCommand(NewCmdCollege())

	return cmd
}
package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)


func NewCmdLocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "location",
		Short: "Location of about Sun",
		Long:  `Parmanent and Current Location of Sun`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("He is in Dhaka and from Rangpur")
		},
	}
	return cmd
}
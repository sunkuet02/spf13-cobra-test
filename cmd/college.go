package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)


func NewCmdCollege() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "college",
		Short: "The name of Collge, Sun read",
		Long:  `Name and Location of Sun's College`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Police Lines School and College, Rangpur")
		},
	}
	return cmd
}
package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)


func NewCmdSchool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "school",
		Short: "The name of school, Sun read",
		Long:  `Name and Location of Sun's School`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Dhantala R.U School and College, Gangachara, Rangpur")
		},
	}
	return cmd
}
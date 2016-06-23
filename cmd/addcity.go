package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)


func NewCmdAddingCity() *cobra.Command {
	var homeCity, currentCity string
	cmd := &cobra.Command{
		Use:   "addcity --home=[home city] --current=[current city]",
		Short: "Take two city , return with fomating names",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Home city : ", homeCity, "\n", "Current city :", currentCity)
		},
	}
	cmd.Flags().StringVar(&homeCity, "home", "", "Home city")
	cmd.Flags().StringVar(&currentCity, "current", "", "Current city")
	return cmd
}
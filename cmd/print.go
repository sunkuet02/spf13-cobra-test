package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
)


func NewCmdPrint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "print [strings to print]",
		Short: "Take two city , return with fomating names",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print : ", strings.Join(args, " "))
		},
	}
	return cmd
}
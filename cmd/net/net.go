/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"github.com/spf13/cobra"
)

// netCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "network based commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}

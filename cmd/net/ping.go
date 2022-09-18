/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/spf13/cobra"
)

var (
	showResp bool
	client   = http.Client{Timeout: time.Second * 2}
)

func ping(domain string) (string, error) {
	url := "https://" + domain

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	dumpedResp, err := httputil.DumpResponse(resp, showResp)
	if err != nil {
		return "", err
	}

	return string(dumpedResp), nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping [url(s) to ping]",
	Short: "ping websites",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if resp, err := ping(args[0]); err != nil {
				log.Panic(err)
			} else {
				fmt.Println(resp)
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	NetCmd.AddCommand(pingCmd)

	pingCmd.Flags().BoolVar(&showResp, "showResp", false, "show response")
}

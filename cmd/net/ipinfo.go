/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SwarnimWalavalkar/toolkit/util"
	"github.com/spf13/cobra"
)

type IpInfoResp struct {
	IP       string `json:"ip" header:"IP"`
	Org      string `json:"org" header:"Org"`
	Hostname string `json:"hostname" header:"Hostname"`
	City     string `json:"city" header:"City"`
	Region   string `json:"region" header:"Region"`
	Country  string `json:"country" header:"Country"`
	Loc      string `json:"loc" header:"Location"`
	Timezone string `json:"timezone" header:"Timezone"`
	Postal   string `json:"postal" header:"Postal"`
}

func fetchData(ip string) IpInfoResp {
	url := "http://ipinfo.io/" + ip + "/"
	client = http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	data := IpInfoResp{}

	err = json.Unmarshal(respByte, &data)
	if err != nil {
		log.Panic(err)
	}

	return data
}

// pingCmd represents the ping command
var ipInfoCmd = &cobra.Command{
	Use:   "ipinfo",
	Short: "fetch ip info ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Printf("\n\n---Fetching Data For %s---\n\n", ip)
				data := fetchData(ip)
				util.DisplayTable(data)
			}
		} else {
			fmt.Println("Please provide an IP address")
		}
	},
}

func init() {
	NetCmd.AddCommand(ipInfoCmd)

}

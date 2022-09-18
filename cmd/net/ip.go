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

	"github.com/spf13/cobra"
)

type IP struct {
    Query string `json:"query"`
}

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get your public IP",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    req, err := http.Get("http://ip-api.com/json/"); if err != nil { log.Panic(err) }

    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body); if err != nil { log.Panic(err) }

    var ip IP
    json.Unmarshal(body, &ip)

    fmt.Println(ip.Query)
	},
}

func init() {
	NetCmd.AddCommand(ipCmd)
}

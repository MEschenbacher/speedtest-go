package main

import (
	speedtest "github.com/meschenbacher/speedtest-go"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	showList   = kingpin.Flag("list", "Show available speedtest.net servers.").Short('l').Bool()
	serverIds  = kingpin.Flag("server", "Select server id to speedtest.").Short('s').Ints()
	savingMode = kingpin.Flag("saving-mode", "Using less memory (≒10MB), though low accuracy (especially > 30Mbps).").Bool()
)

func main() {
	kingpin.Version("1.1.1")
	kingpin.Parse()

	speedtester := speedtest.New()
	speedtester.FetchServers()

	speedtester.ShowUser()

	if *showList {
		speedtester.ShowList()
		return
	}
	speedtester.ShowResult(*serverIds)
}

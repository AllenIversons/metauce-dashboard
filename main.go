package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	db2 "metauce-dashboard/db"
	"metauce-dashboard/types"
	"os"
	"os/signal"
	"syscall"
)

var (
	OriginCommandHelpTemplate = `{{.Name}}{{if .Subcommands}} command{{end}}{{if .Flags}} [command options]{{end}} {{.ArgsUsage}}
{{if .Description}}{{.Description}}
{{end}}{{if .Subcommands}}
SUBCOMMANDS:
  {{range .Subcommands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}{{end}}{{if .Flags}}
OPTIONS:
{{range $.Flags}}   {{.}}
{{end}}
{{end}}`
)

var app *cli.App

var Duration int = 0
var MetisUrl = "https://stardust.metis.io/?owner=588"
var Port = "8547"
var ContractAddressStrList = []string{"", ""}
var (
	portFlag = cli.StringFlag{
		Name:  "port",
		Usage: "restful rpc port",
		Value: "8547",
	}
	configFlag = cli.StringFlag{
		Name:  "config",
		Usage: "config file path",
		Value: "./config.yml",
	}
	durationFlag = cli.IntFlag{
		Name:  "dur",
		Usage: "track cycle duration",
		Value: 1440,
	}
)

type Tracker interface {
	start() error
	close() error
	GetRankingList() ([]types.AddressRank, error)
	GetAddressRank(address common.Address) (types.AddressRank, error)
}

type ConfigYml struct {
	Url             string   `yaml:"url"`
	Duration        int      `yaml:"duration"`
	Port            string   `yaml:"port"`
	ContractAddress []string `yaml:"contract_address"`
}

func init() {
	app = cli.NewApp()
	app.Version = "v1.0.0"
	app.Action = Start
	app.Flags = []cli.Flag{
		portFlag,
		configFlag,
		durationFlag,
	}

	cli.CommandHelpTemplate = OriginCommandHelpTemplate
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Start(ctx cli.Context) {
	contractAddressList := make([]common.Address, 0)
	configPath := ctx.String(configFlag.Name)
	b, err := ioutil.ReadFile(configPath)
	if err != nil && ctx.IsSet(configFlag.Name) {
		utils.Fatalf("read config file error %v", err)
	}
	if err == nil {
		var configYml ConfigYml
		e := yaml.Unmarshal(b, &configYml)
		if e != nil && ctx.IsSet(configFlag.Name) {
			utils.Fatalf("unmarshal config error %v", err)
		}

		if configYml.Url != "" {
			MetisUrl = configYml.Url
		}

		if configYml.Port != "" {
			Port = configYml.Port
		}

		if configYml.Duration != 0 {
			Duration = configYml.Duration
		}

		if len(configYml.ContractAddress) != 0 {
			for _, addrStr := range configYml.ContractAddress {
				contractAddressList = append(contractAddressList, common.HexToAddress(addrStr))
			}
		}
	}

	if ctx.IsSet(portFlag.Name) {
		Port = ctx.String(portFlag.Name)
	}

	if ctx.IsSet(durationFlag.Name) {
		Duration = ctx.Int(durationFlag.Name)
	}

	if len(contractAddressList) == 0 {
		for _, addrStr := range ContractAddressStrList {
			contractAddressList = append(contractAddressList, common.HexToAddress(addrStr))
		}
	}

	//init DB service
	dbService, err := db2.Init(dbpassword, dburl)
	if err != nil {
		utils.Fatalf("init db service error %v", err)
	}
	defer dbService.Close()

	//init start tracker

	//init and start restful rpc

	waitToExit()
}

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	if !signal.Ignored(syscall.SIGHUP) {
		signal.Notify(sc, syscall.SIGHUP)
	}
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range sc {
			fmt.Printf("received exit signal:%v", sig.String())
			close(exit)
			break
		}
	}()
	<-exit
}

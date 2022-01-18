package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"metauce-dashboard/consts"
	db2 "metauce-dashboard/db"
	"metauce-dashboard/restful"
	tracker2 "metauce-dashboard/tracker"
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
var CodeContractStr = "0xE5547D0378944B786BB38fB1bea4027052FC883E"
var PlayContractStr = "0x7Ab359a33F14C23571f79b62f1B8a5dc510c8070"

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
	codeFlag = cli.StringFlag{
		Name:  "code",
		Usage: "code contract address",
		Value: "0xE5547D0378944B786BB38fB1bea4027052FC883E",
	}
	playFlag = cli.StringFlag{
		Name:  "play",
		Usage: "play contract address",
		Value: "0x7Ab359a33F14C23571f79b62f1B8a5dc510c8070",
	}
)

type ConfigYml struct {
	Url          string `yaml:"url"`
	Duration     int    `yaml:"duration"`
	Port         string `yaml:"port"`
	CodeContract string `yaml:"code_contract"`
	PlayContract string `yaml:"play_contract"`
}

func init() {
	app = cli.NewApp()
	app.Version = "v1.0.0"
	app.Action = Start
	app.Flags = []cli.Flag{
		portFlag,
		configFlag,
		durationFlag,
		codeFlag,
		playFlag,
	}

	cli.CommandHelpTemplate = OriginCommandHelpTemplate
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Start(ctx *cli.Context) {
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

		if configYml.CodeContract != "" {
			CodeContractStr = configYml.CodeContract
		}

		if configYml.PlayContract != "" {
			PlayContractStr = configYml.PlayContract
		}
	}

	if ctx.IsSet(portFlag.Name) {
		Port = ctx.String(portFlag.Name)
	}

	if ctx.IsSet(durationFlag.Name) {
		Duration = ctx.Int(durationFlag.Name)
	}

	if ctx.IsSet(codeFlag.Name) {
		CodeContractStr = ctx.String(codeFlag.Name)
	}

	if ctx.IsSet(playFlag.Name) {
		PlayContractStr = ctx.String(playFlag.Name)
	}

	//init DB service
	dbService, err := db2.Init(consts.Dbpassword, consts.Dburl)
	if err != nil {
		utils.Fatalf("init db service error %v", err)
	}
	defer dbService.Close()

	//init start tracker
	codeContract := common.HexToAddress(CodeContractStr)
	playContract := common.HexToAddress(PlayContractStr)
	tracker := tracker2.Init(dbService, MetisUrl, playContract, codeContract, Duration)
	err = tracker.Start()
	if err != nil {
		utils.Fatalf("start tracker error %v", err)
	}
	defer tracker.Close()

	//init and start restful rpc
	rpcServer := restful.Init(Port, tracker)
	err = rpcServer.Start()
	if err != nil {
		utils.Fatalf("start rpc server error %v", err)
	}
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

package main

import (
	"context"
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/walletconnect"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

const (
	Version      = "1.0.2"
	PROGRAM_NAME = "dapp-server"
)

const(
	CMD_FLAG_NAME_SUB_TOPIC = "topic"
)

var (
	BuildTime = "2022-02-23"
	GitCommit = ""
)

func grace() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	go func() {
		for {
			select {
			case s := <-sigChannel:
				{
					if s != nil && s == os.Interrupt {
						fmt.Printf("Ctrl+C signal captured, program exiting...\n")
						close(sigChannel)
						os.Exit(0)
					}
				}
			}
		}
	}()
}

func init() {
	log.SetLevel("debug")
}

func main() {
	grace()

	app := &cli.App{
		Name:     PROGRAM_NAME,
		Usage:    "[options] wc-uri",
		Version:  fmt.Sprintf("v%s %s commit %s", Version, BuildTime, GitCommit),
		Flags:    []cli.Flag{
			&cli.StringFlag{
				Name:  CMD_FLAG_NAME_SUB_TOPIC,
				Usage: "topic to subscribe",
				Required: true,
			},
		},
		Commands: nil,
		Action:  func(cctx *cli.Context) error {
			if cctx.Args().First() == "" {
				return log.Errorf("wc uri required")
			}
			dapp, err := walletconnect.NewDApp(cctx.Args().First())
			if err != nil {
				return log.Errorf("Dapp instance is nil")
			}
			topic := cctx.String(CMD_FLAG_NAME_SUB_TOPIC)
			err = dapp.SubscribeTopic(topic, Subscriber)
			if err != nil {
				return log.Errorf("subscribe error [%s]", err)
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
	//strWCURI := "wc:38d80e38-1c9a-470f-aed1-27b8ced3338d@1?bridge=https://6.bridge.walletconnect.org&key=d8089b579aeef1ba4692d72bd3a0bb870f8bbceafaab325f23739635e2e54f05"
	//qrcode.QrencodeTerminal(strWCURI, qrencode.ECLevelL)
	ch := make(chan interface{})
	<-ch
	log.Infof("program exit...")
}

func Subscriber(ctx context.Context, topic, typ string,  msg []byte) bool {
	log.Infof("message [%s]", msg)
	return true
}
package cli

import (
	"github.com/alecthomas/kong"
	"github.com/blue-monads/lpweb/code/proxy"
	"github.com/blue-monads/lpweb/code/tunnel"
	"github.com/k0kubun/pp"
)

type CLI struct {
	WebProxy struct {
		Port int `help:"port run proxy on"`
	} `cmd:"" help:"web proxy for forwrading local requests to target libweb server"`

	HttpTunnel struct {
		Port int `help:"port to tunnel"`
	} `cmd:"" help:"http tunnel to a http service running in local port"`

	Key string

	ctx *kong.Context
}

func RunCLI() {
	cli := &CLI{}
	cli.ctx = kong.Parse(cli)

	switch cli.ctx.Command() {
	case "web-proxy":
		cli.runWebProxy()
	case "http-tunnel":
		cli.runHttpTunnel()
	default:
		panic("Not implemented" + cli.ctx.Command())
	}

}

func (c *CLI) runWebProxy() {
	pp.Println("@start webproxy")
	if c.WebProxy.Port == 0 {
		c.WebProxy.Port = 8080
	}

	wproxy := proxy.New(c.WebProxy.Port)

	pp.Println("@run", wproxy.Run())

}

func (c *CLI) runHttpTunnel() {
	pp.Println("@start http tunnel")
	if c.HttpTunnel.Port == 0 {
		panic("tunnel port is needed")
	}

	htun := tunnel.New(c.HttpTunnel.Port, false)

	pp.Println("@run", htun.Run())
}

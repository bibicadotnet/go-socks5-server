package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/armon/go-socks5"
	"github.com/caarlos0/env/v6"
)

type params struct {
	Creds           string        `env:"PROXY_CREDENTIALS" envDefault:""`
	User            string        `env:"PROXY_USER" envDefault:""`
	Password        string        `env:"PROXY_PASSWORD" envDefault:""`
	Port            string        `env:"PROXY_PORT" envDefault:"1080"`
	AllowedDestFqdn string        `env:"ALLOWED_DEST_FQDN" envDefault:""`
	AllowedIPs      []string      `env:"ALLOWED_IPS" envSeparator:"," envDefault:""`
	Timeout         time.Duration `env:"DIAL_TIMEOUT" envDefault:"3s"`
}

func main() {
	cfg := params{}
	if err := env.Parse(&cfg); err != nil {
		os.Exit(1)
	}

	socks5conf := &socks5.Config{
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			d := net.Dialer{Timeout: cfg.Timeout}
			return d.DialContext(ctx, network, addr)
		},
	}

	creds, err := getCredentials(cfg)
	if err != nil {
		os.Exit(1)
	}

	if len(creds) > 0 {
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		socks5conf.AuthMethods = []socks5.Authenticator{cator}
	}

	if cfg.AllowedDestFqdn != "" {
		socks5conf.Rules = PermitDestAddrPattern(cfg.AllowedDestFqdn)
	}

	server, err := socks5.New(socks5conf)
	if err != nil {
		os.Exit(1)
	}

	if len(cfg.AllowedIPs) > 0 {
		whitelist := make([]net.IP, len(cfg.AllowedIPs))
		for i, ip := range cfg.AllowedIPs {
			whitelist[i] = net.ParseIP(ip)
		}
		server.SetIPWhitelist(whitelist)
	}

	if err := server.ListenAndServe("tcp", ":"+cfg.Port); err != nil {
		os.Exit(1)
	}
}

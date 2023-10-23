package main

import (
	"flag"
	"github.com/gost1k337/wb_demo/publisher/config"
	"github.com/gost1k337/wb_demo/publisher/service"
	"github.com/gost1k337/wb_demo/publisher/stan"
	"log"
	"net"
)

const configPath = "config/config.yml"

func main() {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	cc, err := stan.NewStanClient(cfg.Nats.ClusterID, cfg.Nats.ClientID, net.JoinHostPort(cfg.Nats.Host, cfg.Nats.Port))
	if err != nil {
		log.Fatalf("stan: %v", err)
	}
	defer cc.Close()

	countPtr := flag.Int("c", 5, "number of orders to pub")

	flag.Parse()

	svc := service.NewPubService(cc)

	err = svc.PubRandomData(cfg.Nats.Channel, *countPtr)
	if err != nil {
		log.Fatalf("pub: %v", err)
	}
}

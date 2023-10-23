package service

import (
	"fmt"
	"github.com/gost1k337/wb_demo/publisher/stan"
	"log"
	"os"
)

type PubService struct {
	cc *stan.ClientConn
}

func NewPubService(cc *stan.ClientConn) *PubService {
	return &PubService{cc: cc}
}

func (ps *PubService) PubRandomData(channel string, count int) error {
	file, err := os.ReadFile("./data/model.json")
	if err != nil {
		return fmt.Errorf("no schema: %w", err)
	}

	for i := 1; i <= count; i++ {
		err := ps.cc.SC.Publish(channel, file)
		fmt.Printf("Publish number %d\n", i)
		if err != nil {
			log.Printf("[pub %d]: %v\n", i, err)
		}
	}

	return nil
}

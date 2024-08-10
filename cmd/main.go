package main

import (
	"encoding/json"
	"fmt"

	"github.com/RuanRosa/quake/domain"
	"github.com/RuanRosa/quake/domain/ports"
	"github.com/RuanRosa/quake/gateways"
)

type Game map[string]Score

type Score struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
	KillsType  map[string]int `json:"deaths_from"`
}

func main() {
	fileSvc := gateways.NewFile()
	quakeRxPatterns := ports.NewQuakeRxPatterns()
	quakeDomain := domain.NewReadQuakeLog(fileSvc, quakeRxPatterns)
	logs, err := quakeDomain.ReadQuakeLog()
	if err != nil {
		panic(fmt.Errorf("read quake log error: %w", err))
	}

	// TODO: This code should be in an api output
	output := Game{}
	for i := 1; i <= len(logs); i++ {
		key := fmt.Sprintf("game_%d", i)
		output[key] = Score{
			TotalKills: logs[i].TotalKills,
			Kills:      logs[i].Kills,
			Players:    logs[i].Players,
			KillsType:  make(map[string]int),
		}

		for _, v := range logs[i].KillsType {
			output[key].KillsType[v.Type] = v.Quantity
		}
	}

	b, err := json.Marshal(output)
	if err != nil {
		panic(fmt.Errorf("json marshal output: %w", err))
	}
	//

	fmt.Println(string(b))
}

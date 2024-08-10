package domain

import (
	"fmt"
	"strings"

	"github.com/RuanRosa/quake/domain/ports"
	"github.com/RuanRosa/quake/utils"
)

type fileSvc interface {
	GetLines() ([]string, error)
}

type readQuakeLog struct {
	fileSvc fileSvc
	rx      ports.QuakeRxPatterns
	games   ports.Game
	score   ports.Score
}

func (r *readQuakeLog) ReadQuakeLog() (ports.ReadQuakeLogOutput, error) {
	gameCount := 0

	lines, err := r.fileSvc.GetLines()
	if err != nil {
		return ports.ReadQuakeLogOutput(r.games), fmt.Errorf("get file: %w", err)
	}

	for _, line := range lines {
		if r.rx.InitGame.MatchString(line) {
			r.score = ports.Score{
				Kills: make(map[string]int),
			}
			gameCount++
		}

		if killsMatch := r.rx.Killed.FindStringSubmatch(line); killsMatch != nil {
			r.Kills(killsMatch)
		}

		if playersMatch := r.rx.ClientUserInfoChanged.FindStringSubmatch(line); playersMatch != nil {
			r.Player(playersMatch)
		}

		if r.rx.ShutdownGame.MatchString(line) {
			r.games[gameCount] = r.score
			r.score = ports.Score{}
		}

	}

	return ports.ReadQuakeLogOutput(r.games), nil
}

func (r *readQuakeLog) Kills(matchs []string) {
	killer := strings.TrimSpace(matchs[1])
	victim := strings.TrimSpace(matchs[2])

	r.KillsType(strings.TrimSpace(matchs[3]))

	if suicide := killer == victim; suicide {
		r.score.TotalKills++
		return
	}

	if killer == "<world>" {
		r.score.Kills[victim] -= 1
	}

	if killer != "<world>" {
		r.score.Kills[killer] += 1
	}

	r.score.TotalKills++
}

func (r *readQuakeLog) KillsType(match string) {
	for k, v := range r.score.KillsType {
		if v.Type == match {
			r.score.KillsType[k].Quantity += 1
			return
		}
	}

	r.score.KillsType = append(r.score.KillsType, ports.KillType{
		Type:     match,
		Quantity: 1,
	})
}

func (r *readQuakeLog) Player(matchs []string) {
	player := strings.TrimSpace(matchs[1])
	if !utils.ExistsStringInArray(player, r.score.Players) {
		r.score.Players = append(r.score.Players, player)
	}
}

func NewReadQuakeLog(
	fileSvc fileSvc,
	rx ports.QuakeRxPatterns,
) readQuakeLog {
	return readQuakeLog{
		fileSvc: fileSvc,
		games:   make(ports.Game),
		rx:      rx,
	}
}

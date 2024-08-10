package ports

import "regexp"

type Game map[int]Score

type Score struct {
	TotalKills int
	Players    []string
	Kills      map[string]int
	KillsType  []KillType
}

type KillType struct {
	Type     string
	Quantity int
}

type ReadQuakeLogOutput Game

type QuakeRxPatterns struct {
	InitGame              *regexp.Regexp
	ShutdownGame          *regexp.Regexp
	Killed                *regexp.Regexp
	ClientUserInfoChanged *regexp.Regexp
}

func NewQuakeRxPatterns() QuakeRxPatterns {
	return QuakeRxPatterns{
		InitGame:              regexp.MustCompile(`InitGame`),
		ShutdownGame:          regexp.MustCompile(`ShutdownGame`),
		Killed:                regexp.MustCompile(`(<world>|[\w\s]+) killed ([\w\s]+) by (\w+)`),
		ClientUserInfoChanged: regexp.MustCompile(`ClientUserinfoChanged: \d+ n\\([^\\]+)`),
	}
}

package main

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	// ErrInvalidLineFormat is returned by parsing functions that cannot parse a given line
	ErrInvalidLineFormat = errors.New("invalid line format")

	// ErrReasonEmpty is returned when the reason is empty
	ErrReasonEmpty = errors.New("the reason is empty")

	// 0: Full
	// 1: ID Voter
	// 2: Name Voter
	// 3: ID Victim
	// 4: Name Victim
	// 5: Reason
	// 6: Command
	// 7: Forced
	startVoteKickRegex = regexp.MustCompile(`'([\d]{1,2}):(.*)' voted kick '([\d]{1,2}):(.*)' reason='(.{1,20})' cmd='(.*)' force=([\d])`)
	// 0: Full
	// 1: ID Voter
	// 2: Name Voter
	// 3: ID Victim
	// 4: Name Victim
	// 5: Reason
	// 6: Command
	// 7: Forced
	startVoteSpecRegex = regexp.MustCompile(`'([\d]{1,2}):(.*)' voted spectate '([\d]{1,2}):(.*)' reason='(.{1,20})' cmd='(.*)' force=([\d])`)
)

var Parsers = []func(string) (string, error){
	StartVoteKick,
	StartVoteSpec,
}

func StartVoteKick(logLine string) (string, error) {
	match := startVoteKickRegex.FindStringSubmatch(logLine)
	if len(match) == 0 {
		return "", fmt.Errorf("%w: StartVoteKick: %s", ErrInvalidLineFormat, logLine)
	}
	reason := match[5]

	if reason == "" {
		return "", ErrReasonEmpty
	}
	return reason, nil
}

func StartVoteSpec(logLine string) (string, error) {
	match := startVoteSpecRegex.FindStringSubmatch(logLine)
	if len(match) == 0 {
		return "", fmt.Errorf("%w: StartVoteSpec: %s", ErrInvalidLineFormat, logLine)
	}
	reason := match[5]
	if reason == "" {
		return "", ErrReasonEmpty
	}
	return reason, nil
}

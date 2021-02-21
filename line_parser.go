package main

func lineParser(line string) (string, error) {
	var err error
	var reason string
	for _, parse := range Parsers {
		reason, err := parse(line)
		if err == nil {
			return reason, nil
		}
	}
	return reason, err
}

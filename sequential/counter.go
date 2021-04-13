package sequential

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

//Parse content of file in sequential mode.
//Read each line separated by the char '.'
//For each line split by space and accumulate
//the number of words in that line.
func Counter(fn string) (int, error) {
	lns, err := splitFile(fn)
	if err != nil {
		log.Print(err.Error())
		return -1, err
	}
	acc := 0
	for _, ln := range lns {
		if ln == "" {
			continue //ignore blank lines
		}
		acc = acc + len(strings.Split(ln, " "))
	}
	return acc, nil
}

func splitFile(fn string) ([]string, error) {
	fl, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, errors.New("Could not open file " + fn)
	}
	return strings.Split(string(fl), "."), nil
}
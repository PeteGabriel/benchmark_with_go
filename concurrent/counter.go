package concurrent

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

//CountWords count amount of words in a given file using concurrency.
//Read content of file.
//Start a goroutine for each statement inside the file.
//Wait for notifications about number of words and accumulate the number.
func CountWords(fn string) (int, error) {
	lns, err := splitFile(fn)
	if err != nil {
		log.Print(err.Error())
		return -1, err
	}


	lnsChl := make(chan int, len(lns))

	var wg sync.WaitGroup

	for _, ln := range lns {
		wg.Add(1)
		log.Print("wait + 1")
		go countWords(&wg, ln, lnsChl)
	}

	wg.Wait()
	close(lnsChl)

	acc := 0
	for c := range lnsChl {
		acc = acc + c
	}

	return acc, nil
}

func countWords(wg *sync.WaitGroup, ln string, lnChl chan int){
	defer wg.Done()
	if ln == "" {
		log.Print("send value")
		lnChl <- 0 //ignore blank lines
	}else {
		log.Print("send value")
		lnChl <- len(strings.Split(ln, " "))
	}
	log.Print("wait - 1")
}


func splitFile(fn string) ([]string, error) {
	fl, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, errors.New("Could not open file " + fn)
	}
	return strings.Split(string(fl), "."), nil
}
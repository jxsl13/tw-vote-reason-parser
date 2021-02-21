package main

import (
	"bufio"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"os"

	"github.com/jszwec/csvutil"
)

var reasonSet CSet = NewCSet()

func main() {
	wg := &sync.WaitGroup{}
	err := filepath.Walk("./logs", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".txt") {
			return nil
		}
		fmt.Printf("started processing of file: %q\n", path)
		wg.Add(1)
		go processLogFile(wg, path)
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Waiting for processing to finish")
	wg.Wait()
	fmt.Println("Parsing finished")

	fh, err := os.OpenFile("reasons.csv", os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	reasons := reasonSet.SortedList()
	b, err := csvutil.Marshal(reasons)
	if err != nil {
		fmt.Println("error:", err)
	}

	_, err = fh.Write(b)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Wrote reasons to reasons.csv.... done!")
}

func processLogFile(wg *sync.WaitGroup, path string) error {
	defer wg.Done()
	fh, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		reason, err := lineParser(line)
		if err != nil {
			continue
		}
		if reason != "" {
			reasonSet.Add(reason)
		}
	}
	return nil
}

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/akanshat/phaidra/pkg/datastore"
)

type query struct {
	name, qmin, qmax string
}

func runService(input io.ReadCloser, q chan query, r chan []datastore.SensorType) error {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}
	input.Close()

	ds, err := datastore.NewDataStore(data)
	if err != nil {
		return err
	}

	for v := range q {
		r <- ds.Query(v.name, v.qmin, v.qmax)
	}
	return nil
}

func stringToQuery(str string) (query, error) {
	s := strings.Split(str, " ")
	if len(s) != 3 {
		return query{}, errors.New("invalid query")
	}

	return query{
		name: s[0],
		qmin: s[1],
		qmax: s[2],
	}, nil
}

func main() {
	var path string
	flag.StringVar(&path, "filepath", "./data.json", "File containing sensor and equipment data")
	flag.Parse()

	file, err := os.Open(path)
	if err != nil {
		return
	}

	qch := make(chan query)
	resch := make(chan []datastore.SensorType)

	go (func() {
		for v := range resch {
			fmt.Printf("%+v\n", v)
		}
	})()

	go runService(file, qch, resch)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		q, err := stringToQuery(scanner.Text())
		if err != nil {
			continue
		}
		qch <- q
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}

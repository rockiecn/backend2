package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/rockiecn/backend2/filter"
	"github.com/rockiecn/backend2/parser"
	"github.com/rockiecn/backend2/recorder"
)

func main() {
	// new Filter
	a, err := filter.New()
	if err != nil {
		log.Fatal("new Filter failed")
	}

	// initialze
	a.Initialize()

	// get all logs of a block
	logs, err := a.Logs(big.NewInt(3104475))
	if err != nil {
		log.Fatal(err)
	}

	// parser
	p := parser.New()
	p.Initialize()

	// show logs
	fmt.Printf("==== all logs, %d\n", len(logs))
	for i, v := range logs {
		fmt.Printf("--> log %d:\n", i)

		// get parser func with topics[0]
		parseFunc := p.ParseMethods[v.Topics[0]]
		if parseFunc != nil {
			// call a parser
			fmt.Println("parsing...")
			e, err := parseFunc(logs[5])
			if err != nil {
				log.Fatal(err)
			}

			// call recorder
			fmt.Println("recording...")
			err = recorder.ReAccRecorder(e)
			if err != nil {
				log.Fatal("record ReAcc failed")
			}
		} else {
			continue
		}

	}

}

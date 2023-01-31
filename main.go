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
	f, err := filter.New()
	if err != nil {
		log.Fatal("new Filter failed")
	}
	// initialize filter
	f.Initialize()

	// get logs of a specified block
	logs, err := f.Logs(big.NewInt(3104475))
	if err != nil {
		log.Fatal(err)
	}

	// new parser
	p := parser.New()
	// initialize parser
	p.Initialize()

	// parse all logs into event struct and record
	fmt.Printf("==== all logs, %d\n", len(logs))
	for i, v := range logs {
		fmt.Printf("\n--> log %d:\n", i)
		fmt.Printf("block: %d\n", v.BlockNumber)

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
			fmt.Println("no parser func for this log")
			continue
		}

	}

}

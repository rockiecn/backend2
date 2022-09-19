package filter

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Filter struct {
	// endpoint of chain
	Endpoint string
}

// create an empty Filter, initialization needed
func New() (*Filter, error) {
	ep := ""

	a := &Filter{
		Endpoint: ep,
	}

	return a, nil
}

// initialize an Filter with endpoint, fill all hashes of event
func (a *Filter) Initialize() {
	a.Endpoint = "https://testchain.metamemo.one:24180"
}

// get all logs of a specified lock, with filter table
func (a *Filter) Logs(block *big.Int) ([]types.Log, error) {
	if a.Endpoint == "" {
		return nil, errors.New("need initialize")
	}

	// dial block chain
	client, err := ethclient.Dial(a.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fb := new(big.Int).Sub(block, big.NewInt(1))
	tb := block
	fmt.Printf("from: %s, to: %s\n", fb.String(), tb.String())

	// query a block for all logs, need a flter table
	query := ethereum.FilterQuery{
		FromBlock: fb,
		ToBlock:   tb,

		// filter table
		// Addresses: []common.Address{
		// 	common.HexToAddress(contractAddr),
		// },
		Addresses: []common.Address{},
	}

	// filter logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	if len(logs) == 0 {
		return nil, nil
	}

	return logs, nil
}

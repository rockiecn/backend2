package filter

import (
	"context"
	"errors"
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

	f := &Filter{
		Endpoint: ep,
	}

	return f, nil
}

// initialize an Filter with endpoint, fill all hashes of event
func (f *Filter) Initialize() {
	f.Endpoint = "https://testchain.metamemo.one:24180"
}

// get all logs of a specified lock, with filter table
func (f *Filter) Logs(block *big.Int) ([]types.Log, error) {
	if f.Endpoint == "" {
		return nil, errors.New("need initialize")
	}

	// dial block chain
	client, err := ethclient.Dial(f.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// query a block for all logs, need a flter table
	query := ethereum.FilterQuery{
		FromBlock: block,
		ToBlock:   block,

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

package parser

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/backend2/events"
)

// hash table
var (
	ReAccHash = common.HexToHash("e22ccfc11bbbb1d3350012248b26ad616e32a8a13eece408dd14f1953fe24752")
)

// log parser func type
type ParserFunc func(types.Log) (interface{}, error)

type Parser struct {
	// all parsers of each event
	ParseMethods map[common.Hash]ParserFunc
}

// new parser
func New() *Parser {
	pm := make(map[common.Hash]ParserFunc)
	p := &Parser{ParseMethods: pm}
	return p
}

// initialization
func (p *Parser) Initialize() error {
	p.ParseMethods[ReAccHash] = ReAccParser
	return nil
}

// parser for register account
func ReAccParser(log types.Log) (interface{}, error) {
	ev := &events.ReAcc{}
	// get topics
	ev.Topics[0] = log.Topics[0]

	// get data
	ev.Data.Address.SetBytes(log.Data[0:32])
	ev.Data.Index = new(big.Int).SetBytes(log.Data[32:64])

	return ev, nil
}

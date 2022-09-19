package events

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// register account event
type ReAcc struct {
	Topics [1]common.Hash
	Data   struct {
		Address common.Address
		Index   *big.Int
	}
}

package contract

import (
	"gin-api/contract/hello"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Hello struct {
}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) DeployContract(auth *bind.TransactOpts, client *ethclient.Client ) (common.Address, *types.Transaction, *hello.Hello, error) {
	return hello.DeployHello(auth, client)
}

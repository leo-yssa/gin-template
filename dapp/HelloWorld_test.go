package dapp_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/suite"
)

type HelloworldTestSuite struct {
	suite.Suite
	ctx context.Context
	auth       *bind.TransactOpts
	address    common.Address
	gAlloc     types.GenesisAlloc
	sim        *simulated.Backend
	helloworld *Dapp
}

func TestRunHelloworldSuite(t *testing.T) {
	suite.Run(t, new(HelloworldTestSuite))
}

func (s *HelloworldTestSuite) SetupTest() {
	s.ctx = context.Background()
	key, _ := crypto.GenerateKey()
	s.auth, _ = bind.NewKeyedTransactorWithChainID(key, big.NewInt(int64(1337)))
	s.address = s.auth.From
	s.gAlloc = map[common.Address]types.Account{
		s.address: {Balance: big.NewInt(350000000000000)},
	}
	s.sim = simulated.NewBackend(s.gAlloc)
	c := s.sim.Client()
	c.SendTransaction(s.ctx, )
	_, _, hw, e := DeployDapp(s.auth, s.sim)
	s.helloworld = hw
	s.Nil(e)
	s.sim.Commit()
	// s.sim.SendTransaction()
}

func (s *HelloworldTestSuite) TestSay() {
	str, err := s.helloworld.Say(
		&bind.CallOpts{
			From: s.auth.From,
			Context: s.ctx,
		},
	)
	s.Equal("hello etherworld", str)
	s.Nil(err)
}
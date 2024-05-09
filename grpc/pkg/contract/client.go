package contract

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Network int

const (
	ETHEREUM_TESTNET = iota
	ETHEREUM_MAINNET
	HEDERA_TESTNET
	HEDERA_MAINNET
)

type Client struct {
	Ethclient *ethclient.Client
	chainId *big.Int
	Hello *Hello
}

func NewClient(network Network) (*Client, error) {
	var address string
	var chainId *big.Int
	switch network {
	case ETHEREUM_TESTNET:
		address = "https://rpc2.sepolia.org"
		chainId = big.NewInt(11155111)
	case HEDERA_TESTNET:
		address = "https://testnet.hashio.io/api"
		chainId = big.NewInt(296)
	}
	ethclient, err := ethclient.Dial(address)
	if err != nil {
		return nil, err
	}
	return &Client{
		Ethclient: ethclient,
		chainId: chainId,
		Hello: NewHello(),
	}, nil
}

func HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	if strings.HasPrefix(hexkey, "0x") {
		hexkey = hexkey[2:]
	}
	return crypto.HexToECDSA(hexkey)
}

func ToAddress(key *ecdsa.PrivateKey) (common.Address, error) {
	return crypto.PubkeyToAddress(key.PublicKey), nil
}

func (c *Client) pendingNonceAt(from common.Address) (*big.Int, error) {
	nonce, err := c.Ethclient.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(nonce)), nil
}

func (c *Client) suggestGasPrice() (*big.Int, error) {
	return c.Ethclient.SuggestGasPrice(context.Background())
}

func (c *Client) NewTransactor(from common.Address, value *big.Int, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainId)
	if err != nil {
		return nil, err
	}
	nonce, err := c.pendingNonceAt(from)
	if err != nil {
		return nil, err
	}
	gasPrice, err := c.suggestGasPrice()
	if err != nil {
		return nil, err
	}
	auth.Nonce = nonce
	auth.Value = value
	auth.GasPrice = gasPrice
	return auth, nil
}



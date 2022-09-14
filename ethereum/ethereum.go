package ethereum

import (
	"crypto/ecdsa"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InitEthereum() *ethclient.Client {
	client, err := ethclient.Dial(os.Getenv("CHAIN_LINK"))
	if err != nil {
		panic(err)
	}
	return client
}

func CreateWallet() *ecdsa.PrivateKey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	return privateKey
}

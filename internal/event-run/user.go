package eventrun

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"sync"

	"github.com/qwertyqq2/event-listener/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Params struct {
		PrivateKey string `yaml:"private_key"`
	} `yaml:"params_user"`
}

type User struct {
	conf   *Config
	client *ethclient.Client
}

var (
	U            *User
	AddrContract common.Address
	once         sync.Once
)

func GetUserConfig() *Config {
	conf := &Config{}
	once.Do(func() {
		if err := cleanenv.ReadConfig("config.yaml", conf); err != nil {
			log.Fatal(err)
		}
	})
	log.Println(conf.Params)
	return conf
}

func init() {
	U = &User{
		conf: GetUserConfig(),
	}
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	U.client = client
	auth, err := U.CreateAuth(0)
	if err != nil {
		log.Fatal(err)
	}

	addr, tx, instance, err := api.DeployApi(auth, client)
	log.Println("Contact created: ", addr)
	AddrContract = addr
	if err != nil {
		panic(err)
	}

	U.client = client
	_, _ = instance, tx
}

func (u *User) CreateAuth(val int64) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(u.conf.Params.PrivateKey)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := u.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	chainID, err := u.client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(val)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(1000000)
	return auth, nil
}

func (u *User) Deposit(val int64) {
	auth, err := u.CreateAuth(val)
	if err != nil {
		log.Fatal(err)
	}
	instance, err := api.NewApi(AddrContract, u.client)
	if err != nil {
		log.Fatal(err)
	}

	_, err = instance.Deposit(auth)

	if err != nil {
		log.Fatal(err)
	}

}

func (u *User) Close() {
	auth, err := u.CreateAuth(0)
	if err != nil {
		log.Fatal(err)
	}
	instance, err := api.NewApi(AddrContract, u.client)
	if err != nil {
		log.Fatal(err)
	}

	_, err = instance.Close(auth)

	if err != nil {
		log.Fatal(err)
	}
}

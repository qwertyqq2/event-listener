package provider

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/qwertyqq2/event-listener/internal/event-listener/provider/model"
	"github.com/qwertyqq2/event-listener/pkg/mongodb"
	"github.com/qwertyqq2/event-listener/pkg/mongodb/event/events"

	"github.com/qwertyqq2/event-listener/api"
	eventrun "github.com/qwertyqq2/event-listener/internal/event-run"
)

type Provider struct {
	client          *ethclient.Client
	contractAddress common.Address
	privateKey      *ecdsa.PrivateKey
	api             *api.Api
	auth            *bind.TransactOpts
}

func (p *Provider) SetUp() {
	conf := GetConfig()

	pClient, err := ethclient.Dial(conf.ParamsProvider.ProviderURL)
	if err != nil {
		log.Fatal(err)
	}
	p.client = pClient
	pPrivateKey, err := crypto.HexToECDSA(conf.ParamsProvider.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	p.privateKey = pPrivateKey
	p.contractAddress = eventrun.AddrContract
	pContractClient, err := api.NewApi(p.contractAddress, p.client)
	if err != nil {
		log.Fatal(err)
	}
	p.api = pContractClient
	p.auth = bind.NewKeyedTransactor(p.privateKey)
	log.Println("Provider created: ", p.contractAddress.Hex())
}

func (p *Provider) ListenToEvent() {
	depositAbi, err := abi.JSON(strings.NewReader(string(api.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{p.contractAddress},
	}
	ctx := context.Background()

	var eventCh = make(chan types.Log)

	sub, err := p.client.SubscribeFilterLogs(ctx, query, eventCh)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case err := <-sub.Err():
			log.Fatal(err)
		case logEvent := <-eventCh:
			switch logEvent.Topics[0].Hex() {
			case model.LogCreateSigHash:
				depositEvent := &model.LogDeposit{}
				err := depositAbi.UnpackIntoInterface(depositEvent, "CreateDeposit", logEvent.Data)
				if err != nil {
					log.Fatal(err)
				}
				dep := events.NewDeposit("CreateDeposit", depositEvent.From.String(), int(depositEvent.Value.Int64()))
				if err := mongodb.Insert(ctx, dep); err != nil {
					log.Fatal(err)
				}

				fmt.Printf("owner: %s, value: %d \n", depositEvent.From, depositEvent.Value)

			case model.LogCloseSigHash:
				closeEvent := &model.LogClose{}
				err := depositAbi.UnpackIntoInterface(closeEvent, "CloseDeposit", logEvent.Data)
				if err != nil {
					log.Fatal(err)
				}
				close := events.NewClose("CloseDeposit", closeEvent.From.String())
				if err := mongodb.Insert(ctx, close); err != nil {
					log.Fatal(err)
				}

				fmt.Printf("owner: %s \n", closeEvent.From)
			}

		}
	}
}

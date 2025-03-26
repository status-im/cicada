package feeds

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumEventFeed struct {
	rpcURL     string
	contract   common.Address
	eventSig   common.Hash
	startBlock uint64
	client     *ethclient.Client
}

func NewEthereumEventFeed(rpcURL, contractAddr, eventSig string, startBlock uint64) (*EthereumEventFeed, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &EthereumEventFeed{
		rpcURL:     rpcURL,
		contract:   common.HexToAddress(contractAddr),
		eventSig:   crypto.Keccak256Hash([]byte(eventSig)), // e.g. Transfer(address,address,uint256)
		startBlock: startBlock,
		client:     client,
	}, nil
}

func (e *EthereumEventFeed) Name() string {
	return fmt.Sprintf("eth:%s:%s", e.contract.Hex(), e.eventSig.Hex())
}

func (e *EthereumEventFeed) FetchItems() ([]FeedItem, error) {
	ctx := context.Background()

	// Get latest block
	latest, err := e.client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	// Define query
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(e.startBlock)),
		ToBlock:   big.NewInt(int64(latest)),
		Addresses: []common.Address{e.contract},
		Topics:    [][]common.Hash{{e.eventSig}},
	}

	logs, err := e.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	// Advance block for next poll
	e.startBlock = latest + 1

	var items []FeedItem
	for _, logEntry := range logs {
		// TODO: careful of too many logs in one poll potential
		//  issue, try caching block timestamps by number to
		//  avoid fetching the same block multiple times.

		var timestamp time.Time
		blockNumber := big.NewInt(int64(logEntry.BlockNumber))
		block, err := e.client.BlockByNumber(ctx, blockNumber)
		if err != nil {
			log.Printf("Error getting block %d: %v", logEntry.BlockNumber, err)
			timestamp = time.Now() // fallback
		} else {
			timestamp = time.Unix(int64(block.Time()), 0)
		}

		items = append(items, FeedItem{
			ID:        logEntry.TxHash.Hex() + fmt.Sprintf(":%d", logEntry.Index),
			Title:     fmt.Sprintf("Event at block %d", logEntry.BlockNumber),
			Link:      fmt.Sprintf("https://etherscan.io/tx/%s", logEntry.TxHash.Hex()),
			Timestamp: timestamp,
		})
	}

	return items, nil
}

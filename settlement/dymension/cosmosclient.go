package dymension

import (
	"context"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/cosmosclient/cosmosclient"
	rollapptypes "github.com/dymensionxyz/dymension/x/rollapp/types"
	sequencertypes "github.com/dymensionxyz/dymension/x/sequencer/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// CosmosClient is an interface for interacting with cosmos client chains.
// It is a wrapper around the cosmos client in order to provide with an interface which can be implemented by
// other clients and can easily be mocked for testing purposes.
// Currently it contains only the methods that are used by the dymension hub client.
type CosmosClient interface {
	Context() sdkclient.Context
	StartEventListener() error
	StopEventListener() error
	EventListenerQuit() <-chan struct{}
	SubscribeToEvents(ctx context.Context, subscriber string, query string, outCapacity ...int) (out <-chan ctypes.ResultEvent, err error)
	BroadcastTx(accountName string, msgs ...sdktypes.Msg) (cosmosclient.Response, error)
	GetRollappClient() rollapptypes.QueryClient
	GetSequencerClient() sequencertypes.QueryClient
}

type cosmosClient struct {
	cosmosclient.Client
}

var _ CosmosClient = &cosmosClient{}

// NewCosmosClient creates a new cosmos client
func NewCosmosClient(client cosmosclient.Client) CosmosClient {
	return &cosmosClient{client}
}

func (c *cosmosClient) StartEventListener() error {
	return c.Client.RPC.WSEvents.Start()
}

func (c *cosmosClient) StopEventListener() error {
	return c.Client.RPC.WSEvents.Stop()
}

func (c *cosmosClient) EventListenerQuit() <-chan struct{} {
	return c.Client.RPC.GetWSClient().Quit()
}

func (c *cosmosClient) SubscribeToEvents(ctx context.Context, subscriber string, query string, outCapacity ...int) (out <-chan ctypes.ResultEvent, err error) {
	return c.Client.RPC.WSEvents.Subscribe(ctx, subscriber, query, outCapacity...)
}

func (c *cosmosClient) GetRollappClient() rollapptypes.QueryClient {
	return rollapptypes.NewQueryClient(c.Context())
}

func (c *cosmosClient) GetSequencerClient() sequencertypes.QueryClient {
	return sequencertypes.NewQueryClient(c.Context())
}

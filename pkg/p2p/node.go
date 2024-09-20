package p2p

import (
	"context"
	"math/rand"
	"time"

	"github.com/JITLiq/node/pkg/logger"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"go.uber.org/zap"
)

const (
	serviceNode         = "node"
	DiscoveryServiceTag = "jit-liq-node"
)

type NodeConfig struct {
	Addr string `json:"addr"`
}

type Node struct {
	host   host.Host
	logger *zap.Logger
	topic  *pubsub.Topic
}

func New(cfg *NodeConfig) (*Node, error) {
	var err error
	node := &Node{
		logger: logger.NewLogger(serviceNode),
	}
	node.host, err = libp2p.New(libp2p.ListenAddrStrings(cfg.Addr))
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (n *Node) HandlePeerFound(pi peer.AddrInfo) {
	n.logger.Info("discovered new peer", zap.String("id", pi.ID.String()))
	rand.Seed(time.Now().UnixNano())

	// generate a random integer between 0 and 9
	randomInt := rand.Intn(5)
	<-time.After(time.Duration(randomInt) * time.Second)
	err := n.host.Connect(context.Background(), pi)
	if err != nil {
		n.logger.Error("error connecting to peer", zap.String("id", pi.ID.String()), zap.Error(err))
	}
}

func (n *Node) SetupDiscovery() error {
	s := mdns.NewMdnsService(n.host, DiscoveryServiceTag, n)
	return s.Start()
}

func (n *Node) SetupPubSub(ctx context.Context, topic string) error {
	queue, err := pubsub.NewGossipSub(ctx, n.host)
	if err != nil {
		return err
	}

	n.topic, err = queue.Join(topic)
	return err
}

func (n *Node) Subscriber() (*pubsub.Subscription, error) {
	return n.topic.Subscribe()
}

func (n *Node) Publish(ctx context.Context, data []byte) error {
	return n.topic.Publish(ctx, data)
}

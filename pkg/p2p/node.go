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
	DiscoveryServiceTag = "jit-liq-node-discovery"
	MaxConnectDelay     = 5
	maxRetries          = 5
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
	n.connectWithRetry(context.Background(), pi)
}

func (n *Node) connectWithRetry(ctx context.Context, pi peer.AddrInfo) {
	for i := 0; i < maxRetries; i++ {
		n.logger.Info("trying to connect to peer", zap.String("id", pi.ID.String()))

		<-time.After(time.Second * time.Duration(rand.Intn(MaxConnectDelay)))
		err := n.host.Connect(ctx, pi)
		if err == nil {
			n.logger.Info("successfully connected to peer", zap.String("id", pi.ID.String()))
			return
		}
		n.logger.Warn("failed to connect to peer, retrying", zap.String("id", pi.ID.String()), zap.Error(err), zap.Int("attempt", i+1))
	}
	n.logger.Error("failed to connect to peer after multiple attempts", zap.String("id", pi.ID.String()))
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

func (n *Node) Close() error {
	if n.topic != nil {
		n.topic.Close()
	}
	return n.host.Close()
}

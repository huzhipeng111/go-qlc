package p2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-discovery"
	"github.com/libp2p/go-libp2p-host"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	localdiscovery "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/qlcchain/go-qlc/config"
)

const (
	QlcProtocolID           = "qlc/1.0.0"
	QlcProtocolFOUND        = "/qlc/discovery/1.0.0"
	discoveryConnTimeout    = time.Second * 30
	EachTimeFoundPeersLimit = 20
)

func (node *QlcNode) dhtFoundPeers() ([]pstore.PeerInfo, error) {
	//discovery peers
	peers, err := discovery.FindPeers(node.ctx, node.dis, QlcProtocolFOUND, EachTimeFoundPeersLimit)
	if err != nil {
		return nil, err
	}
	logger.Infof("Found %d peers!\n", len(peers))
	for _, p := range peers {
		logger.Info("Peer: ", p)
	}
	return peers, nil
}

// HandlePeerFound attempts to connect to peer from `PeerInfo`.
func (node *QlcNode) HandlePeerFound(p pstore.PeerInfo) {
	ctx, cancel := context.WithTimeout(node.ctx, discoveryConnTimeout)
	defer cancel()
	if err := node.host.Connect(ctx, p); err != nil {
		logger.Error("Failed to connect to peer found by discovery: ", err)
	}
	logger.Info("find a local peer , ID:", p.ID.Pretty())
}
func setupDiscoveryOption(cfg *config.Config) DiscoveryOption {
	if cfg.MdnsEnabled {
		return func(ctx context.Context, h host.Host) (localdiscovery.Service, error) {
			if cfg.MdnsInterval == 0 {
				cfg.MdnsInterval = 5
			}
			return localdiscovery.NewMdnsService(ctx, h, time.Duration(cfg.MdnsInterval)*time.Second, QlcProtocolID)
		}
	}
	return nil
}

type DiscoveryOption func(context.Context, host.Host) (localdiscovery.Service, error)

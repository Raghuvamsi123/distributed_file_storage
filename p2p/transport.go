package p2p

// peer is an interface that represents the remote node

type Peer interface {
	Close() error
}

// transport is anything that handles the communication between the nodes in the network.

// Can be of the form tcp/udp/ websockets.

type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}

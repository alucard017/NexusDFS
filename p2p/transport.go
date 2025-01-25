package p2p

//Peer is an interface that represents a remote node in the network
type Peer interface{}

//Transport is an interface that handles communication between the nodes in the network.
//This can be in any of the form i.e either TCP, UDP, WebSockets... etc
type Transport interface {
	ListenAndAccept() error
}

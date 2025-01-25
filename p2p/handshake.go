package p2p

//It is returned if the handshake between the local and remote node couldn't be established
// var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }

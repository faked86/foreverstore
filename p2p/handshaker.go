package p2p

import "errors"

// ErrInvalidHandshake is returned if the handshake between
// the Local and remote node could not be established.
var ErrInvalidHandshake = errors.New("invalid handshake")

// HandshakeFunk... ?
type HandshakeFunc func(Peer) error

func NOPHandshkeFunk(Peer) error { return nil }

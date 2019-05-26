package rpc

import (
	"bytes"
	"crypto/x509"
	"time"

	clientpb "github.com/bishopfox/sliver/protobuf/client"
	"github.com/bishopfox/sliver/server/certs"
	"github.com/bishopfox/sliver/server/core"

	"github.com/golang/protobuf/proto"
)

func rpcPlayers(_ []byte, timeout time.Duration, resp RPCResponse) {

	clientCerts := certs.OperatorClientListCertificates()

	players := &clientpb.Players{Players: []*clientpb.Player{}}
	for _, cert := range clientCerts {
		players.Players = append(players.Players, &clientpb.Player{
			Client: &clientpb.Client{
				Operator: cert.Subject.CommonName,
			},
			Online: isPlayerOnline(cert),
		})
	}

	data, err := proto.Marshal(players)
	if err != nil {
		rpcLog.Errorf("Error encoding rpc response %v", err)
	}
	resp(data, err)
}

// isPlayerOnline - Is a player connected using a given certificate
func isPlayerOnline(cert *x509.Certificate) bool {
	for _, client := range *core.Clients.Connections {
		if client.Certificate == nil {
			continue // Server certificate is nil
		}
		if bytes.Equal(cert.Raw, client.Certificate.Raw) {
			return true
		}
	}
	return false
}

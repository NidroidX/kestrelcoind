package client

import (
	"context"
	"github.com/NidroidX/kestrelcoind/cmd/kestrelcoinwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/NidroidX/kestrelcoind/cmd/kestrelcoinwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the kestrelcoinwalletd server, and returns the client instance
func Connect(address string) (pb.kestrelcoinwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("kestrelcoinwallet daemon is not running, start it with `kestrelcoinwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewkestrelcoinwalletdClient(conn), func() {
		conn.Close()
	}, nil
}

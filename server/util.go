package server

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

func GetIDTokenFromMetadata(ctx context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get grpc metadata from context")
	}

	xheader, ok := md["X-Firebase-ID-Token"]
	if !ok {
		return nil, fmt.Errorf("missing X-Firebase-ID-Token header")
	}

	idtoken := strings.Join(xheader, ",")
	return &idtoken, nil
}

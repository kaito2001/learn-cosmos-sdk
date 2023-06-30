package keeper

import (
	"context"

	"vbi-cosmos-basic/x/vbicosmosbasic/types"

	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SayHello(goCtx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QuerySayHelloResponse{Name: fmt.Sprintf("Hello, %s handsome!", req.Name)}, nil
}

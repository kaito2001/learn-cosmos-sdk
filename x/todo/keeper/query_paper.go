package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"vbi-cosmos-basic/x/todo/types"
)

func (k Keeper) PaperAll(goCtx context.Context, req *types.QueryAllPaperRequest) (*types.QueryAllPaperResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var papers []types.Paper
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	paperStore := prefix.NewStore(store, types.KeyPrefix(types.PaperKey))

	pageRes, err := query.Paginate(paperStore, req.Pagination, func(key []byte, value []byte) error {
		var paper types.Paper
		if err := k.cdc.Unmarshal(value, &paper); err != nil {
			return err
		}

		papers = append(papers, paper)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPaperResponse{Paper: papers, Pagination: pageRes}, nil
}

func (k Keeper) Paper(goCtx context.Context, req *types.QueryGetPaperRequest) (*types.QueryGetPaperResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	paper, found := k.GetPaper(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPaperResponse{Paper: paper}, nil
}

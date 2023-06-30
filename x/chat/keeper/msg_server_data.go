package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"vbi-cosmos-basic/x/chat/types"
)

func (k msgServer) CreateData(goCtx context.Context, msg *types.MsgCreateData) (*types.MsgCreateDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var data = types.Data{
		Creator: msg.Creator,
		Content: msg.Content,
	}

	id := k.AppendData(
		ctx,
		data,
	)

	return &types.MsgCreateDataResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateData(goCtx context.Context, msg *types.MsgUpdateData) (*types.MsgUpdateDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var data = types.Data{
		Creator: msg.Creator,
		Id:      msg.Id,
		Content: msg.Content,
	}

	// Checks that the element exists
	val, found := k.GetData(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetData(ctx, data)

	return &types.MsgUpdateDataResponse{}, nil
}

func (k msgServer) DeleteData(goCtx context.Context, msg *types.MsgDeleteData) (*types.MsgDeleteDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetData(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveData(ctx, msg.Id)

	return &types.MsgDeleteDataResponse{}, nil
}

package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"vbi-cosmos-basic/x/todo/types"
)

func (k msgServer) CreatePaper(goCtx context.Context, msg *types.MsgCreatePaper) (*types.MsgCreatePaperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var paper = types.Paper{
		Creator: msg.Creator,
		Body:    msg.Body,
		Content: msg.Content,
	}

	id := k.AppendPaper(
		ctx,
		paper,
	)

	return &types.MsgCreatePaperResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePaper(goCtx context.Context, msg *types.MsgUpdatePaper) (*types.MsgUpdatePaperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var paper = types.Paper{
		Creator: msg.Creator,
		Id:      msg.Id,
		Body:    msg.Body,
		Content: msg.Content,
	}

	// Checks that the element exists
	val, found := k.GetPaper(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPaper(ctx, paper)

	return &types.MsgUpdatePaperResponse{}, nil
}

func (k msgServer) DeletePaper(goCtx context.Context, msg *types.MsgDeletePaper) (*types.MsgDeletePaperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPaper(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePaper(ctx, msg.Id)

	return &types.MsgDeletePaperResponse{}, nil
}

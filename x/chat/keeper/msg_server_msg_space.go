package keeper

import (
	"context"

	"vbi-cosmos-basic/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
)

func (k msgServer) SendMsgSpace(goCtx context.Context, msg *types.MsgSendMsgSpace) (*types.MsgSendMsgSpaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	k.AppendData(ctx, types.Data{
		Creator: msg.User,
		Content: msg.Content,
	})

	// Construct the packet
	var packet types.MsgSpacePacketData

	packet.User = msg.User
	packet.Content = msg.Content

	// Transmit the packet
	_, err := k.TransmitMsgSpacePacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendMsgSpaceResponse{}, nil
}

package keeper

import (
	"errors"

	"vbi-cosmos-basic/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// TransmitMsgSpacePacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitMsgSpacePacket(
	ctx sdk.Context,
	packetData types.MsgSpacePacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvMsgSpacePacket processes packet reception
func (k Keeper) OnRecvMsgSpacePacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgSpacePacketData) (packetAck types.MsgSpacePacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	k.AppendData(ctx, types.Data{
		Creator: data.User,
		Content: data.Content,
	})

	return packetAck, nil
}

// OnAcknowledgementMsgSpacePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementMsgSpacePacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgSpacePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.MsgSpacePacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutMsgSpacePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutMsgSpacePacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgSpacePacketData) error {

	// TODO: packet timeout logic

	return nil
}

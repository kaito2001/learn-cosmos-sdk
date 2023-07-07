package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePaper = "create_paper"
	TypeMsgUpdatePaper = "update_paper"
	TypeMsgDeletePaper = "delete_paper"
)

var _ sdk.Msg = &MsgCreatePaper{}

func NewMsgCreatePaper(creator string, body string, content string) *MsgCreatePaper {
	return &MsgCreatePaper{
		Creator: creator,
		Body:    body,
		Content: content,
	}
}

func (msg *MsgCreatePaper) Route() string {
	return RouterKey
}

func (msg *MsgCreatePaper) Type() string {
	return TypeMsgCreatePaper
}

func (msg *MsgCreatePaper) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePaper) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePaper) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePaper{}

func NewMsgUpdatePaper(creator string, id uint64, body string, content string) *MsgUpdatePaper {
	return &MsgUpdatePaper{
		Id:      id,
		Creator: creator,
		Body:    body,
		Content: content,
	}
}

func (msg *MsgUpdatePaper) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePaper) Type() string {
	return TypeMsgUpdatePaper
}

func (msg *MsgUpdatePaper) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePaper) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePaper) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePaper{}

func NewMsgDeletePaper(creator string, id uint64) *MsgDeletePaper {
	return &MsgDeletePaper{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePaper) Route() string {
	return RouterKey
}

func (msg *MsgDeletePaper) Type() string {
	return TypeMsgDeletePaper
}

func (msg *MsgDeletePaper) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePaper) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePaper) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

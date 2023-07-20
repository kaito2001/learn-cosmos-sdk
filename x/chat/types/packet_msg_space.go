package types

// ValidateBasic is used for validating the packet
func (p MsgSpacePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p MsgSpacePacketData) GetBytes() ([]byte, error) {
	var modulePacket ChatPacketData

	modulePacket.Packet = &ChatPacketData_MsgSpacePacket{&p}

	return modulePacket.Marshal()
}

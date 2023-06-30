package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:   PortID,
		DataList: []Data{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in data
	dataIdMap := make(map[uint64]bool)
	dataCount := gs.GetDataCount()
	for _, elem := range gs.DataList {
		if _, ok := dataIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for data")
		}
		if elem.Id >= dataCount {
			return fmt.Errorf("data id should be lower or equal than the last id")
		}
		dataIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

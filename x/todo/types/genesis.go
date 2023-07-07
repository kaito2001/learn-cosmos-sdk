package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PaperList: []Paper{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in paper
	paperIdMap := make(map[uint64]bool)
	paperCount := gs.GetPaperCount()
	for _, elem := range gs.PaperList {
		if _, ok := paperIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for paper")
		}
		if elem.Id >= paperCount {
			return fmt.Errorf("paper id should be lower or equal than the last id")
		}
		paperIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

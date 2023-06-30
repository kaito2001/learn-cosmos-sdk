package types

const (
	// ModuleName defines the module name
	ModuleName = "chat"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_chat"

	// Version defines the current version the IBC module supports
	Version = "chat-1"

	// PortID is the default port id that module binds to
	PortID = "chat"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("chat-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DataKey      = "Data/value/"
	DataCountKey = "Data/count/"
)

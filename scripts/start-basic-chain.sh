MONIKER='vbi'
KEY='test1'
KEYRING_BACKEND=file
CHAIN_ID='test1'
TOKEN_BALANCES=100000000stake
TOKEN_STAKE=70000000stake

### Run

vbi-cosmos-basicd init $MONIKER --chain-id $CHAIN_ID

vbi-cosmos-basicd keys add $KEY --keyring-backend $KEYRING_BACKEND

vbi-cosmos-basicd add-genesis-account $KEY $TOKEN_BALANCES

vbi-cosmos-basicd gentx $KEY $TOKEN_STAKE --chain-id $CHAIN_ID --keyring-backend $KEYRING_BACKEND

vbi-cosmos-basicd collect-gentxs

vbi-cosmos-basicd start
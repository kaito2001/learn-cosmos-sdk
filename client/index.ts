import { SigningStargateClient } from '@cosmjs/stargate';
import { DirectSecp256k1HdWallet, OfflineSigner } from '@cosmjs/proto-signing';
import { StdFee } from '@cosmjs/amino';
import { Coin } from 'cosmjs-types/cosmos/base/v1beta1/coin';
import { toBase64 } from '@cosmjs/encoding';
import { TxRaw } from 'cosmjs-types/cosmos/tx/v1beta1/tx'
import axios from 'axios';
import * as dotenv from 'dotenv'
dotenv.config()

const network = {
    provider: "http://localhost:26657",
    api: 'http://localhost:1317',
    bech32Prefix: "cosmos",
    nativeDenom: "stake",
    defaultTxFee: 100,
    defaultGas: 200000,
};

function getDefaultStdFee(): StdFee {
    return {
        amount: [
            {
              amount: network.defaultTxFee.toString(),
              denom: network.nativeDenom,
            },
        ],
        gas: network.defaultGas.toString(),
    }
}

async function getSigningClient(signer: OfflineSigner): Promise<SigningStargateClient> {
    return await SigningStargateClient.connectWithSigner(network.provider, signer, {prefix: network.bech32Prefix})
}

async function sendTransaction ({ to, amount, mnemonic} : {to: string, amount: number, mnemonic: string}): Promise<any> {
    const signer = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {prefix: network.bech32Prefix})
    const signingClient = await getSigningClient(signer)

    const accs = await signer.getAccounts()

    const sendAmt: Coin[] = [{ denom: network.nativeDenom, amount: amount.toString() }]

    const fee = getDefaultStdFee()

    // create tx_bytes
    let msgs = [
        {
            typeUrl: '/cosmos.bank.v1beta1.MsgSend',
            value: {
                fromAddress: accs[0].address,
                toAddress: to,
                amount: sendAmt,
            },
        },
    ]
    // sign tx
    let bodyBytes = await signingClient.sign(accs[0].address, msgs, fee, "")

    // construct payload with signed tx and broadcast it through API
    let payload = {
        tx_bytes: toBase64(TxRaw.encode(bodyBytes).finish()),
        mode: 'BROADCAST_MODE_SYNC'
    }

    let res = await axios.post(network.api + '/cosmos/tx/v1beta1/txs', payload);

    return res.data
}

(async () => {
    try {
        if (process.env.MNEMONIC == undefined) {
            throw new Error("MNEMONIC is not defined")
        }

        if (process.env.TO_ADDRESS == undefined) {
            throw new Error("TO_ADDRESS is not defined")
        }

        // will send 100000 stake to TO_ADDRESS
        let res = await sendTransaction({
            to: process.env.TO_ADDRESS,
            amount: 100000,
            mnemonic: process.env.MNEMONIC
        })

        console.log(res)
    } catch(e) {
        console.log(e)
    }
})();
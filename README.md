# Onchain merklized issuer demo

> **NOTE**: This is a demo only. Do not use it in a production environment.

### Description

This is a demo service that communicates with the [merklized on-chain issuer](https://github.com/0xPolygonID/contracts/blob/674043ddd96c1944db15079c6a00e543731724bc/contracts/examples/IdentityExample.sol) to issuer merklized credential. At first glance, it may appear to be a decentralized issuer, but it is not. In the case of a merklized credential, the service issues a credential locally, then calculates the Merkle root for the credential and saves the core claim of the credential on-chain. As a result, the user will have a MTP proof, and the issuer will manage all of their trees on-chain. After that, the user can use the MTP proof to build a ZK proof to prove the claim on-chain with an [on-chain verifier](https://devs.polygonid.com/docs/verifier/on-chain-verification/overview/).</br>
If you want to have a decentralized on-chain issuer, consider using this [contract](https://github.com/0xPolygonID/contracts/blob/main/scripts/deployBalanceCredentialIssuer.ts) and this [demo](https://github.com/0xPolygonID/onchain-nonmerklized-issuer-demo).

### Quick Start Installation

**Requirements:**
- Docker
- Docker-compose
- Ngrok

**Steps to run:**

1. Deploy the [on-chain merklized issuer contract](https://github.com/0xPolygonID/contracts/blob/main/contracts/examples/IdentityExample.sol). [Script to deploy](https://github.com/0xPolygonID/contracts/blob/main/scripts/deployIdentityExample.ts) or use the [npm command](https://github.com/0xPolygonID/contracts/blob/d308e1f586ea177005b34872992d16c3cb20e474/package.json#L60).

2. Copy `.env.example` to `.env`
    ```sh
    cp .env.example .env
    ```

3. Run `ngrok` on 8080 port.
    ```sh
    ngrok http 8080
    ```

4. Use the utility to calculate the issuerDID from the smart contract address:
    ```bash
    go run utils/convertor.go --contract_address=<ADDRESS_OF_ONCHAIN_ISSUER_CONTRACT>
    ```
    Available flags:
    - `contract_address` - contract address that will convert to did
    - `network` - network of the contract. Default: **polygon**
    - `chain` - chain of the contract. Default: **amoy**

5. Fill the `.env` config file with the proper variables:
    ```bash
    SUPPORTED_RPC="80002=<RPC_POLYGON_AMOY>"
    ISSUERS_PRIVATE_KEY="<ISSUER_DID>=<PRIVATE_KEY_OF_THE_CONTRACT_DEPLOYER>"
    EXTERNAL_HOST="<NGROK_URL>"
    ```
    `ISSUERS_PRIVATE_KEY` supports a list of issuers in the format `"issuerDID=ket,issuerDID2=key2"`

6. Use the docker-compose file:
    ```bash
    docker-compose build
    docker-compose up -d
    ```

7. Open: http://localhost:3000

## How to verify the balance claim:
1. Visit [https://verifier-demo.polygonid.me/](https://verifier-demo.polygonid.me/).
2. Choose `custom` from the drop-down menu.
3. Select the verifier based on your network:
    - For polygon mumbai: `Mumbai`
    - For polygon mainnet: `Mainnet`
    - For polygon amoy: `Amoy`
4. Fill out the form:
    - **Circuit Id**: Credential Atomic Query MTP
    - **Url**: ipfs://QmbbTKPTJy5zpS2aWBRps1duU8V3zC84jthpWDXE9mLHBX
    - **Type**: BalanceCredential
    - **Field**: balance
    - **Operator**: All operators work for the claim
    - **Value**: Set the value that you want to compare
5. Press submit.
6. Use the mobile application to verify.

## License

onchain-merklized-issuer-demo is part of the 0xPolygonID project copyright 2024 ZKID Labs AG

This project is licensed under either of

- [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE))
- [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT))

at your option.

# Signers and Addresses
> A signer scheme determines the way a transaction issigned and the mapping to a public address.

## Adresses and accounts
In Orbs platform there is no notation of a native account only of an address. Accounts or other signatrue based databases can be implemented by using the public address as the key for the database. 

## SDK Functions

#### `GetSignerAddress()`
> Returns the adddress of transaction signer. 

#### `GetCallerAddress()`
> Returns the adddress of the function caller.
* If the caller is a transaction, then CallerAddress = SignerAddress
* If the caller is a smart contractm then CallerAddress = SHA256(contract_name)
* If the caller is a system call then - TBD.


<!-- TBD 
#### `GetSignerScheme()`
#### `GetSigner()`
#### `IsSignerValid()`
#### `VerifyNetworkType(Address)`
---> 

## Usage Examples
#### Token get_my_balance()
Return BalancesDB[`GetCallerAddress()`]

#### Token transfer
If `token` <= BalancesDB[`GetCallerAddress()`] then
  BalancesDB[target] += tokens
  BalancesDB[`GetCallerAddress()`] -= tokens


## Signature schemes

#### `EdDSA01Signer`
> Ed25519 based signatrue scheme
* Parameters: 
  * network_type 
  * public_key - Ed25519 public key
* Signature = Ed25519Signature(private_key, txhash)
  * txhash = SHA256(Transaction) <!-- TBD - relation to the tarnsaction txhash - should it incldue the signature>
* Signer checks
  * network_type matches the network
  * Valid signature
* Address = {scheme, network_type, SHA256(signer)}

#### `SmartContractCaller`
> Set when a function is called by a smart contract, can't be sent in a transaction
* Parameters: 
  * network_type 
  * smart contract name
* Address = {scheme, network_type, SHA256(contract_name)}

#### `


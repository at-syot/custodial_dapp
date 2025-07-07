## TODO
- deploy contract in go: DONE
- interact contract in go: DONE
- deep-drive into smart contract - layer: WORKING
---

IPFS = InterPlantary File System
Vault service [AWS KWS, Harshikop Vault]
SLAs = Service Level agreements - a matrix indicate the service is
   [secure enough, uptime, availability, etc..]

DID = Decentalize Identifier
CID = Content Identifier - uniq hash of content - got from IPFS 

IPFS providers
:testing  web3.storage
:production Pinata, ***Infura***


## Signup flow
user signup -> system -> handover to bc-interactor-service 
bc-interactor-service -> automatically create wallet for each user
store user's keyspair (securely) && save to *vault service

## Interaction to Blockchain
*the service = bc-interactor-service

service -> construct data -> convert data struct to json like - json(struct{})
-> store that json in IPFS -> IPFS return CID *hash 
-> append hash to blockchain via Smart Contract


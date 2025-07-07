package main

import (
	"context"
	"log"
	"math/big"
	"time"

	// "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/at-syot/poc-dapp/internal/contracts/todo"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const networkURL = "http://127.0.0.1:8545"

var ks *keystore.KeyStore

func init() {
	ks = keystore.NewKeyStore("./keystore", keystore.LightScryptN, keystore.LightScryptP)
}

func main() {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, networkURL)
	if err != nil {
		log.Fatal(err)
	}

	addr := common.HexToAddress("0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199")
	balance, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		log.Fatalf("get balance err - %v", err)
	}

	log.Printf("addr: {%s}, balance: %d", addr.Hex(), balance)

	deployTodoContract(ctx, client)
	// createKeyStore()
	// generateWallet()
}

func deployTodoContract(ctx context.Context, client *ethclient.Client) {
	cid, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("retrive chainId err - %v", cid)
	}

	//   Account #0: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 (10000 ETH)
	//   Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
	//   ***Remove prefix 0x
	hhHexPk := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	hhPk, err := crypto.HexToECDSA(hhHexPk)
	if err != nil {
		log.Fatalf("derive hh pk err %v", err)
	}

	// ##### Get pub key from hh account
	hhAddr := crypto.PubkeyToAddress(hhPk.PublicKey)
	// hhAccBalance, err := client.BalanceAt(ctx, hhAddr, nil)
	// if err != nil {
	// 	log.Fatalf("get hh acc balance %v", err)
	// }
	// log.Printf("hh acc balance %v\n", hhAccBalance.Not)

	//###### Get account from KS
	// accKURL := ks.Accounts()[0].URL.String()
	// accKURL = strings.Split(accKURL, "//")[1:][0]
	// accKURL = filepath.Clean(accKURL)
	// kBytes, err := os.ReadFile(accKURL)
	// if err != nil {
	// 	log.Fatalf("retrive key from ks err - %v", err)
	// }

	// need wallet's account that gonna be the one that deploy || own
	// deploying smart contract!
	signer, err := bind.NewKeyedTransactorWithChainID(hhPk, cid)
	if err != nil {
		log.Fatal(err)
	}

	addr, tx, contx, err := todo.DeployTodo(signer, client)
	if err != nil {
		log.Fatalf("deploy contract err - %v", err)
	}
	log.Printf("Contract pending deploy: 0x%x\n", addr)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	time.Sleep(time.Second)
	log.Println("deploy:done")

	// ##### After deploy success
	// ##### Let's interact with it

	// Get next nonce
	nextN, err := client.PendingNonceAt(ctx, hhAddr)
	if err != nil {
		log.Printf("get next nonce err - %v", err)
	}
	nonce := &big.Int{}
	nonce = nonce.SetInt64(int64(nextN))

	txOpts := &bind.TransactOpts{
		From:  hhAddr,
		Nonce: nonce,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.NewPragueSigner(cid), hhPk)
		},
	}
	addTx, err := contx.Add(txOpts, "t-0")
	if err != nil {
		log.Fatalf("transact:add err - %v", err)
	}
	log.Printf("create addTx ok %+v", addTx)

	// tasks, err := contx.List(&bind.CallOpts{Pending: true})
	// if err != nil {
	// 	log.Fatalf("list todo tasks err - %v", err)
	// }
	//
	// for _, t := range tasks {
	// 	log.Printf("#task %+v\n", t)
	// }
}

func createKeyStore() {
	ks := keystore.NewKeyStore("./keystore", keystore.LightScryptN, keystore.LightScryptP)
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	_ = am

	// ks.Wallets
	for _, acc := range ks.Accounts() {
		log.Printf("fond acc %+v\n", acc)
	}

	// acc, err := ks.NewAccount("somepw")
	// if err != nil {
	// 	log.Fatalf("creating acc err - %v", err)
	// }
	// log.Printf("created acc - %s:done", acc.Address.Hex())
	log.Println()
	for _, w := range ks.Wallets() {
		log.Printf("wallet found: %+v", w)
	}
}

func generateWallet() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("generate pvk err - %v", err)
	}

	addr := crypto.PubkeyToAddress(pvk.PublicKey)
	log.Printf("wallet addr - %s", addr)
}

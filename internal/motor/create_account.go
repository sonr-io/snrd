package motor

import (
	"encoding/json"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/pkg/vault"
	rt "github.com/sonr-io/sonr/x/registry/types"
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func (mtr *MotorNode) CreateAccount(requestBytes []byte) (rtmv1.CreateAccountResponse, error) {
	// decode request
	var request rtmv1.CreateAccountRequest
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// create motor
	fmt.Printf("initializing motor... ")
	if err := initMotor(mtr); err != nil {
		return rtmv1.CreateAccountResponse{}, nil
	}
	fmt.Println("done.")

	// Request from Faucet
	fmt.Printf("requesting initial balance... ")
	err := mtr.Cosmos.RequestFaucet(mtr.Address)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, nil
	}
	fmt.Println("done.")

	// Create Initial Shards
	fmt.Printf("creating shards... ")
	deviceShard, sharedShard, recShard, unusedShards, err := mtr.Wallet.CreateInitialShards()
	if err != nil {
		return rtmv1.CreateAccountResponse{}, nil
	}
	mtr.deviceShard = deviceShard
	mtr.sharedShard = sharedShard
	mtr.recoveryShard = recShard
	mtr.unusedShards = unusedShards
	fmt.Println("done.")

	// Create the DID Document
	doc, err := did.NewDocument(mtr.DID.String())
	if err != nil {
		return rtmv1.CreateAccountResponse{}, nil
	}
	mtr.DIDDocument = doc

	// create Vault shards to make sure this works before creating WhoIs
	fmt.Printf("creating account... ")
	vc := vault.New()
	if _, err := createWhoIs(mtr); err != nil {
		return rtmv1.CreateAccountResponse{}, nil
	}
	fmt.Println("done.")

	// ecnrypt dscShard with DSC
	fmt.Printf("encrypting shards... ")
	dscShard, err := dscEncrypt(mtr.deviceShard, request.AesDscKey)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, nil
	}

	// encrypt pskShard with psk (must be generated)
	pskShard, psk, err := pskEncrypt(mtr.sharedShard)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// password protect the recovery shard
	pwShard, err := crypto.AesEncryptWithPassword(request.Password, mtr.recoveryShard)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}
	fmt.Println("done.")

	// create vault
	fmt.Printf("setting up vault... ")
	vaultService, err := vc.CreateVault(
		mtr.Address,
		mtr.unusedShards,
		mtr.DeviceID,
		dscShard,
		pskShard,
		pwShard,
	)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}
	fmt.Println("done.")

	// update DID Document
	fmt.Printf("updating WhoIs... ")
	mtr.DIDDocument.AddService(vaultService)

	// update whois
	if _, err = updateWhoIs(mtr); err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}
	fmt.Println("done.")

	fmt.Println("account created successfully.")
	return rtmv1.CreateAccountResponse{
		AesPsk:  psk,
		Address: mtr.Address,
	}, err
}

func createWhoIs(m *MotorNode) (*sdk.TxResponse, error) {
	docBz, err := m.DIDDocument.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgCreateWhoIs(m.Address, m.PubKey, docBz, rt.WhoIsType_USER)
	txRaw, err := tx.SignTxWithWallet(m.Wallet, "/sonrio.sonr.registry.MsgCreateWhoIs", msg1)
	if err != nil {
		return nil, err
	}

	resp, err := m.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	if resp.TxResponse.RawLog != "[]" {
		return nil, errors.New(resp.TxResponse.RawLog)
	}
	return resp.TxResponse, nil
}

func updateWhoIs(m *MotorNode) (*sdk.TxResponse, error) {
	docBz, err := m.DIDDocument.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgUpdateWhoIs(m.Address, docBz)
	txRaw, err := tx.SignTxWithWallet(m.Wallet, "/sonrio.sonr.registry.MsgUpdateWhoIs", msg1)
	if err != nil {
		return nil, err
	}

	resp, err := m.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	if resp.TxResponse.RawLog != "[]" {
		return nil, errors.New(resp.TxResponse.RawLog)
	}
	return resp.TxResponse, nil
}

func pskEncrypt(shard []byte) ([]byte, []byte, error) {
	key, err := crypto.NewAesKey()
	if err != nil {
		return nil, nil, err
	}

	cipherShard, err := crypto.AesEncryptWithKey(key, shard)
	if err != nil {
		return nil, key, err
	}

	return cipherShard, key, nil
}

func dscEncrypt(shard, dsc []byte) ([]byte, error) {
	if len(dsc) != 32 {
		return nil, errors.New("dsc must be 32 bytes")
	}
	return crypto.AesEncryptWithKey(dsc, shard)
}

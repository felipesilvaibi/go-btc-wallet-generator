package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// 1. **Geração de Entropia**
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}

	// 2. **Geração de Mnemônico**
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}

	// 3. **Geração de Seed**
	seed := bip39.NewSeed(mnemonic, "password")
	fmt.Println("Seed (mnemonic + password):", mnemonic)

	// 4. **Geração de Chave Mestra**
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatal(err)
	}

	// 5. **Derivação de Chave Filha**
	purposeKey, err := masterKey.Child(hdkeychain.HardenedKeyStart + 49)
	if err != nil {
		log.Fatal(err)
	}
	coinTypeKey, err := purposeKey.Child(hdkeychain.HardenedKeyStart)
	if err != nil {
		log.Fatal(err)
	}
	accountKey, err := coinTypeKey.Child(hdkeychain.HardenedKeyStart)
	if err != nil {
		log.Fatal(err)
	}
	externalKey, err := accountKey.Child(0)
	if err != nil {
		log.Fatal(err)
	}

	// Deriva uma chave filha específica para uso.
	childKey, err := externalKey.Child(0)
	if err != nil {
		log.Fatal(err)
	}

	// 6. **Geração do Endereço Público**
	pubKey, err := childKey.ECPubKey()
	if err != nil {
		log.Fatal(err)
	}
	pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
	addressPubKeyHash, err := btcutil.NewAddressPubKeyHash(pubKeyHash, &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endereço Testnet:", addressPubKeyHash.EncodeAddress())

	// 7. **Conversão para WIF**

	// Extrair a chave privada ECDSA.
	privKey, err := childKey.ECPrivKey()
	if err != nil {
		log.Fatal(err)
	}

	wif, err := btcutil.NewWIF(privKey, &chaincfg.TestNet3Params, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Chave Privada Filho WIF:", wif.String())
}

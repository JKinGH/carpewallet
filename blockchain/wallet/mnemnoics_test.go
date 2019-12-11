package wallet

import (
	"encoding/hex"

	"testing"
)


func TestGenerateMnemonics(t *testing.T) {

	m := NewMnemonicWithLanguage(DefaultEntropySize,ENGLISH)

	mnemonic,err := m.GenerateMnemonics()
	if err != nil {
		t.Errorf("Failed to GenerateMnemonics: %v",err)
	}

	t.Log("mnemonic=",mnemonic)
}


func TestGenerateSeed(t *testing.T){
	mnemonic := "trend memory raccoon escape crush nut arm alley melody spread spin cute"
	password := DefaultSeedPass
	seed ,err  := GenerateSeed(mnemonic,password)

	if err != nil {
		t.Errorf("Failed to GenerateMnemonics: %v",err)
	}

	t.Log("seed=",hex.EncodeToString(seed))
}

/*func TestDF(t *testing.T){
	seed := "4538d2e6045f770d26e253c355ac1c5a14e0f587b61412fbf96384f6dbeb254a4226bc07f7f9c2a37f3df9c24b4d0fcec5a7a91877528d036d84303b1fca4f69"
	seedBytes,_ := hex.DecodeString(seed)
	masterKey, _ := bip32.NewMasterKey(seedBytes)
	publicKey := masterKey.PublicKey()

	childKey, err := masterKey.DerivePath(BIP44PATH)
	if err != nil {
		fmt.Errorf("Failed to derive address node")
		return ""
	}
	privKey, err := childKey.ECPrivKey()
	fmt.Println("privKey="+ hex.EncodeToString(privKey.Serialize()))
	if err != nil {
		fmt.Errorf("Failed to generate privatekey")
		return ""
	}

	// Display mnemonic and keys

	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)
}*/
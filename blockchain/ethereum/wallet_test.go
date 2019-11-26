package ethereum

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"testing"
)

const (
	keystoreDir = "/Users/wujinquan/workspace/eth/"
)

//generate Mnemonics
func TestGenerateMnemonics(t *testing.T) {
	mnemonic := GenerateMnemonics()
	if mnemonic == "" {
		t.Error("GenerateMnemonics err!")
	}

	t.Log("mnemonic=", mnemonic)
}

//TestImportKeystore
func TestImportKeystore(t *testing.T) {

	ks := NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	password := "12345678"
	keyjsonStr := "{\"address\":\"ec9c88fc291ddc0e18dc321d82e29aa5454efb9d\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"525b496910610bcc48c5488edc2a1daf19cf57f553d747fb214994fd32145096\",\"cipherparams\":{\"iv\":\"5c49b2e7f13f6afc5321a43d91ee55d0\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"783f5d34ebb98b7a06b39f65f172b5e1e11a0c126a0a8e621dabe5906de2d307\"},\"mac\":\"3fc4239a550687f4b038a5740802371be876e4bb9b68a174850220b6abb140bc\"},\"id\":\"299ff3d4-14cd-49c9-aec3-561fd8ce88a8\",\"version\":3}"
	keyjson := []byte(keyjsonStr)

	err , address := ks.ImportWalletByKeystore(password,keyjson)
	if err != nil {
		t.Errorf("Failed to import account: %v", err)
	}

	t.Log("TestImportKeystore: new account=",address)
}

//Test Import PrivateKey, Save as keystore
func TestImportWalletByPrivateKey(t *testing.T){

	privkey := "51979504a2a370942b621d347262e54e796592561528822f3b7a9207d45f1c98"
	password := "12345678"

	err , address := ImportWalletByPrivateKey(keystoreDir,password,privkey)
	if err != nil {
		t.Errorf("Failed to import account: %v",err)
	}

	t.Log("TestImportWalletByPrivateKey: new account=",address)
}

//Test Export keystore by Mnemonics
func TestExportkeystoreByMnemonics(t *testing.T){

}

//Test Export Keystore
func TestExportKeystore(t *testing.T){

	address := "822a0303b4aeeb56d02838f77b35d9b2366141ea"
	password := "12345678"
	keyJson , err := ExportKeystore(keystoreDir,address,password)
	if err != nil {
		t.Errorf("Failed to export keystore: %v",err)
	}

	keyJsonStr := string(keyJson)
	t.Log("keyJsonStr=",keyJsonStr)
}

//Test Export PrivateKey by Mnemonics
func TestExportPrivateKeyByMnemonics(t *testing.T){

}

//Test Export PrivateKey by Keystore
func TestExportPrivateKeyByKeystore(t *testing.T){

	address := "1c05cb077d2e2d28bfffb73a51cca25af22bc355"
	password := "12345678"
	privateKey , err := ExportPrivateKeyByKeystore(keystoreDir,address,password)
	if err != nil {
		t.Errorf("Failed to export keystore: %v",err)
	}

	privateKeyStr := hex.EncodeToString(privateKey)
	t.Log("privateKeyStr=",privateKeyStr)
}

//修改keystore密码
func TestUpdateAccount(t *testing.T){

	oldpass := "87654321"
	newpass := "87654321"

	ks := keystore.NewKeyStore("/Users/wujinquan/workspace/eth/", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := utils.MakeAddress(ks, "72fca09899865d329b5a5b7f3beede2547374c52")
	if err != nil {
		utils.Fatalf("Could not list accounts: %v", err)
	}

	// Sign a transaction with multiple manually cancelled authorizations
	if err := ks.Unlock(account, oldpass); err != nil {
		t.Fatalf("Failed to unlock account: %v", err)
	}

	// Update the passphrase on the account created above inside the local keystore
	if err := ks.Update(account,oldpass,newpass); err != nil {
		t.Fatalf("Failed to update account: %v", err)
	}
}

//删除keystore
func TestDeleteAccount(t *testing.T){

	password := "87654321"

	ks := keystore.NewKeyStore("/Users/wujinquan/workspace/eth/", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := utils.MakeAddress(ks, "72fca09899865d329b5a5b7f3beede2547374c52")
	if err != nil {
		utils.Fatalf("Could not list accounts: %v", err)
	}

	// Delete the account updated above from the local keystore
	if err := ks.Delete(account, password); err != nil {
		t.Fatalf("Failed to delete account: %v", err)
	}

}

//
func TestImportWalletByKeystore(t *testing.T){



	password := "12345678"
	ks := keystore.NewKeyStore("/Users/wujinquan/workspace/eth/", keystore.StandardScryptN, keystore.StandardScryptP)

	// Create a new account to sign transactions with
	newaccount, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("Failed to create signer account: %v", err)
	}

	println("address="+newaccount.Address.Hex())

}





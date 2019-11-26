package ethereum

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
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

	err , address := ks.ImportKeystore(password,keyjsonStr)
	if err != nil {
		t.Errorf("Failed to import account: %v", err)
	}

	t.Log("TestImportKeystore: new account=",address)
}

//Test Import PrivateKey, Save as keystore
func TestImportPrivateKey(t *testing.T){

	ks := NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	privkey := "a7c0cf50fecdf99c570b987e21c9f63f3f152e5b74885e5f7a16dd3bbebe4d7b"
	password := "87654321"

	err , address := ks.ImportPrivateKey(password,privkey)
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

	ks := NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	address := "ec9c88fc291ddc0e18dc321d82e29aa5454efb9d"
	password := "12345678"
	keyJson , err := ks.ExportKeystore(address,password)
	if err != nil {
		t.Errorf("Failed to export keystore: %v",err)
	}

	t.Log("keyJson=",keyJson)
}

//Test Export PrivateKey by Mnemonics
func TestExportPrivateKeyByMnemonics(t *testing.T){

}

//Test Export PrivateKey by Keystore
func TestExportPrivateKeyByKeystore(t *testing.T){

	ks := NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	address := "1c05cb077d2e2d28bfffb73a51cca25af22bc355"
	password := "12345678"
	privateKey , err := ks.ExportPrivateKeyByKeystore(address,password)
	if err != nil {
		t.Errorf("Failed to export keystore: %v",err)
	}

	t.Log("privateKey=",privateKey)
}

//Test DeleteAccountByKeystore
func TestDeleteKeystoreByAddress(t *testing.T){

	ks := NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "12345678"
	address := "51c9180cf26dd1e26b481cac0f555677bfe51d95"

	if err := ks.DeleteKeystoreByAddress(address,password); err != nil {
		t.Errorf("Failed DeleteAccountByKeystore: %v",err)
	}

	t.Log("TestDeleteAccountByKeystore successful")
}

//TestUpdateKeystorePassword
func TestUpdateKeystorePassword(t *testing.T){

	ks := NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	address := "ec9c88fc291ddc0e18dc321d82e29aa5454efb9d"
	oldpass := "12345678"
	newpass := "12345678"

	if err := ks.UpdateKeystorePassword(address,oldpass,newpass); err != nil {
		t.Errorf("Failed to UpdateKeystorePassword: %v",err)
	}

	t.Log("TestUpdateKeystorePassword successful")
}






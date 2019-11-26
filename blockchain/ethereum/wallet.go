package ethereum

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	address_prefix = "0x"
)

// KeyStore manages a key storage directory on disk.
type KeyStore struct{ keystore *keystore.KeyStore }
// NewKeyStore creates a keystore for the given directory.
func NewKeyStore(keydir string, scryptN, scryptP int) *KeyStore {
	return &KeyStore{keystore: keystore.NewKeyStore(keydir, scryptN, scryptP)}
}

//Generate Mnemonics string, saprated by space, default language is EN(english)
func GenerateMnemonics() string {

	/*mn := NewMnemonicWithLanguage(ENGLISH)
	words, err := mn.GenerateMnemonic()
	if err != nil {
		return ""
	}*/
	//return words
	return "situate job interest junior demand wear abstract elephant draw suit right sell"
}

//Import Wallet By Mnemonics
//return newaddress
func ImportWalletByMnemonics(mnemonics string) string{

	return "3057a0f2a3dab4eca9d93b2c78ffccc30e4482fd"
}

//Import Wallet By Keystore  //wjq 增加去重判断:keystore.Find()
func (ks *KeyStore) ImportKeystore(password string, keyjson string) (error ,string){

	// Import back the account we've exported (and then deleted) above with yet
	// again a fresh passphrase
	newAccount, err := ks.keystore.Import(common.CopyBytes([]byte(keyjson)), password, password)
	if err != nil {
		return err, ""
	}
	return nil , newAccount.Address.Hex()
}

//Import PrivateKey, Save as keystore
func (ks *KeyStore) ImportPrivateKey(password, privateKey string ) (error ,string){

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err, ""
	}

	newAccount,err := ks.keystore.ImportECDSA(privKey,password)
	if err != nil {
		return err, ""
	}
	return nil , newAccount.Address.Hex()
}

//Export keystore by Mnemonics
func (ks *KeyStore) ExportkeystoreByMnemonics(){

}

//Export keystore by keystore
func (ks *KeyStore) ExportKeystore(address, password string )(string,error) {

	account, err := utils.MakeAddress(ks.keystore, address)
	if err != nil {
		return "", err
	}

	// Export the newly created account with a different passphrase. The returned
	// data from this method invocation is a JSON encoded, encrypted key-file
	keyJson,err := ks.keystore.Export(account,password,password)
	if err != nil {
		return "", err
	}

	return string(keyJson),nil
}

//Export PrivateKey by Mnemonics
func (ks *KeyStore) ExportPrivateKeyByMnemonics(){

}

//Export PrivateKey by keystore
func (ks *KeyStore) ExportPrivateKeyByKeystore(address, password string )(string,error) {

	keyJson , err := ks.ExportKeystore(address,password)
	if err != nil {
		return "", err
	}

	key, err := keystore.DecryptKey([]byte(keyJson), password)
	//seckey := math.PaddedBigBytes(key.PrivateKey.D, key.PrivateKey.Params().BitSize/8)
	//fmt.Println("seckey="+ hex.EncodeToString(seckey))
	if err != nil {
		return "", err
	}
	privateKey := crypto.FromECDSA(key.PrivateKey)

	return hex.EncodeToString(privateKey), nil
}

//Delete  Mnemonics
func DeleteMnemonics(){

}

//Delete Keystore By Address
func (ks *KeyStore) DeleteKeystoreByAddress(address, password string ) error {

	account, err := utils.MakeAddress(ks.keystore, address)
	if err != nil {
		utils.Fatalf("Could not list accounts: %v", err)
	}

	// Delete the account updated above from the local keystore
	if err := ks.keystore.Delete(account, password); err != nil {
		return err
	}

	return nil
}


//Update Keystore password
func (ks *KeyStore) UpdateKeystorePassword(address, oldpass , newpass string ) error {

	account, err := utils.MakeAddress(ks.keystore, address)
	if err != nil {
		utils.Fatalf("Could not list accounts: %v", err)
	}

	// Update the passphrase on the account created above inside the local keystore
	if err := ks.keystore.Update(account,oldpass,newpass); err != nil {
		return err
	}

	return nil
}
package ethereum

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
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
func (ks *KeyStore)ImportWalletByKeystore(password string, keyjson []byte) (error ,string){

	// Import back the account we've exported (and then deleted) above with yet
	// again a fresh passphrase
	newAccount, err := ks.keystore.Import(keyjson, password, password)
	if err != nil {
		return err, ""
	}
	return nil , newAccount.Address.Hex()
}

//Import PrivateKey, Save as keystore  //wjq 已有去重判断
func ImportWalletByPrivateKey(keystoreDir,password, privateKey string ) (error ,string){

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err, ""
	}

	newAccount,err := ks.ImportECDSA(privKey,password)
	if err != nil {
		return err, ""
	}
	return nil , newAccount.Address.Hex()
}

//Export keystore by Mnemonics
func ExportkeystoreByMnemonics(){

}

//Export keystore by keystore
func ExportKeystore(keystoreDir,address, password string )([]byte,error) {

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := utils.MakeAddress(ks, address)
	if err != nil {
		return nil, err
	}

	// Export the newly created account with a different passphrase. The returned
	// data from this method invocation is a JSON encoded, encrypted key-file
	keyJson,err := ks.Export(account,password,password)
	if err != nil {
		return nil, err
	}

	return keyJson,nil
}

//Export PrivateKey by Mnemonics
func ExportPrivateKeyByMnemonics(){

}

//Export PrivateKey by keystore
func ExportPrivateKeyByKeystore(keystoreDir,address, password string )([]byte,error) {

	keyJson , err := ExportKeystore(keystoreDir,address,password)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(keyJson, password)
	//seckey := math.PaddedBigBytes(key.PrivateKey.D, key.PrivateKey.Params().BitSize/8)
	//fmt.Println("seckey="+ hex.EncodeToString(seckey))
	if err != nil {
		return nil, err
	}
	privateKey := crypto.FromECDSA(key.PrivateKey)

	return privateKey, nil
}
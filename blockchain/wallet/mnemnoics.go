package wallet

import (
	"errors"
	"github.com/JKinGH/carpewallet/blockchain/commons"
	"github.com/tyler-smith/go-bip39/wordlists"
)

// List Mnemonic language support
const (
	ENGLISH  string = "EN"
	JAPANESE        = "JP"
	FRENCH          = "FR"
	ITALIAN         = "IT"
	KOREAN          = "KR"
	SPANISH         = "ES"
	ZH_CN			= "ZH_CN"
	ZH_TW			= "ZH_TW"
)

const(
	DefaultEntropySize = 128

	// Default seed pass. it used to generate seed from mnemonic( BIP39 ). Don't change if determined
	DefaultSeedPass = ""
)

type Mnemonic struct {
	EntropySize int  //单词个数 = EntropySize/8
	Language    string //语言
}

func  NewMnemonicWithLanguage(entropySize int,language string) *Mnemonic {
	commons.SetWordList(loadWordList(language))
	return &Mnemonic{EntropySize: entropySize, Language: language}
}

//Generate Mnemonics string, saprated by space, default language is EN(english)
func (m *Mnemonic ) GenerateMnemonics() (string ,error){

	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := commons.NewEntropy(m.EntropySize)
	if err != nil {
		return "",err
	}
	mnemonic, err := commons.NewMnemonic(entropy)
	if err != nil {
		return "",err
	}

	return mnemonic,nil
}

// Generate seed from mnemonic and pass
func GenerateSeed(mnemonics ,password string) ([]byte, error) {
	if !commons.IsMnemonicValid(mnemonics) {
		return nil, errors.New("invalidate mnemonic")
	}
	return commons.NewSeed(mnemonics, password), nil
}



// loadWordList returns word lists base on language setting in the configuration
func loadWordList(language string) []string {
	switch language {
	case JAPANESE:
		return wordlists.Japanese
	case ITALIAN:
		return wordlists.Italian
	case KOREAN:
		return wordlists.Korean
	case SPANISH:
		return wordlists.Spanish
	case FRENCH:
		return wordlists.French
	case ZH_CN :
		return wordlists.ChineseSimplified
	case ZH_TW :
		return wordlists.ChineseTraditional
	default:
		return wordlists.English
	}
}
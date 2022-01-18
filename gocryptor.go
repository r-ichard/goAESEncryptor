package main

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var filePath string
	var decrypt bool
	var encrypt bool
	var keyPath string
	var ivString string
	var tagString string
	var aadString string
	flag.StringVar(&filePath, "f", "", "The path to the file you want to encrypt/decrypt")
	flag.BoolVar(&decrypt, "d", false, "Used to tell the Command-line that you want to decrypt the file")
	flag.BoolVar(&encrypt, "e", false, "Used to tell the Command-line that you want to encrypt the file")
	flag.StringVar(&keyPath, "p", "", "The key file used to encrypt/decrypt the file (16/24/32 bytes) - Base64 encoded expected")
	flag.StringVar(&ivString, "i", "", "The identification vector (should change for every encryption) - Base64 Expected")
	flag.StringVar(&tagString, "t", "", "The tag - Base64 Expected")
	flag.StringVar(&aadString, "a", "", "The additional authentication data string")
	flag.Parse()

	if decrypt {
		DecryptFile(filePath, keyPath, ivString, tagString, aadString)
	} else if encrypt {
		EncryptFile(filePath, keyPath, ivString, tagString, aadString)
	}
}

type encryptionInfo struct {
	file []byte
	key  []byte
	iv   []byte
	tag  []byte
	aad  []byte
}

func EncryptFile(filePath string, keyPath string, ivString string, tagString string, aadString string) {
	ei, err := transformAndPrepare(filePath, keyPath, ivString, tagString, aadString)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new AES cipher
	block, err := aes.NewCipher(ei.key)

	nonceSize := len(ei.iv)
	gcm, err := cipher.NewGCMWithNonceSize(block, nonceSize)
	if err != nil {
		log.Fatal(err)
	}

	encryptedFileWithTag := gcm.Seal(nil, ei.iv, ei.file, ei.aad)

	tagFile := encryptedFileWithTag[len(encryptedFileWithTag)-gcm.Overhead():]
	fmt.Println("******TAG BELOW******")
	fmt.Println(b64.StdEncoding.EncodeToString(tagFile))
	fmt.Println("******TAG ABOVE******")

	encryptedFilePath := filePath + ".encrypted"
	// Now, we write the encryption to the file
	ioutil.WriteFile(encryptedFilePath, encryptedFileWithTag, 0777)
}

func DecryptFile(filePath string, keyPath string, ivString string, tagString string, aadString string) {

	ei, err := transformAndPrepare(filePath, keyPath, ivString, tagString, aadString)
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(ei.key)
	if err != nil {
		log.Fatal(err)
	}

	nonceSize := len(ei.iv)
	gcm, err := cipher.NewGCMWithNonceSize(block, nonceSize)
	if err != nil {
		log.Panic(err)
	}

	encryptedFileWithTag := append(ei.file, ei.tag...)

	decryptedFile, err := gcm.Open(nil, ei.iv, encryptedFileWithTag, ei.aad)
	if err != nil {
		log.Panic(err)
	}
	decryptedFilePath := filePath + ".decrypted"
	err = ioutil.WriteFile(decryptedFilePath, decryptedFile, 0777)
	if err != nil {
		log.Panic(err)
	}
}

func transformAndPrepare(filePath string, keyPath string, ivString string, tagString string, aadString string) (encryptionInfo, error) {
	encryptedFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return encryptionInfo{}, err
	}

	keyFile, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return encryptionInfo{}, err
	}
	key, err := b64.StdEncoding.DecodeString(string(keyFile))
	if err != nil {
		return encryptionInfo{}, err
	}

	nonce, err := b64.StdEncoding.DecodeString(ivString)
	if err != nil {
		return encryptionInfo{}, err
	}

	var authTag []byte

	if tagString != "" { // not needed in encryption mode
		authTag, err = b64.StdEncoding.DecodeString(tagString)
		if err != nil {
			return encryptionInfo{}, err
		}
	}

	aad := []byte(aadString)

	ei := encryptionInfo{
		file: encryptedFile,
		key:  key,
		iv:   nonce,
		tag:  authTag,
		aad:  aad,
	}

	return ei, nil
}

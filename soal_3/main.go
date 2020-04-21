package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

const (
	aesSecretKey = "1234privy5678"
	iv = ""
	stringToEncrypt = "Selamat Datang"
)

func main() {
	resultEncrypt := AES256_CBC_Encrypt(aesSecretKey, iv, stringToEncrypt)
	fmt.Println("encrypt the string('"+stringToEncrypt+"') \t\t\t: "+resultEncrypt)

	resultDescrypt := AES256_CBC_Decrypt(aesSecretKey, iv, resultEncrypt)
	fmt.Println("decrypt the encode AES256('"+resultEncrypt+"') \t: "+resultDescrypt)
}

// AES256_CBC_Encrypt is function to encrypt
func AES256_CBC_Encrypt(secretKey, IV, textToBeEncrypt string) string{
	var (
		byteSecretKey 	= make([]byte, 32) // to init AES 256
		byteIV 			= make([]byte, 16)
	)
	// init AES256
	copy(byteSecretKey[:], []byte(secretKey))
	chiperBlock, err := aes.NewCipher(byteSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	// add padding in to the byte inside text that will be encrypt
	bPlaintext := PKCS5Padding([]byte(textToBeEncrypt), aes.BlockSize, len(textToBeEncrypt))
	copy(byteIV[:], []byte(IV))
	// create mode CBC
	blockMode := cipher.NewCBCEncrypter(chiperBlock, []byte(byteIV))
	// make variable output which the length of byte must be same with plainText or string that will be encrypt
	output := make([]byte, len(bPlaintext))
	blockMode.CryptBlocks(output, bPlaintext)
	return base64.StdEncoding.EncodeToString(output)
}

// PKCS5Padding is a schema to pad cleartext to be multiples of 8-byte blocks
func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	dif 	:= len(ciphertext)%blockSize
	padding := blockSize - dif
	// the padding
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	// added into the original text that will be encrypted
	return append(ciphertext, padtext...)
}

func AES256_CBC_Decrypt(secretKey, iv, stringToDecode string) string{
	var (
		byteSecretKey 	= make([]byte, 32) // init that secret is AES256
		byteIV 			= make([]byte, 16)
	)
	copy(byteSecretKey[:], []byte(secretKey))

	byteToDecode, err := base64.StdEncoding.DecodeString(stringToDecode)
	if err != nil {
		log.Fatal(err)
	}
	//create AES256
	chiperBlock, err := aes.NewCipher(byteSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	if len(byteToDecode) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		log.Fatal(err)
	}
	copy(byteIV[:], []byte(iv))
	// create mode CBC
	blockMode := cipher.NewCBCDecrypter(chiperBlock, byteIV)
	output := make([]byte, len(byteToDecode))
	// descrypt
	blockMode.CryptBlocks(output, byteToDecode)
	// remove whitepace
	outputWithoutSpace := bytes.TrimSpace(output)
	return string(outputWithoutSpace)
}

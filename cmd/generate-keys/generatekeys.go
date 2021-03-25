package generatekeys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
)

type GenerateRSAArgs struct {
	PrivateKeyFilename string
	PublicKeyFilename  string
	KeySize            int
	ReplaceKeys        bool
}

var errFileMustBePresent = errors.New("fileMustBePresentError")

func FileMustBePresentError(filename string) error {
	return fmt.Errorf(`%w: file "%v" must not be present`, errFileMustBePresent, filename)
}

func GenerateRSA(args GenerateRSAArgs) error {
	if !args.ReplaceKeys && fileExists(args.PrivateKeyFilename) {
		return FileMustBePresentError(args.PrivateKeyFilename)
	}

	if !args.ReplaceKeys && fileExists(args.PublicKeyFilename) {
		return FileMustBePresentError(args.PublicKeyFilename)
	}

	key, err := rsa.GenerateKey(rand.Reader, args.KeySize)
	checkError(err)

	savePrivateKey(args.PrivateKeyFilename, key)
	log.Println("Generated private key", args.PrivateKeyFilename)

	savePublicPEMKey(args.PublicKeyFilename, &key.PublicKey)
	log.Println("Generated public key", args.PublicKeyFilename)

	return nil
}

func savePrivateKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)

	privateKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)

	err = outFile.Close()
	checkError(err)
}

func savePublicPEMKey(filename string, pubKey *rsa.PublicKey) {
	pemKey := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pubKey),
	}

	pemFile, err := os.Create(filename)
	checkError(err)

	err = pem.Encode(pemFile, pemKey)
	checkError(err)

	err = pemFile.Close()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

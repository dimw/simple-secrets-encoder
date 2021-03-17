package generate_keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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

func GenerateRSA(args GenerateRSAArgs) error {
	if !args.ReplaceKeys && fileExists(args.PrivateKeyFilename) {
		return fmt.Errorf(`file "%v" must not be present`, args.PrivateKeyFilename)
	}

	if !args.ReplaceKeys && fileExists(args.PublicKeyFilename) {
		return fmt.Errorf(`file "%v" must not be present`, args.PublicKeyFilename)
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
	var pemKey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pubKey),
	}

	pemFile, err := os.Create(filename)
	checkError(err)
	defer pemFile.Close()

	err = pem.Encode(pemFile, pemKey)
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

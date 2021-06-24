package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/tzdq/go-utils/file"
)

// RSAGenKeyToFile Generate public and private keys and save them in files
func RSAGenKeyToFile(bits int, publicKeyPath, privateKeyPath string) error {
	// Generate private keys
	// 1、Use the GenerateKey method in RSA to generate a private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	// 2、Serialize the obtained RAS private key into ASN.1 DER encoded string through X509 standard
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	//3、Set the private key string into the pem format block
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateStream,
	}
	//4、Encode the set data through pem and write to disk file
	privateFile, err := file.CreateFile(privateKeyPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = privateFile.Close()
	}()
	err = pem.Encode(privateFile, block)
	if err != nil {
		return err
	}

	// Generate public keys
	publicKey := &privateKey.PublicKey
	publicStream, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicStream,
	}
	publicFile, err := file.CreateFile(publicKeyPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = publicFile.Close()
	}()
	err = pem.Encode(publicFile, block)
	if err != nil {
		return err
	}
	return nil
}

// RSAGenKey Generate public and private keys. Return value is public key,private key,error
func RSAGenKey(bits int) ([]byte, []byte, error) {
	// Generate private keys
	// 1、Use the GenerateKey method in RSA to generate a private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	// 2、Serialize the obtained RAS private key into ASN.1 DER encoded string through X509 standard
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3、Set the private key string into the pem format block
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateStream,
	}
	// 4、Encode the set data through pem
	prvKey := pem.EncodeToMemory(block)

	// Generate public keys
	publicKey := &privateKey.PublicKey
	publicStream, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicStream,
	}
	pubKey := pem.EncodeToMemory(block)
	return pubKey, prvKey, nil
}

// RSAEncrypt Use the public key to encrypt the plaintext
func RSAEncrypt(plaintext, publicKey []byte) ([]byte, error) {
	// decrypt the public key in pem format
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New(sErrPublicKeyErr)
	}
	// parse the public key
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// type assertion
	pub := pubInterface.(*rsa.PublicKey)
	// encrypt
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plaintext)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// RSAEncryptFromFile Encrypt the plaintext after reading the public key from the file
func RSAEncryptFromFile(plaintext []byte, pubKeyPath string) ([]byte, error) {
	publicKey, err := file.ReadFile(pubKeyPath)
	if err != nil {
		return nil, err
	}
	return RSAEncrypt(plaintext, publicKey)
}

// RSADecrypt Use the private key to decrypt the ciphertext
func RSADecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New(sErrPrivateKeyErr)
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, private, ciphertext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// RSADecryptFromFile Decrypt the ciphertext after reading the private key from the file
func RSADecryptFromFile(ciphertext []byte, priKeyPath string) ([]byte, error) {
	priKey, err := file.ReadFile(priKeyPath)
	if err != nil {
		return nil, err
	}
	return RSADecrypt(ciphertext, priKey)
}

// RSAVerySign Use the public key to verify whether the signed message has been tampered
func RSAVerySign(data, signature, publicKey []byte) bool {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false
	}
	pubKey := pubInterface.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, HashBytes(data, HtSha256), signature)
	if err != nil {
		return false
	}
	return true
}

// RSAVerySignFromFile Use the public key (read from the file) to verify whether the signed message has been tampered
func RSAVerySignFromFile(data, signature []byte, pubKeyPath string) bool {
	publicKey, err := file.ReadFile(pubKeyPath)
	if err != nil {
		return false
	}
	return RSAVerySign(data, signature, publicKey)
}

// RSASign Use private key to sign information
func RSASign(data, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New(sErrPrivateKeyErr)
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA256, HashBytes(data, HtSha256))
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// RSASignFromFile Use private key to sign information After reading the private key from the file
func RSASignFromFile(data []byte, priKeyPath string) ([]byte, error) {
	privateKey, err := file.ReadFile(priKeyPath)
	if err != nil {
		return nil, err
	}
	return RSASign(data, privateKey)
}

/*
NAME: db201110-d875-4133-9709-2732a47f252f.go
RULE: Protect against ransomware
CREATED: 2023-01-03
*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
	"os/user"
	"path/filepath"

	Endpoint "github.com/preludeorg/test/endpoint"
)

func Encrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func test() {
	usr, _ := user.Current()
	println("[*] Targeting user: " + usr.Username)

	println("[*] Generating sample files...")
	Endpoint.Write(filepath.Join(usr.HomeDir, "one.txt"), make([]byte, 1000))
	Endpoint.Write(filepath.Join(usr.HomeDir, "two.xlsx"), make([]byte, 7500000))
	Endpoint.Write(filepath.Join(usr.HomeDir, "three.pdf"), make([]byte, 5500))

	key, err := GenerateKey()
	if err != nil {
		println(err)
		os.Exit(1)
	}
	println("[+] Generated an encryption key")

	arr := [3]string{"one.txt", "two.xlsx", "three.pdf"}
	for _, name := range arr {
		println("[+] Starting encryption process: " + name)
		bytesRead := Endpoint.Read(filepath.Join(usr.HomeDir, name))

		encryptBytes, err := Encrypt(key, bytesRead)
		if err != nil {
			println(err)
		}

		Endpoint.Write(filepath.Join(usr.HomeDir, name), encryptBytes)
	}

	os.Exit(100)
}

func clean() {
	usr, _ := user.Current()
	arr := [3]string{"one.txt", "two.xlsx", "three.pdf"}

	for _, name := range arr {
		Endpoint.Remove(filepath.Join(usr.HomeDir, name))
	}
	os.Exit(100)
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		clean()
	} else {
		test()
	}
}

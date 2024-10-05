package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const key = "1234567891234567"

func GetConnectionAtrribs(connectionString []byte) (string, string, error) {
	type data struct {
		IP   string `json:"ip"`
		Name string `json:"name"`
		Node string `json:"node"`
	}

	var d data
	err := json.Unmarshal(connectionString, &d)
	if err != nil {
		return "", "", err
	}

	d.IP, err = decryptMessage([]byte(key), d.IP)
	if err != nil {
		return "", "", err
	}
	d.Name, err = decryptMessage([]byte(key), d.Name)
	if err != nil {
		return "", "", err
	}
	d.Node, err = decryptMessage([]byte(key), d.Node)
	if err != nil {
		return "", "", err
	}

	fmt.Println(d.IP, d.Name, d.Node)

	cs := fmt.Sprintf("opc.tcp://%s/%s", d.IP, d.Name)

	return cs, d.Node, nil
}

func decryptMessage(key []byte, message string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

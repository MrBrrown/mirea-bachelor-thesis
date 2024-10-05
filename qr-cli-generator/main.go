package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"

	qr "github.com/skip2/go-qrcode"
)

type ConnectionStirng struct {
	Ip   string
	Name string
	Node string
}

func (cs ConnectionStirng) toJSON() string {
	return fmt.Sprintf(`{
        "ip": "%s",
        "name": "%s",
        "node": "%s"
    }`, cs.Ip, cs.Name, cs.Node)
}

func (cs *ConnectionStirng) GenerateQR(key []byte) string {
	ex, err := os.Executable()
	if err != nil {
		return "Ошибка создания файлв: " + err.Error()
	}
	fileName := filepath.Dir(ex) + "/qrcode.png"

	err = Encrypt(cs, key)
	if err != nil {
		return "Ошибка создания файлв: " + err.Error()
	}

	err = qr.WriteFile(cs.toJSON(), qr.Medium, 256, fileName)
	if err != nil {
		return "Ошибка создания файлв: " + err.Error()
	}

	return fmt.Sprintf("Путь к файлу: %s", fileName)
}

func Encrypt(cs *ConnectionStirng, key []byte) error {
	var err error
	cs.Ip, err = encrypt(cs.Ip, key)
	if err != nil {
		return err
	}

	cs.Name, err = encrypt(cs.Name, key)
	if err != nil {
		return err
	}
	cs.Node, err = encrypt(cs.Node, key)
	if err != nil {
		return err
	}

	fmt.Println(cs.Ip)
	fmt.Println(cs.Name)
	fmt.Println(cs.Node)

	return nil
}

func encrypt(text string, key []byte) (string, error) {
	byteMsg := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func main() {
	connectionString := ConnectionStirng{}
	key := ""
	fmt.Println("Ведите секретный ключ:")
	fmt.Scan(&key)
	if len(key) != 16 {
		return
	}

	fmt.Println("Введите IP адрес сервера")
	fmt.Scan(&connectionString.Ip)

	fmt.Println("Введите имя сервера")
	fmt.Scan(&connectionString.Name)

	fmt.Println("Введите имя узла")
	fmt.Scan(&connectionString.Node)

	fmt.Println(connectionString.GenerateQR([]byte(key)))
}

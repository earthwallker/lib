package strlib
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func TestEnDecrypt() {
	key := []byte("example key 1234") // 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256

	text := "Hello, World!"
	fmt.Println("Original Text:", text)

	ciphertext, err := Encrypt(key, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted:", ciphertext)

	decrypted, err := Decrypt(key, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted Text:", decrypted)
}

//对字符串进行加密和解密。
//请注意，这里的key长度必须是16、24或32字节，对应AES-128、AES-192或AES-256。
//字符串进行AES加密，并返回Base64编码的密文。
//decrypt 函数接受Base64编码的密文并解密为原始字符串。
func Encrypt(key []byte, text string) (string, error) {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return string(ciphertext), nil
}


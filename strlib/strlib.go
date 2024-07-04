package strlib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// func TestEnDeCrypt() {
// 	key := []byte("dlgdabrnabrndlgd") // 需要 16, 24 或 32 字节的密钥
// 	plaintext := "Hello, World!"

// 	// 加密
// 	ciphertext, err := Encrypt(key, plaintext)
// 	if err != nil {
// 		fmt.Println("加密错误:", err)
// 		return
// 	}
// 	fmt.Println("加密后的文本:", ciphertext)

// 	// 解密
// 	decryptedText, err := Decrypt(key, ciphertext)
// 	if err != nil {
// 		fmt.Println("解密错误:", err)
// 		return
// 	}
// 	fmt.Println("解密后的文本:", decryptedText)
// }

func newGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

//对字符串进行加密和解密。
//请注意，这里的key长度必须是16、24或32字节，对应AES-128、AES-192或AES-256。
//字符串进行AES加密，并返回Base64编码的密文。
//decrypt 函数接受Base64编码的密文并解密为原始字符串。
func Encrypt(key []byte, plaintext string) (string, error) {
	// 生成随机的 nonce
	gcm, err := newGCM(key)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// 返回 Base64 编码的密文
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	gcm, err := newGCM(key)
	if err != nil {
		return "", err
	}

	// 提取 nonce
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("密文太短")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 解密
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
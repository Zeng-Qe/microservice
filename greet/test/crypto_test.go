package test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var serverPrivateKey *rsa.PrivateKey
var publickey *rsa.PublicKey

func init() {
	// 在实际应用中，应该从安全的密钥存储（如 HSM）中加载私钥
	var err error
	serverPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("无法生成私钥: %v", err)
	}
	publickey = &serverPrivateKey.PublicKey
	fmt.Println("公钥：", publickey)
	// fmt.Println("私钥：", serverPrivateKey)

}

func main() {
	// username := "testuser"
	// receivedEncryptedPassword := "base64EncodedEncryptedPassword" // 这应该是从前端接收的加密密码
	receivedSalt := []byte("somesalt") // 在实际应用中，这可能不需要，因为 bcrypt 会生成盐

	// 加密
	// Decrypt 解密
	decrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publickey, receivedSalt, nil)
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	fmt.Println(decrypted)
	password := base64.StdEncoding.EncodeToString(decrypted)
	fmt.Println(password)

	// 解密
	decryptedPwd, err := DecryptPassword(password)
	if err != nil {
		log.Fatalf("密码解密失败: %v", err)
	}

	hashedPwd, err := HashPassword(decryptedPwd, receivedSalt)
	if err != nil {
		log.Fatalf("密码哈希失败: %v", err)
	}
	fmt.Println(hashedPwd)
	// fmt.Println(username)

}

func DecryptPassword(encryptedPassword string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedPassword)
	if err != nil {
		return nil, fmt.Errorf("解码密文失败: %v", err)
	}

	// Decrypt 解密
	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, serverPrivateKey, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("解密失败: %v", err)
	}
	fmt.Println(decrypted)
	fmt.Println("解密密码值：", string(decrypted))

	return decrypted, nil
}

func HashPassword(decryptedPassword []byte, salt []byte) ([]byte, error) {
	// 注意：在Go中，bcrypt.GenerateFromPassword 已经包含了salt的生成
	// 因此，这里我们直接使用 bcrypt.GenerateFromPassword
	hashedPassword, err := bcrypt.GenerateFromPassword(decryptedPassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("密码哈希失败: %v", err)
	}
	return hashedPassword, nil
}

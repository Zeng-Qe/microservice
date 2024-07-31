package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// 假设这是从安全存储中获取的服务器私钥
var serverPrivateKey *rsa.PrivateKey

func init() {
	// 在实际应用中，应该从安全的密钥存储（如 HSM）中加载私钥
	var err error
	serverPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("无法生成私钥: %v", err)
	}
	publickey := &serverPrivateKey.PublicKey
	fmt.Println("公钥：", publickey)
	fmt.Println("私钥：", serverPrivateKey)

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

func Encrypt(publicKey *rsa.PublicKey) {
	// 假设已经有了一个 RSA 公钥
	// publicKey := getRSAPublicKey() // 这应该是从之前生成的密钥对中获取的公钥
	// 要加密的数据
	message := []byte("这是一个需要加密的秘密消息")
	// 使用公钥进行加密
	encryptedMessage, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	// 加密后的数据
	log.Printf("加密后的消息: %x\n", encryptedMessage)
}

func Decrypt(privateKey *rsa.PrivateKey, encryptedMessage []byte) {
	// 假设已经有了一个 RSA 私钥
	// privateKey := getRSAPrivateKey() // 这应该是从之前生成的密钥对中获取的私钥

	// 假设这是已加密的数据
	// encryptedMessage := getEncryptedMessage()

	// 使用私钥进行解密
	decryptedMessage, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedMessage, nil)
	if err != nil {
		log.Fatalf("解密失败: %v", err)
	}

	// 解密后的数据
	log.Printf("解密后的消息: %s\n", string(decryptedMessage))

}

func SignPKCS(privateKey *rsa.PrivateKey) {
	// 假设已经有了一个 RSA 私钥
	// privateKey := GetRSAPrivateKey()

	// 准备要签名的数据
	data := []byte("这是一个需要签名的重要消息")
	hashedData := sha256.Sum256(data)

	// 使用私钥创建签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashedData[:])
	if err != nil {
		log.Fatalf("签名失败: %v", err)
	}

	// 输出签名
	log.Printf("生成的签名: %x\n", signature)

}

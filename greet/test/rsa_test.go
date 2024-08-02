package test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// var password = "eDdka1FFVlR3enZ4ZVZLdVljQWt0RVNNT25mUHdpTU5tbXk0dkYzSThTcTUzRWN6VWl3STdIUzRTZTM1MXFROTl5V2xkbmtwOTQwZDVpZVl6b2NwZVF0RXNSc21aSmZ3a3RES3BwbVpWRURLNGJzZHVjSFhXTzd2eDY3VmFsQThjbjEwSnp2d0xNKzZVeHpiK2VnTTJqRUd6aFhTMGZEQ0ZmcEJPSEdmb1FMV1l5eTN3RWtZc2lFUzlxWjZ4WTlZbEN4Y2dibk9jeURuVFV0N3RlalM0UFMzR3BpMnFEWHRLWlFPVkpndEJqaTNWb1F2dG5yS3VpcURpSFhyaTdXVTRSY3BDbGcrb1UvLzcyc0FyN0huRkp1TjdWZHozSitmVFBWdWdiL0k2enhPQjhVVldsOUhxcit3UVkrZy9QckZZSWJ3RHVFSlBpVkpwbW5LUWROOUVRPT0="

var RSAPrivateKey *rsa.PrivateKey
var RSAPublicKey *rsa.PublicKey

// 初始化私钥
func RsaDecryptInit(filePath string) (err error) {
	key, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New("加载私钥错误1：" + err.Error())
	}
	block, _ := pem.Decode(key)
	if block == nil {
		return errors.New("加载私钥错误2：")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return errors.New("加载私钥错误3：" + err.Error())
	}
	RSAPrivateKey = privateKey

	return err
}

// 初始化公钥
func RsaDecryptPublicInit(filePath string) (err error) {
	key, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New("加载私钥错误1：" + err.Error())
	}
	block, _ := pem.Decode(key)
	if block == nil {
		return errors.New("加载私钥错误2：")
	}

	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return errors.New("加载私钥错误3：" + err.Error())
	}
	publicKey, flag := publicInterface.(*rsa.PublicKey)
	if flag == false {
		return errors.New("加载私钥错误4：" + err.Error())
	}

	fmt.Println("公钥：", publicKey)
	RSAPublicKey = publicKey
	return err
}

// 解密
func DecryptPasswordV2(encryptedPassword string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedPassword)
	if err != nil {
		return nil, fmt.Errorf("解码密文失败: %v", err)
	}
	fmt.Println("私钥：", string(ciphertext))

	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, RSAPrivateKey, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("解密失败: %v", err)
	}
	fmt.Println("私钥：", string(decrypted))

	return decrypted, nil
}

// 生成 hash
func HashPasswordV2(decryptedPassword []byte, salt []byte) (password []byte, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(decryptedPassword, bcrypt.DefaultCost)
	if err != nil {
		return password, fmt.Errorf("密码哈希失败: %v", err)
	}
	return hashedPassword, err
}

// 加密
func Encrypt() string {
	message := []byte("123456salt")
	encryptedMessage, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, RSAPublicKey, message, nil)
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	log.Printf("加密后的消息: %x\n", encryptedMessage)
	encryptPassword := base64.StdEncoding.EncodeToString(encryptedMessage)
	return encryptPassword
}

func TestRSA(t *testing.T) {
	RsaDecryptInit("../rsa_key/private_key.pem")
	RsaDecryptPublicInit("../rsa_key/public.pem")

	password := Encrypt()
	fmt.Println(password)

	passwords, err := DecryptPassword(password)
	if err != nil {
		return
	}
	fmt.Println("私钥：", string(passwords))

	// 加密生成 hash
	// passwordss, err := HashPassword(passwords, []byte("salt"))
	passwordss := strings.TrimRight(string(passwords), "salt")

	fmt.Println("私钥：", string(passwordss))

	t.Logf("%v", passwordss)
}

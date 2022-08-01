package AES

import (
	"log"
	"testing"
	"time"
)

func TestAes(t *testing.T) {
	//加密
	str, _ := EncryptByAes([]byte("abcdef"))
	//解密
	str1, _ := DecryptByAes(str)
	//打印
	log.Printf(" 加密：%v\n 解密：%s\n ",
		str, str1,
	)
}

func TestAesFile(t *testing.T) {
	startTime := time.Now()
	EncryptFile("res", "t.dat")
	log.Printf("耗时：%v", time.Since(startTime))
}

func TestDecFile(t *testing.T) {
	startTime := time.Now()
	DecryptFile("encryptFile_t.dat", "dec")
	log.Printf("耗时：%v", time.Since(startTime))
}

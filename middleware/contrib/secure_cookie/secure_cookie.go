package secure_cookie

// DEMO
//import (
//	"Gin_API_Framework/middleware/contrib/secure_cookie"
//)

//func DemoHandler(c *gin.Context) {

//     Set Secure Cookie
//     secure_cookie.SetSecureCookie(
//		c, "user_token","1", 222, "/asd","*",true,true)

//     Get Secure Cookie
//     v, _ := secure_cookie.GetSecureCookie(c, "user_token",10)

//
//	c.JSON(http.StatusOK, gin.H{
//		"status":  "success",
//		"is_created": success,
//		"cookie" : v,
//	})
//
//}

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


import (
	"errors"
	"encoding/hex"
	"crypto/sha1"
	"crypto/hmac"
	"fmt"
	"encoding/base64"
	"time"
	"strings"
	"strconv"
)

const (
	SINGNATURE_V1  = 1
	SINGNATURE_V2  = 2
	DEFAULT_SIGNATURE_VER = SINGNATURE_V1
	SECRET_KEY = "jkfldakjdksjafds&(*%^&GHJVE$%CUVuifgdiaisdf(&*&676"
)

func GetSecureCookie(
	c *gin.Context,
	name string,
	maxAgeDays int) (string,error){

	rawCookie, err:= c.Cookie(name)
	var value = ""
	if err != nil{
		return value, err
	}

	value, err = decodeSignedValue(name,rawCookie,maxAgeDays)
	return value, err
}

func decodeSignedValue(
	name string,
	value string,
	maxAgeDays int) (string,error){

	version := DEFAULT_SIGNATURE_VER

	var v string
	var err error
	if version == SINGNATURE_V1{
		v,err = decodeSignedValueV1(
			SECRET_KEY,name,value,maxAgeDays)
	}else if version == SINGNATURE_V2 {
		//V2没实现
		v,err = value,nil
	}
	return v,err
}

func decodeSignedValueV1(
	secret string,
	name string,
	value string,
	maxAgeDays int) (string, error){

	clock := time.Now().Unix()
	parts := strings.Split(value,"|")

	if len(parts) != 3{
		return "", errors.New("secure cookie format error")
	}

	signature := createSignatureV1(SECRET_KEY,name,parts[0],parts[1])
	if parts[2] != signature{
		return "", errors.New("secure cookie singnature error")
	}

	int64Timestamp,_ := strconv.ParseInt(parts[1], 10, 64)
	if int64Timestamp < int64(int64(clock) - int64(maxAgeDays * 86400)){
		return "", errors.New("secure cookie expired")
	}

	v, _ := base64.StdEncoding.DecodeString(parts[0])
	s := string(v)
	return s, nil
}


func  SetSecureCookie(
	c *gin.Context,
	name string,
	value string,
	maxAge int,
	path string,
	domain string,
	secure bool,
	httpOnly bool, ) {

	signedValue := createSignedValue(SECRET_KEY, name, value)
	if path == "" {
		path = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    signedValue,
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
}

func createSignedValue(
	secret string,
	name string,
	value string, ) string{

	timestamp := fmt.Sprintf(
		"%d", time.Now().Unix())

	version := DEFAULT_SIGNATURE_VER
	b64Value := base64.StdEncoding.EncodeToString(
		[]byte(value))

	var signedValue string
	if version == SINGNATURE_V1{

		signature:= createSignatureV1(
			secret,name,b64Value,timestamp)
		signedValue = b64Value+ "|" + timestamp + "|" + signature

	}else if version == SINGNATURE_V2 {
		//V2没实现
		signedValue = value
	}
	return signedValue
}


func createSignatureV1(
	secret string,
	name string,
	value string,
	timestamp string) string{

	key := secret
	h := hmac.New(sha1.New, []byte(key))
	h.Write([] byte(name + value + timestamp))
	sign := hex.EncodeToString(h.Sum(nil))
	return sign
}

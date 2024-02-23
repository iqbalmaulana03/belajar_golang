package encodedicode

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"
)

func EncodeToString() {
	data := "john wick"

	encodeString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encode :", encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodeString = string(decodeByte)
	fmt.Println("decode :", decodeString)
}

func Encode() {
	data := "john wick"

	var encode = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encode, []byte(data))
	var encodedString = string(encode)
	fmt.Println("encodeString :", encodedString)

	var decode = make([]byte, base64.StdEncoding.DecodedLen(len(encode)))
	var _, err = base64.StdEncoding.Decode(decode, encode)
	if err != nil {
		fmt.Println(err.Error())
	}

	var decodedString = string(decode)
	fmt.Println("decode :", decodedString)
}

func EncodeDataURL() {
	data := "https://kalipare.com/"

	encodeString := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("encode UR: :", encodeString)

	var decodeByte, _ = base64.URLEncoding.DecodeString(encodeString)
	decodeString := string(decodeByte)
	fmt.Println("decode URL :", decodeString)
}

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func SaltingSHA() {
	var text = "this is secret"
	fmt.Printf("original : %s\n\n", text)

	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)
	// 929fd8b1e58afca1ebbe30beac3b84e63882ee1a

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)
	// cda603d95286f0aece4b3e1749abe7128a4eed78

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)
	// 9e2b514bca911cb76f7630da50a99d4f4bb200b4

	_, _, _ = salt1, salt2, salt3
}

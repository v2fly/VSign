package sign

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"github.com/xiaokangwang/VSign/sign/signify"
	"golang.org/x/crypto/sha3"
)

func GenerateKeyFromSeed(seed string, password string) ([]byte, []byte) {
	shaw := sha3.NewShake256()
	shaw.Write([]byte(seed))
	pub, prv, err := signify.GenerateKey(shaw)
	if err != nil {
		panic(err)
	}

	pubb := signify.MarshalPublicKey(pub)
	prvb, err := signify.MarshalPrivateKey(prv, rand.Reader, []byte(password), 42)
	if err != nil {
		panic(err)
	}

	return prvb, pubb

}

func Sign(key []byte, password string, msg []byte) ([]byte, error) {
	pvkey, err := signify.ParsePrivateKey(key, []byte(password))
	if err != nil {
		return nil, err
	}
	out := bytes.NewBuffer(nil)
	outb := base64.NewEncoder(base64.StdEncoding, out)
	outb.Write(signify.MarshalSignature(signify.Sign(pvkey, msg)))
	outb.Close()
	return out.Bytes(), nil
}

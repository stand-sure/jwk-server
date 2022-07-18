package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/lestrrat-go/jwx/jwk"
)

// makeJwk makes a JSON Web Key
func makeJWK() (jwk.Key, error) {
	keyText := getKey()
	keyBytes, err := base64.StdEncoding.DecodeString(keyText)

	if err != nil {
		return nil, err
	}

	myKey, err := jwk.New(keyBytes)

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write(keyBytes)
	kid := hex.EncodeToString(hash.Sum(nil))

	myKey.Set(jwk.KeyIDKey, kid)
	myKey.Set(jwk.KeyUsageKey, jwk.ForSignature)

	return myKey, err
}

func makeSet() (jwk.Set, error) {
	set := jwk.NewSet()
	key, err := makeJWK()

	if err != nil {
		return nil, err
	}

	set.Add(key)

	return set, nil
}

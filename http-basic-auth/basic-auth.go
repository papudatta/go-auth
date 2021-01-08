package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	//"net/http"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// go list -m -versions github.com/dgrijalva/jwt-go

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("Invalid session ID")
	}
	return nil
}

var key = []byte{}

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
	fmt.Println()
	pass := "12340987abcd"
	hashedPass, err := hashP(pass)
	if err != nil {
		panic(err)
	}

	compareP(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in!")
	}
	log.Println("Logged in!!")
}

func hashP(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error generating hash: %w", err)
	}
	return bs, nil
}

func compareP(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("You forgot your password: %w", err)
	}
	return nil
}

func signMsg(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMsg while hashing: %w", err)
	}
	sig := h.Sum(nil)
	return sig, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMsg(msg)
	if err != nil {
		return false, fmt.Errorf("Error in signMsg while hashing: %w", err)
	}
	same := hmac.Equal(newSig, sig)
	return same, nil
}

func craeteToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}
	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm!")
		}

		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error in parseToken while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("Error in parseToken, token is invalid")
	}

	return t.Claims.(*UserClaims), nil
}

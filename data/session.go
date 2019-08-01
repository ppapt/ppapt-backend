// The data package defines data structures and functions to handle them
//
// This file session.go contains the data structures to handle sessions for
//logged in users.
package data

import (
	"math/rand"
	"time"
	"unsafe"
)

// UserSession contains the token and its expire time.
type UserSession struct {
	Token  string
	Expire time.Time
}

// Session stores UserSession data per user (EMail)
type Session map[string]UserSession

// TokenResponse is the data structure returned to the frontend, when requesting
//a token or logging in.
type TokenResponse struct {
	Token string `json:"token"`
}

// session_length defines, how long a token is valid
const session_length = 3600

// characters used for token generation
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."

// 6 bits to represent a letter index
const letterIdxBits = 6

// All 1-bits, as many as letterIdxBits
const letterIdxMask = 1<<letterIdxBits - 1

// # of letter indices fitting in 63 bits
const letterIdxMax = 63 / letterIdxBits

// length of the token string
const token_len = 32

// Random token string generation, see
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func newToken() string {

	rand.Seed(time.Now().UnixNano())
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, token_len)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := token_len-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// Add (or replace) a session to a sessions map.
// EMail is the users EMail address (primary key)
func (s Session) AddSession(EMail string) string {
	t := newToken()
	e := time.Now().Add(time.Minute * time.Duration(session_length))
	u := UserSession{
		Token:  t,
		Expire: e,
	}
	s[EMail] = u
	return t
}
package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtEncoding interface {
	MustEncode(c *Clamis) string
	Decode(s string) (*Clamis, bool)
	Expire() int
}

type jwtClamis struct {
	jwt.RegisteredClaims
	Payload *Clamis `json:"payload"`
}

type jwtEncoding struct {
	secret []byte
	expire int
}

func NewJwt(secret []byte, expire int) JwtEncoding {
	return &jwtEncoding{
		secret: secret,
		expire: expire,
	}
}

func (j *jwtEncoding) MustEncode(c *Clamis) string {
	ins := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClamis{
		Payload: c,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"Coorify Frontend"},
			Issuer:    "Coorify Backend",
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Expire() * int(time.Second)))),
		},
	})

	tk, err := ins.SignedString(j.secret)
	if err != nil {
		panic(err)
	}

	return tk
}

func (j *jwtEncoding) Decode(s string) (*Clamis, bool) {
	ins, err := jwt.ParseWithClaims(s, &jwtClamis{}, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil || ins == nil {
		return nil, false
	}

	v, ok := ins.Claims.(*jwtClamis)
	if ok {
		return v.Payload, true
	}

	return nil, false
}

func (j *jwtEncoding) Expire() int {
	return j.expire
}

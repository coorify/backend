package plugin

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/option"
	"github.com/emmansun/gmsm/sm2"
	"github.com/gin-gonic/gin"
)

func Signature(opt *option.Option) gin.HandlerFunc {
	if opt.Signature == nil {
		return func(ctx *gin.Context) {}
	}

	var err error
	var priKey *sm2.PrivateKey = nil
	var pubkey *ecdsa.PublicKey = nil

	priBytes, _ := hex.DecodeString(opt.Signature.Pri)
	priKey, err = sm2.NewPrivateKey(priBytes)
	if err != nil {
		panic(err)
	}

	pubBytes, _ := hex.DecodeString(opt.Signature.Pub)
	pubkey, err = sm2.NewPublicKey(pubBytes)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set(field.SYS_SIGPRIKEY, priKey)
		c.Set(field.SYS_SIGPUBKEY, pubkey)
	}
}

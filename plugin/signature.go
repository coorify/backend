package plugin

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/option"
	"github.com/coorify/go-value"
	"github.com/emmansun/gmsm/sm2"
	"github.com/gin-gonic/gin"
)

func Signature(opt interface{}) gin.HandlerFunc {
	o := value.GetWithDefault(opt, "Signature", nil).(*option.SignatureOption)

	if o == nil {
		return func(ctx *gin.Context) {}
	}

	var err error
	var priKey *sm2.PrivateKey = nil
	var pubkey *ecdsa.PublicKey = nil

	priBytes, _ := hex.DecodeString(o.Pri)
	priKey, err = sm2.NewPrivateKey(priBytes)
	if err != nil {
		panic(err)
	}

	pubBytes, _ := hex.DecodeString(o.Pub)
	pubkey, err = sm2.NewPublicKey(pubBytes)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set(field.SYS_SIGPRIKEY, priKey)
		c.Set(field.SYS_SIGPUBKEY, pubkey)
	}
}

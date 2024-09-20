package plugin

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/coorify/backend/field"
	"github.com/coorify/go-value"
	"github.com/emmansun/gmsm/sm2"
	"github.com/gin-gonic/gin"
)

func Signature(s Server) gin.HandlerFunc {
	o := s.Option()

	enable := value.MustGet(o, "Signature.Enable").(bool)
	pri := value.MustGet(o, "Signature.Pri").(string)
	pub := value.MustGet(o, "Signature.Pub").(string)
	if !enable {
		return func(ctx *gin.Context) {}
	}

	var err error
	var priKey *sm2.PrivateKey = nil
	var pubkey *ecdsa.PublicKey = nil

	priBytes, _ := hex.DecodeString(pri)
	priKey, err = sm2.NewPrivateKey(priBytes)
	if err != nil {
		panic(err)
	}

	pubBytes, _ := hex.DecodeString(pub)
	pubkey, err = sm2.NewPublicKey(pubBytes)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set(field.SYS_SIGPRIKEY, priKey)
		c.Set(field.SYS_SIGPUBKEY, pubkey)
	}
}

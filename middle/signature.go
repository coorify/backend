package middle

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/reply"
	"github.com/coorify/go-value"
	"github.com/emmansun/gmsm/sm2"
	"github.com/gin-gonic/gin"
)

type replySignatureWriter struct {
	gin.ResponseWriter
	pri *sm2.PrivateKey
}

func (w replySignatureWriter) Write(b []byte) (int, error) {
	rawBytes, _ := w.pri.SignWithSM2(rand.Reader, nil, b)
	signature := hex.EncodeToString(rawBytes)

	w.Header().Add("signature", signature)
	return w.ResponseWriter.Write(b)
}

func Signature(c *gin.Context) {
	if value.MustGet(c.MustGet(field.SYS_OPTION), "Signature") == nil {
		return
	}

	pri := c.MustGet(field.SYS_SIGPRIKEY).(*sm2.PrivateKey)
	pub := c.MustGet(field.SYS_SIGPUBKEY).(*ecdsa.PublicKey)

	raws := make([]byte, 0)
	csig := c.GetHeader("signature")

	c.Writer = &replySignatureWriter{
		ResponseWriter: c.Writer,
		pri:            pri,
	}

	rawBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Abort()
		reply.NewReply(http.StatusForbidden, nil, "Forbidden", c)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(rawBytes))
	raws = append(raws, rawBytes...)

	rawBytes = []byte(c.Request.URL.RawQuery)
	raws = append(raws, rawBytes...)

	bcsin, err := hex.DecodeString(csig)
	if err != nil {
		c.Abort()
		reply.NewReply(http.StatusForbidden, nil, "Forbidden", c)
		return
	}

	if !sm2.VerifyASN1WithSM2(pub, nil, raws, bcsin) {
		c.Abort()
		reply.NewReply(http.StatusForbidden, nil, "Forbidden", c)
	}
}

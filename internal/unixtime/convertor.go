package b64

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var errBadRequest = errors.New("bad request")

type handler struct{}

func (h handler) decode(c *gin.Context) {
	b, ok := c.GetQuery("base64")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, errBadRequest)
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.String(http.StatusOK, string(decoded))
}

func (h handler) encode(c *gin.Context) {
	text, ok := c.GetQuery("text")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, errBadRequest)
		return
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	c.String(http.StatusOK, encoded)
}


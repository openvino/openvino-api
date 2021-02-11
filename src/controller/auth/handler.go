package auth

import (
	"crypto"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/openvino/openvino-api/src/repository"

	"github.com/openvino/openvino-api/src/config"
	customHTTP "github.com/openvino/openvino-api/src/http"
)

type authData struct {
	Expire    int64  `json:"expire"`
	Address   string `json:"address"`
	Role      string `json:"role"`
	Integrity string `json:"integrity"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {

	var authData authData
	authData.Address = r.URL.Query().Get("public_key")
	authData.Expire = time.Now().UnixNano()/int64(time.Millisecond) + 900000
	authData.Role = "Guest"
	authData.Address = authData.Address[2:len(authData.Address)]

	var domain, err = repository.GetDomain(authData.Address)
	if err == nil && strings.HasSuffix(domain, "rinkibino.eth") {
		authData.Role = "Worker"
	}

	total := authData.Address + "$" +
		strconv.FormatInt(authData.Expire, 10) + "$" +
		authData.Role + "$" +
		config.Config.Secret
	hasher := crypto.SHA512.New()
	hasher.Write([]byte(total))
	authData.Integrity = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	customHTTP.ResponseJSON(w, authData)

}

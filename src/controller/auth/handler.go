package auth

import (
	"crypto"
	"encoding/base64"
	"github.com/openvino/openvino-api/src/repository"
	"net/http"
	"strconv"
	"strings"
	"time"

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

	var domain, err = repository.GetDomain(authData.Address)
	if err != nil {
		customHTTP.NewErrorResponse(w, 400, "Wrong address")
	} else if strings.HasSuffix(domain, "rinkibino.eth") {
		authData.Role = "Worker"
	}

	authData.Address = authData.Address[2:len(authData.Address)]
	total := authData.Address + "$" +
		strconv.FormatInt(authData.Expire, 10) + "$" +
		authData.Role + "$" +
		config.Config.Secret
	hasher := crypto.SHA512.New()
	hasher.Write([]byte(total))
	authData.Integrity = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	customHTTP.ResponseJSON(w, authData)

}

package middleware

import (
	"bytes"
	basecrypto "crypto"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/openvino/openvino-api/src/config"
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/util"
)

const ApiKeyHeader = "X-API-Key"
const AuthHeader = "Authorization"

// AuthMiddleware - Checks if valid API Key provided
func AuthMiddleware(next http.Handler, scopes []customHTTP.Scope) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !util.ContainsScope(customHTTP.GuestScope.String(), scopes) {
			data, errS := getSignatureVerification(w, r)
			if len(errS) != 0 {
				customHTTP.NewErrorResponse(w, http.StatusUnauthorized, errS)
				return
			}

			var split = strings.Split(data, "$")
			data = split[0] + "$" + split[1] + "$" + split[3] + "$" + config.Config.Secret
			hasher := basecrypto.SHA512.New()
			hasher.Write([]byte(data))

			if split[2] != base64.URLEncoding.EncodeToString(hasher.Sum(nil)) {
				customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Invalid integrity proof")
				return
			}

			timestamp, err := strconv.Atoi(split[1])
			if err != nil || int64(timestamp) <= time.Now().UnixNano()/int64(time.Millisecond) {
				customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "The key has expired")
				return
			}

			if !util.ContainsScope(split[3], scopes) {
				customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Unauthorized scope")
				return
			}

		}
		next.ServeHTTP(w, r)
	})
}

func getSignatureVerification(w http.ResponseWriter, r *http.Request) (string, string) {

	signedData := r.Header.Get(ApiKeyHeader)
	addressHex := strings.Split(signedData, "$")[0]
	addressBytes, err := hex.DecodeString(addressHex)
	if err != nil || len(addressHex) != 40 {
		return "", "No valid timestamp header present"
	}

	signatureHex := r.Header.Get(AuthHeader)
	signatureBytes, err := hex.DecodeString(signatureHex)

	if len(signatureHex) == 0 || err != nil {
		return "", "No " + AuthHeader + " header present"
	}

	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(signedData[:]), signedData[:])
	hash := crypto.Keccak256Hash([]byte(msg))

	sigPublicKeyBytes, err := crypto.Ecrecover(hash.Bytes(), signatureBytes)
	if err != nil {
		return "", err.Error() + "3"
	}

	sigPublicKeyECDSA, err := crypto.UnmarshalPubkey(sigPublicKeyBytes)
	if err != nil {
		return "", err.Error() + "4"
	}

	sigAddress := crypto.PubkeyToAddress(*sigPublicKeyECDSA)

	privateKeyOk := bytes.Equal(sigAddress.Bytes(), addressBytes)

	if !privateKeyOk {
		return "", "Unauthorized signature"
	}

	return signedData, ""

}

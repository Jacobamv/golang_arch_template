package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/Jacobamv/golang_arch_template/pkg/bootstrap/http/misc/response/customRW"
	"github.com/rs/zerolog"
)

func (m *provider) RequestLog(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dateStart := time.Now()
		ctx := r.Context()
		lrw := customRW.NewResponseWriter(w)
		bodyBytes, _ := io.ReadAll(r.Body)

		r.Body.Close() //  must close
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		next(lrw, r)

		dateEnd := time.Now()

		duration := dateEnd.Unix() - dateStart.Unix()

		requestHeaders := make(map[string]string)

		for key, val := range r.Header {
			if key == "Authorization" {
				continue
			}
			requestHeaders[key] = val[0]
		}

		requestParams := make(map[string][]string)

		for key, val := range r.URL.Query() {
			requestParams[key] = val
		}

		var requestBody map[string]interface{}

		bodyString := string(bodyBytes)

		bodyString = string(regexp.MustCompile(`base64,.*"`).ReplaceAll([]byte(bodyString), []byte("base64,<basefile>")))

		_ = json.Unmarshal(bodyBytes, &requestBody)

		var responseBodyString string
		var responseBodyJSON map[string]interface{}

		responseBodyString = string(lrw.GetContent())

		responseBodyString = string(regexp.MustCompile(`base64,.*"`).ReplaceAll([]byte(responseBodyString), []byte("base64,<basefile>")))

		_ = json.Unmarshal(lrw.GetContent(), &responseBodyJSON)

		responseHeaders := make(map[string]string)

		for key, val := range lrw.Header() {
			responseHeaders[key] = val[0]
		}

		if len(bodyString) < 3 {
			bodyString = " "
		}
		if len(responseBodyString) < 3 {
			responseBodyString = " "
		}

		m.logger.Info().
			Ctx(ctx).
			Dict("request", zerolog.Dict().
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Str("remote_addr", r.RemoteAddr).
				Any("request_headers", requestHeaders).
				Any("request_params", requestParams).
				Str("user_agent", r.UserAgent()).
				Any("body", requestBody).
				RawJSON("body_raw", []byte(bodyString))).
			Dict("response", zerolog.Dict().
				Any("headers", responseHeaders).
				Int("status", lrw.Status()).
				Any("body", responseBodyJSON).
				RawJSON("body_raw", []byte(responseBodyString))).
			Int("duration", int(duration)).Msg("Request")

	})
}

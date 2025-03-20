package example

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Jacobamv/golang_arch_template/pkg/bootstrap/http/misc/response"
)

func (p *provider) Login(ctx context.Context) (err error) {
	url := p.config.Example.Url + "/login/token"
	method := "POST"

	payload, err := json.Marshal(nil)

	if err != nil {
		return
	}

	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		p.logger.Error().Ctx(ctx).Err(err).Msg("An error occurred while tring to create request in example Login")
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		p.logger.Error().Ctx(ctx).Err(err).Msg("An error occurred while tring to send request in example Login")
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		p.logger.Error().Ctx(ctx).Err(err).Int("response status", res.StatusCode).Msg("An error occurred while tring to send request in example Login")
		return response.ErrSomethingWentWrong
	}

	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		p.logger.Error().Ctx(ctx).Err(err).Msg("An error occurred while tring to read response in example Login")
		return response.ErrSomethingWentWrong
	}

	var responseLogin map[string]string

	err = json.Unmarshal(respBody, &responseLogin)

	if err != nil {
		p.logger.Error().Ctx(ctx).Err(err).Msg("An error occurred while tring to unmarshal response in example Login")
		return response.ErrSomethingWentWrong
	}

	p.AccessToken = responseLogin["access_token"]
	p.RefreshToken = responseLogin["refresh_token"]
	p.TokenType = responseLogin["token_type"]

	return
}

package example

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/misc/response"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/metrics"
)

func (p *provider) Ping(ctx context.Context) (err error) {
	url := p.config.Example.Url + "/health"
	method := "GET"

	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		p.logger.Error().Ctx(ctx).Err(err).Send()
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		p.logger.Error().Ctx(ctx).Err(err).Send()
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		metrics.SetMicroserviceStatus(0)
		p.logger.Error().Ctx(ctx).Err(err).Send()
		return response.ErrSomethingWentWrong
	}

	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		metrics.SetMicroserviceStatus(0)
		p.logger.Error().Ctx(ctx).Err(err).Msg("An error occurred while tring to read response in scoring Ping")
		return response.ErrSomethingWentWrong
	}

	var health map[string]string

	err = json.Unmarshal(respBody, &health)

	if err != nil {
		metrics.SetMicroserviceStatus(0)
		p.logger.Error().Ctx(ctx).Err(err).Msg("An error occurred while tring to unmarshal response in scoring Ping")
		return response.ErrSomethingWentWrong
	}

	if health["status"] == "ok" {
		metrics.SetMicroserviceStatus(1)
	} else {
		metrics.SetMicroserviceStatus(0)
	}

	return
}

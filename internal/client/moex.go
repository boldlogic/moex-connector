package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

var BaseURL string = "https://iss.moex.com/iss"

func (c *Client) GetSecurityInfo(ctx context.Context, ticker string) ([]byte, error) {

	path := "securities"
	rawUrl := fmt.Sprintf("%s/%s/%s", BaseURL, path, ticker)

	rawQuery := make(map[string]string)
	rawQuery["iss.only"] = "description"

	reqURL, err := url.Parse(rawUrl)
	if err != nil {
		c.logger.Error("некорректный URL", zap.Error(err))
		return nil, err
	}
	query := reqURL.Query()
	for key, value := range rawQuery {
		query.Set(key, value)
	}
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		c.logger.Error("ошибка создания запроса", zap.Error(err))
		return nil, err
	}

	code, resp, err := c.SendRequest(ctx, req)
	if err != nil {
		c.logger.Error("ошибка выполнения запроса", zap.Error(err))
		return nil, err
	}
	if code != http.StatusOK {
		c.logger.Error("ошибка выполнения запроса", zap.Int("code", code))
		return nil, err
	}
	return resp, nil

}

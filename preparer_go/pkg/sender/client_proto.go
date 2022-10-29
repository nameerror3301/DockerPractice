package sender

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ClientPrep struct {
	conn *http.Client
	url  string
}

func NewClient(addr string, port int) (*ClientPrep, error) {
	return &ClientPrep{
		conn: &http.Client{},
		url:  addr + ":" + strconv.Itoa(port),
	}, nil
}

func (c *ClientPrep) SendDataOnPrep(data []byte) error {
	if c.conn == nil {
		return fmt.Errorf("invalid init prep client")
	}

	bode := bytes.NewReader(data)

	req, err := http.NewRequest(http.MethodPost, c.url+"/store/data", bode)
	if err != nil {
		return err
	}

	res, err := c.conn.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		resData, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed request. Code %d", res.StatusCode)
		}
		return fmt.Errorf("failed request. Code %d. Body: %v", res.StatusCode, string(resData))
	}

	return nil
}

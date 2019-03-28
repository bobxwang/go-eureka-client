package eureka

import (
	"net/http"
	"strings"
	"time"
)

func (c *Client) SendHeartbeat(appId, instanceId string, second int) error {
	values := []string{"apps", appId, instanceId}
	path := strings.Join(values, "/")
	resp, err := c.Put(path, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusNotFound:
		return newError(ErrCodeInstanceNotFound,
			"Instance resource not found when sending heartbeat", 0)
	}

	go sendHeartbeat(c, path, second)

	return nil
}

func sendHeartbeat(c *Client, path string, second int) {

	d := (time.Duration(second) * time.Second)
	tick := time.Tick(d)
	for {
		select {
		case <-tick:
			c.Put(path, nil)
		}
	}
}

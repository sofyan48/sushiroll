package argo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sofyan48/sushiroll/src/pkg/requester"
	"github.com/sofyan48/sushiroll/src/presentations"
)

type argoRollout struct {
	request   requester.Contract
	Url       string
	Username  string
	Password  string
	Namespace string
}

func NewArgoRolloutLibrary(request requester.Contract) ArgoRolloutLibrary {
	return &argoRollout{
		request:   request,
		Url:       os.Getenv("ARGO_ROLLOUT_URL"),
		Username:  os.Getenv("ARGO_ROLLOUT_USERNAME"),
		Password:  os.Getenv("ARGO_ROLLOUT_PASSWORD"),
		Namespace: os.Getenv("ARGO_ROLLOUT_NAMESPACE"),
	}
}

func (a *argoRollout) client(url, method string, body io.Reader) (*http.Request, error) {
	reqs, err := a.request.RAW(method, url, body)
	if err != nil {
		return nil, err
	}
	reqs.SetBasicAuth(a.Username, a.Password)
	return reqs, nil
}

func (a *argoRollout) GetList() (*presentations.RolloutList, error) {
	endpoint := a.Url + "/api/v1/rollouts/" + a.Namespace + "/info"
	req, err := a.client(endpoint, "GET", nil)
	if err != nil {
		return nil, err
	}
	result := &presentations.RolloutList{}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *argoRollout) Detail(service string) (*presentations.RolloutDetail, error) {
	endpoint := a.Url + "/api/v1/rollouts/" + a.Namespace + "/" + service + "/info"
	req, err := a.client(endpoint, "GET", nil)
	if err != nil {
		return nil, err
	}
	result := &presentations.RolloutDetail{}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *argoRollout) Promote(isFull bool, service string) (*presentations.PromoteArgoResponse, error) {
	endpoint := a.Url + "/api/v1/rollouts/" + a.Namespace + "/" + service + "/promote"
	payload := map[string]interface{}{
		"full": isFull,
	}
	bodyData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := a.client(endpoint, "PUT", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	result := &presentations.PromoteArgoResponse{}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *argoRollout) Rollback(revision, service string) ([]byte, error) {
	endpoint := a.Url + "/api/v1/rollouts/" + a.Namespace + "/" + service + "/undo/" + revision
	payload := map[string]interface{}{}
	bodyData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := a.client(endpoint, "PUT", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *argoRollout) Restart(service string) (*presentations.PromoteArgoResponse, error) {
	endpoint := a.Url + "/api/v1/rollouts/" + a.Namespace + "/" + service + "/restart"
	payload := map[string]interface{}{}
	bodyData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := a.client(endpoint, "PUT", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	result := &presentations.PromoteArgoResponse{}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *argoRollout) Retry(service string) (*presentations.PromoteArgoResponse, error) {
	endpoint := a.Url + "/api/v1/rollouts/" + a.Namespace + "/" + service + "/retry"
	payload := map[string]interface{}{}
	bodyData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := a.client(endpoint, "PUT", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	result := &presentations.PromoteArgoResponse{}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *argoRollout) Abort(service string) (*presentations.PromoteArgoResponse, error) {
	endpoint := a.Url + "/api/v1/rollouts/" + a.Namespace + "/" + service + "/abort"
	payload := map[string]interface{}{}
	bodyData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := a.client(endpoint, "PUT", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	result := &presentations.PromoteArgoResponse{}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

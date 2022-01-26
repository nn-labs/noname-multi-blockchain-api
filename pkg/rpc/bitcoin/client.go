package bitcoin

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"nn-blockchain-api/pkg/errors"
)

type btcClient struct {
	btcEndpoint string
	btcUser     string
	btcPassword string
}

type BaseRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,array"`
}

type BaseResponse struct {
	Id      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

type IBtcClient interface {
	Send(body io.Reader) (*http.Response, error)
	EncodeBaseRequest(request BaseRequest) (*bytes.Buffer, error)
	//DecodeBaseResponse(response *http.Response, msg interface{}) (*BaseResponse, error)
}

func NewBtcClient(btcEndpoint, btcUser, btcPassword string) (IBtcClient, error) {
	if btcEndpoint == "" {
		return nil, errors.NewInternal("failed check btc endpoint")
	}
	if btcUser == "" {
		return nil, errors.NewInternal("failed check btc user")
	}
	if btcPassword == "" {
		return nil, errors.NewInternal("failed check btc password")
	}

	return &btcClient{
		btcEndpoint: btcEndpoint,
		btcUser:     btcUser,
		btcPassword: btcPassword,
	}, nil
}

func (btc *btcClient) Send(body io.Reader) (*http.Response, error) {
	//resp, err := http.Post(e.btcEndpoint, "application/json", body)
	//return resp, err

	client := &http.Client{}
	req, err := http.NewRequest("POST", btc.btcEndpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(btc.btcUser, btc.btcPassword)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//defer resp.Body.Close()

	return resp, nil
}

func (btc *btcClient) EncodeBaseRequest(request BaseRequest) (*bytes.Buffer, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(data)

	return reqBody, nil
}

//func (btc *btcClient) DecodeBaseResponse(response *http.Response, msg interface{}) (*BaseResponse, error) {
//	err := json.NewDecoder(response.Body).Decode(&response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &msg, nil
//}

//
//func Client(body, res interface{}) error {
//	//remoteURL := "http://159.89.6.17:8332"
//	localURL := "http://127.0.0.1:8332"
//
//	//var serverAddr string
//	//
//	//if walletInfo {
//	//	serverAddr = remoteURL + "/wallet/" + walletId // testnet/main net
//	//} else {
//	//	serverAddr = remoteURL
//	//}
//
//	client := &http.Client{}
//
//	jsonBody, _ := json.Marshal(body)
//	reqBody := bytes.NewBuffer(jsonBody)
//	req, err := http.NewRequest("POST", localURL, reqBody)
//	if err != nil {
//		return err
//	}
//
//	req.Header.Add("Content-Type", "application/json")
//	req.Header.Add("Accept", "application/json")
//	req.SetBasicAuth("uuuset", "password123123")
//
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//
//	defer resp.Body.Close()
//
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return err
//	}
//
//	//fmt.Println(string(respBody))
//
//	err = json.Unmarshal(respBody, res)
//	if err != nil {
//		return err
//	}
//
//	if resp.StatusCode != 200 {
//		return errors.New(string(respBody))
//	}
//
//	return nil
//}
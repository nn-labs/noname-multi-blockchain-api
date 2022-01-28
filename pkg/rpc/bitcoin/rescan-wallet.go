package bitcoin

import (
	"errors"
	"time"
)

func RescanWallet(client IBtcClient, walletId, network string) error {
	//msg := struct {
	//	Result struct {
	//		StartHeight int64 `json:"start_height"`
	//		StopHeight  int64 `json:"stop_height"`
	//	} `json:"result"`
	//	Error struct {
	//		Message string `json:"message"`
	//	} `json:"error"`
	//}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "rescanblockchain",
		Params:  []interface{}{},
	}

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return errors.New(err.Error())
	}

	errs := make(chan error, 1)
	go func() {
		response, err := client.Send(body, walletId, network)
		if err != nil {
			errs <- errors.New(err.Error())
		}

		//err = json.NewDecoder(response.Body).Decode(&msg)
		//if err != nil {
		//	errs <- errors.New(err.Error())
		//}
		//
		//if msg.Error.Message != "" {
		//	errs <- errors.New(msg.Error.Message)
		//}

		defer response.Body.Close()
		close(errs)
	}()
	time.Sleep(2 * time.Second)

	if err := <-errs; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

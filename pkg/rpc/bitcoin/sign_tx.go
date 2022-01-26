package bitcoin

//func SignTx(client IBtcClient) (string, error) {
//	req := BaseRequest{
//		JsonRpc: "2.0",
//		Method:  "getblockchaininfo",
//		Params:  []interface{}{},
//	}
//
//	msg := struct {
//		Result StatusNode `json:"result"`
//		Error  struct {
//			Code    int64  `json:"code"`
//			Message string `json:"message"`
//		} `json:"error"`
//	}{}
//
//	body, err := client.EncodeBaseRequest(req)
//	if err != nil {
//		return nil, errors.New(err.Error())
//	}
//
//	response, err := client.Send(body)
//	if err != nil {
//		return nil, errors.New(err.Error())
//	}
//
//	err = json.NewDecoder(response.Body).Decode(&msg)
//	if err != nil {
//		return nil, err
//	}
//
//	return &msg.Result, nil
//}

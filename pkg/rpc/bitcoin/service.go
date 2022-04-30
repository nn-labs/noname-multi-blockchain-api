package bitcoin_rpc

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/google/uuid"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go
type Service interface {
	Status(ctx context.Context, network string) (*StatusNode, error)

	GetCurrentFee(ctx context.Context, network string) (*float64, error)

	CreateTransaction(ctx context.Context, utxos UTXO, fromAddress, toAddress string, amount int64, network string) (*string, *float64, error)
	//CreateTransaction(ctx context.Context,inputs []map[string]interface{}, outputs []map[string]string, network string) (string, error)
	DecodeTransaction(ctx context.Context, tx string, network string) (*DecodedTx, error)
	FundForTransaction(ctx context.Context, createdTx, changeAddress, network string) (string, *float64, error)
	SignTransaction(ctx context.Context, tx, privateKey string, utxos UTXO, network string) (string, error)
	SendTransaction(ctx context.Context, signedTx, network string) (string, error)

	WalletInfo(ctx context.Context, walletId, network string) (*Info, error)
	CreateWallet(ctx context.Context, network string) (string, error)
	LoadWallet(ctx context.Context, walletId, network string) error
	ImportAddress(ctx context.Context, address, walletId, network string) error
	RescanWallet(ctx context.Context, walletId, network string) error
	ListUnspent(ctx context.Context, address, walletId, network string) ([]*Unspent, error)
}

type service struct {
	btcClient Client
}

func NewService(btcClient Client) (Service, error) {
	if btcClient == nil {
		return nil, errors.New("invalid bitcoin rpc client")
	}
	return &service{btcClient: btcClient}, nil
}

func (s *service) Status(ctx context.Context, network string) (*StatusNode, error) {
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "getblockchaininfo",
		Params:  []interface{}{},
	}

	msg := struct {
		Result StatusNode `json:"result"`
		Error  struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := s.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}

func (s *service) GetCurrentFee(ctx context.Context, network string) (*float64, error) {
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "estimatesmartfee",
		Params:  []interface{}{2},
	}

	msg := struct {
		Result struct {
			Feerate float64 `json:"feerate"`
			Blocks  int64   `json:"blocks"`
		}
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := s.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	var fee float64
	fee = msg.Result.Feerate
	// sanity check
	if fee > 0.05 {
		fee = 0.1
	} else if fee < 0 {
		fee = 0
	}
	//fmt.Printf("fee: %f\n", fee)

	if fee == 0 {
		return &fee, err
	}

	return &fee, nil
}

func (s *service) getCurrentFeeRate(ctx context.Context, network string) (*big.Int, error) {
	fee, err := s.GetCurrentFee(ctx, network)
	if err != nil {
		return nil, err
	}

	// convert to satoshis to bytes
	// feeRate := big.NewInt(int64(msg.Result * 1.0E8))
	// convert to satoshis to kb
	feeRate := big.NewInt(int64(*fee * 1.0e5))

	//fmt.Printf("fee rate: %s\n", feeRate)

	return feeRate, nil
}

func (s *service) CreateTransaction(ctx context.Context, utxos UTXO, fromAddress, toAddress string, amount int64, network string) (*string, *float64, error) {
	chainParams := &chaincfg.TestNet3Params

	// Get fee
	feeRate, err := s.getCurrentFeeRate(ctx, network)
	if err != nil {
		return nil, nil, err
	}

	// Calculate all unspent amount
	utxosAmount := big.NewInt(0)
	for idx := range utxos {
		utxosAmount.Add(utxosAmount, new(big.Int).SetInt64(utxos[idx].Amount))
	}

	// Init transaction
	tx := wire.NewMsgTx(2)

	// prepare transaction inputs
	sourceUtxosAmount := big.NewInt(0)
	// var sourceUTXOs []*UnspentList

	for idx := range utxos {
		hashStr := utxos[idx].TxId
		sourceUtxosAmount.Add(sourceUtxosAmount, new(big.Int).SetInt64(utxos[idx].Amount))

		sourceUTXOHash, err := chainhash.NewHashFromStr(hashStr)
		if err != nil {
			return nil, nil, err
		}

		if amount <= sourceUtxosAmount.Int64() {
			sourceUTXOIndex := uint32(utxos[idx].Vout)
			sourceUTXO := wire.NewOutPoint(sourceUTXOHash, sourceUTXOIndex)
			// sourceUTXOs = append(sourceUTXOs, &UnspentList{
			// 	TxId:         utxos[idx].TxId,
			// 	Vout:         utxos[idx].Vout,
			// 	ScriptPubKey: utxos[idx].PKScript,
			// })
			sourceTxIn := wire.NewTxIn(sourceUTXO, nil, nil)

			tx.AddTxIn(sourceTxIn)
			break
		}

		sourceUTXOIndex := uint32(utxos[idx].Vout)
		sourceUTXO := wire.NewOutPoint(sourceUTXOHash, sourceUTXOIndex)
		// sourceUTXOs = append(sourceUTXOs, &UnspentList{
		// 	TxId:         utxos[idx].TxId,
		// 	Vout:         utxos[idx].Vout,
		// 	ScriptPubKey: utxos[idx].PKScript,
		// })
		sourceTxIn := wire.NewTxIn(sourceUTXO, nil, nil)

		tx.AddTxIn(sourceTxIn)
	}

	// create the transaction outputs
	destAddress, err := btcutil.DecodeAddress(toAddress, chainParams)
	if err != nil {
		return nil, nil, err
	}

	destScript, err := txscript.PayToAddrScript(destAddress)
	if err != nil {
		return nil, nil, err
	}

	// tx out to send btc to user
	destOutput := wire.NewTxOut(amount, destScript)
	tx.AddTxOut(destOutput)

	change := new(big.Int).Set(sourceUtxosAmount)
	change = new(big.Int).Sub(change, new(big.Int).SetInt64(amount))
	//change = new(big.Int).Sub(change, totalFee)
	if change.Cmp(big.NewInt(0)) == -1 {
		//svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, nil, err
	}

	if change.Int64() != 0 {
		// our fee address
		//feeSendToAddress, err := btcutil.DecodeAddress(fromWalletPublicAddress, chainParams)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//feeSendToScript, err := txscript.PayToAddrScript(feeSendToAddress)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		////tx out to send change back to us
		//feeOutput := wire.NewTxOut(changeFee.Int64(), feeSendToScript)
		//tx.AddTxOut(feeOutput)

		// our change address
		changeSendToAddress, err := btcutil.DecodeAddress(fromAddress, chainParams)
		if err != nil {
			//svc.log.WithContext(ctx).Errorf(err.Error())
			return nil, nil, err
		}

		changeSendToScript, err := txscript.PayToAddrScript(changeSendToAddress)
		if err != nil {
			//svc.log.WithContext(ctx).Errorf(err.Error())
			return nil, nil, err
		}

		//tx out to send change back to us
		changeOutput := wire.NewTxOut(change.Int64(), changeSendToScript)
		tx.AddTxOut(changeOutput)
	}

	// calculate fees
	txByteSize := big.NewInt(int64(len(tx.TxIn)*180 + len(tx.TxOut)*34 + 10 + len(tx.TxIn)))
	totalFee := new(big.Int).Mul(feeRate, txByteSize)

	// Need add fee to spend amount and then compare
	if (amount - totalFee.Int64()) >= sourceUtxosAmount.Int64() {
		//log.Fatal(errors.New("your balance too low for this transaction"))
		//svc.log.WithContext(ctx).Errorf("your balance too low for this transaction")
		return nil, nil, errors.New("your balance too low for this transaction")
	}

	//log.Printf("%-18s %s\n", "total fee:", totalFee)

	// Change amount of source output transaction
	tx.TxOut[0].Value = amount - totalFee.Int64()

	// Transaction Hash
	notSignedTxBuf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	err = tx.Serialize(notSignedTxBuf)
	if err != nil {
		//log.Fatal(err)
		//svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, nil, err
	}

	createdTx := hex.EncodeToString(notSignedTxBuf.Bytes())
	btcAmountFee := btcutil.Amount(totalFee.Int64()).ToBTC()
	return &createdTx, &btcAmountFee, nil
}

//func (s *service) CreateTransaction(ctx context.Context,client IBtcClient, inputs []map[string]interface{}, outputs []map[string]string, network string) (string, error) {
//	req := BaseRequest{
//		JsonRpc: "2.0",
//		Method:  "createrawtransaction",
//		Params:  []interface{}{inputs, outputs},
//	}
//
//	msg := struct {
//		Result string `json:"result"`
//		Error  struct {
//			Message string `json:"message"`
//		} `json:"error"`
//	}{}
//
//	body, err := client.EncodeBaseRequest(req)
//	if err != nil {
//		return "", err
//	}
//
//	response, err := client.Send(ctx, body, "", network)
//	if err != nil {
//		return "", err
//	}
//
//	defer response.Body.Close()
//
//	err = json.NewDecoder(response.Body).Decode(&msg)
//	if err != nil {
//		return "", err
//	}
//
//	if msg.Error.Message != "" {
//		return "", errors.New(msg.Error.Message)
//	}
//
//	return msg.Result, nil
//}

func (s *service) DecodeTransaction(ctx context.Context, tx string, network string) (*DecodedTx, error) {
	msg := struct {
		Result DecodedTx `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "decoderawtransaction",
		Params:  []interface{}{tx},
	}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := s.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &DecodedTx{
		Txid:     msg.Result.Txid,
		Hash:     msg.Result.Hash,
		Version:  msg.Result.Version,
		Size:     msg.Result.Size,
		Vsize:    msg.Result.Vsize,
		Weight:   msg.Result.Weight,
		Locktime: msg.Result.Locktime,
		Vin:      msg.Result.Vin,
		Vout:     msg.Result.Vout,
	}, nil
}

func (s *service) FundForTransaction(ctx context.Context, createdTx, changeAddress, network string) (string, *float64, error) {
	subtractFeeFromOutputs := []int64{0}

	params := map[string]interface{}{
		"changeAddress":          changeAddress,
		"subtractFeeFromOutputs": subtractFeeFromOutputs,
	}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "fundrawtransaction",
		Params:  []interface{}{createdTx, params},
	}

	msg := struct {
		Result struct {
			Hex string  `json:"hex"`
			Fee float64 `json:"fee"`
		} `json:"result"`
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return "", nil, err
	}

	response, err := s.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return "", nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", nil, err
	}

	if msg.Error.Message != "" {
		return "", nil, errors.New(msg.Error.Message)
	}

	return msg.Result.Hex, &msg.Result.Fee, nil
}

func (s *service) SignTransaction(ctx context.Context, tx, privateKey string, utxos UTXO, network string) (string, error) {
	privateKeyArray := []string{privateKey}
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "signrawtransactionwithkey",
		Params:  []interface{}{tx, privateKeyArray, utxos},
	}

	msg := struct {
		Result struct {
			Hex      string `json:"hex"`
			Complete bool   `json:"complete"`
		} `json:"result"`
		Error struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return "", err
	}

	response, err := s.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", err
	}

	if msg.Error.Message != "" {
		return "", errors.New(msg.Error.Message)
	}

	if !msg.Result.Complete {
		return "", errors.New("signing transaction not complete. Please try again")
	}

	return msg.Result.Hex, nil
}

func (s *service) SendTransaction(ctx context.Context, signedTx, network string) (string, error) {
	msg := struct {
		Result string `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "sendrawtransaction",
		Params:  []interface{}{signedTx},
	}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return "", err
	}

	response, err := s.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", err
	}

	if msg.Error.Message != "" {
		return "", errors.New(msg.Error.Message)
	}

	return msg.Result, nil
}

func (s *service) WalletInfo(ctx context.Context, walletId, network string) (*Info, error) {
	msg := struct {
		Result Info `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "getwalletinfo",
		Params:  []interface{}{},
	}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := s.btcClient.Send(ctx, body, walletId, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}

func (s *service) CreateWallet(ctx context.Context, network string) (string, error) {
	walletId, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "createwallet",
		Params:  []interface{}{walletId},
	}

	msg := struct {
		Result struct {
			Name    string `json:"name"`
			Warning string `json:"warning"`
		} `json:"result"`
		Error struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return "", err
	}

	response, err := s.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", err
	}

	if msg.Result.Warning != "" {
		return "", errors.New(msg.Result.Warning)
	}

	if msg.Error.Message != "" {
		return "", errors.New(msg.Error.Message)
	}

	return msg.Result.Name, nil
}

func (s *service) LoadWallet(ctx context.Context, walletId, network string) error {
	msg := struct {
		Result struct {
			Name    string `json:"name"`
			Warning string `json:"warning"`
		} `json:"result"`
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "loadwallet",
		Params:  []interface{}{walletId},
	}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return err
	}

	response, err := s.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return err
	}

	if msg.Result.Warning != "" {
		return errors.New(msg.Result.Warning)
	}

	if msg.Error.Message != "" {
		return errors.New(msg.Error.Message)
	}

	return nil
}

func (s *service) ImportAddress(ctx context.Context, address, walletId, network string) error {
	msg := struct {
		Result string `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "importaddress",
		Params:  []interface{}{address, "", false},
	}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return err
	}

	response, err := s.btcClient.Send(ctx, body, walletId, network)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return err
	}

	if msg.Error.Message != "" {
		return errors.New(msg.Error.Message)
	}

	return nil
}

func (s *service) RescanWallet(ctx context.Context, walletId, network string) error {
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

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return err
	}

	errs := make(chan error, 1)
	go func() {
		response, err := s.btcClient.Send(ctx, body, walletId, network)
		if err != nil {
			errs <- err
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

	select {
	case <-time.After(10 * time.Second):
		return nil
	case err := <-errs:
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) ListUnspent(ctx context.Context, address, walletId, network string) ([]*Unspent, error) {
	msg := struct {
		Result []*Unspent `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "listunspent",
		Params:  []interface{}{1, 99999999, []string{address}},
	}

	body, err := s.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := s.btcClient.Send(ctx, body, walletId, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return msg.Result, nil
}

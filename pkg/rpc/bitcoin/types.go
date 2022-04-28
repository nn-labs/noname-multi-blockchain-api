package rpc_bitcoin

type BaseRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type BaseResponse struct {
	Id      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

type StatusNode struct {
	Chain                string      `json:"chain"`
	Blocks               interface{} `json:"blocks"`
	Headers              interface{} `json:"headers"`
	Verificationprogress interface{} `json:"verificationprogress"`
	Softforks            struct {
		Bip34 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip34"`
		Bip66 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip66"`
		Bip65 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip65"`
		Csv struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"csv"`
		Segwit struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"segwit"`
		Taproot struct {
			Type string `json:"type"`
			Bip9 struct {
				Status              string      `json:"status"`
				StartTime           interface{} `json:"start_time"`
				Timeout             interface{} `json:"timeout"`
				Since               interface{} `json:"since"`
				MinActivationHeight int         `json:"min_activation_height"`
			} `json:"bip9"`

			Active bool `json:"active"`
		} `json:"taproot"`
	} `json:"softforks"`

	Warnings string `json:"warnings"`
}

type UnspentList struct {
	TxId         string `json:"txid"`
	Vout         int64  `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
}

type DecodedTx struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	Vsize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`

		Sequence int64 `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm     string `json:"asm"`
			Hex     string `json:"hex"`
			Address string `json:"address"`
			Type    string `json:"type"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
}

type UTXO []struct {
	TxId     string
	Vout     int64
	Amount   int64
	PKScript string
}

type Info struct {
	Walletname            string      `json:"walletname"`
	Walletversion         int         `json:"walletversion"`
	Format                string      `json:"format"`
	Balance               float64     `json:"balance"`
	UnconfirmedBalance    float64     `json:"unconfirmed_balance"`
	ImmatureBalance       float64     `json:"immature_balance"`
	Txcount               int         `json:"txcount"`
	Keypoololdest         int         `json:"keypoololdest"`
	Keypoolsize           int         `json:"keypoolsize"`
	Hdseedid              string      `json:"hdseedid"`
	KeypoolsizeHdInternal int         `json:"keypoolsize_hd_internal"`
	Paytxfee              float64     `json:"paytxfee"`
	PrivateKeysEnabled    bool        `json:"private_keys_enabled"`
	AvoidReuse            bool        `json:"avoid_reuse"`
	Scanning              interface{} `json:"scanning"`
	Descriptors           bool        `json:"descriptors"`
}

type Unspent struct {
	Txid          string  `json:"txid"`
	Vout          int     `json:"vout"`
	Address       string  `json:"address"`
	Label         string  `json:"label"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
	Spendable     bool    `json:"spendable"`
	Solvable      bool    `json:"solvable"`
	Safe          bool    `json:"safe"`
}

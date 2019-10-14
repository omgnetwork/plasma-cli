package childchain

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type TransactionGetResponse struct {
	Version string    `json:"version"`
	Success bool      `json:"success"`
	Data    GetTxData `json:"data"`
}

type GetTxData struct {
	Txindex  int       `json:"txindex"`
	Txhash   string    `json:"txhash"`
	Metadata string    `json:"metadata"`
	Txbytes  string    `json:"txbytes"`
	Block    Block     `json:"block"`
	Inputs   []Inputs  `json:"inputs"`
	Outputs  []Outputs `json:"outputs"`
}

type Block struct {
	Timestamp int    `json:"timestamp"`
	Hash      string `json:"hash"`
	EthHeight int    `json:"eth_height"`
	Blknum    int    `json:"blknum"`
}

//Get transaction
func (c *Client) GetTransaction(txHash string) (*TransactionGetResponse, error) {
	// client := &http.Client{}
	postData := map[string]interface{}{"id": txHash}
	rstring, err := c.do(
		"/transaction.get",
		postData,
	)
	response := TransactionGetResponse{}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		var errorInfo ClientError
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			return nil, err
		}
	}

	return &response, nil
}

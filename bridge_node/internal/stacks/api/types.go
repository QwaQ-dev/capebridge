package api

type HiroEventsResponse struct {
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
	Total   int         `json:"total"`
	Results []HiroEvent `json:"results"`
}

type HiroEvent struct {
	EventIndex  int          `json:"event_index"`
	EventType   string       `json:"event_type"`
	TxID        string       `json:"tx_id"`
	ContractLog *ContractLog `json:"contract_log,omitempty"`
}

type ContractLog struct {
	ContractID string     `json:"contract_id"`
	Topic      string     `json:"topic"`
	Value      ClarityVal `json:"value"`
}

type ClarityVal struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}

type HiroTxResponse struct {
	BlockHeight   int64  `json:"block_height"`
	SenderAddress string `json:"sender_address"`
	TxStatus      string `json:"tx_status"`
}

type ReadOnlyCallRequest struct {
	Sender    string   `json:"sender"`
	Arguments []string `json:"arguments"`
}

type ReadOnlyCallResponse struct {
	Okay   bool   `json:"okay"`
	Result string `json:"result"`
}

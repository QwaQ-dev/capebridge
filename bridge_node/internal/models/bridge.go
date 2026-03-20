package models

// DB structure
type BridgeEvent struct {
	ID          int64
	SourceChain string
	TargetChain string
	TxHash      string
	LogIndex    int
	BlockNumber int64
	Sender      string
	Receiver    string
	Amount      string
	Nonce       string
	Status      string
}

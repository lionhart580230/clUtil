package clTron

type Block struct {
	BlockId      string        `json:"blockID"`
	BlockHeader  BlockHeader   `json:"block_header"`
	Transactions []Transaction `json:"transactions"`
}

type BlockHeaderRawData struct {
	Number         uint64 `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Timestamp      uint64 `json:"timestamp"`
	Version        uint32 `json:"version"`
}

//type BlockHeaderRawData struct {
//	"number": 1000000,
//	"txTrieRoot": "e2dd42daa9c853e070df1e4fc927851a7438c601fb12cf945922629d5cec187b",
//	"witness_address": "41f16412b9a17ee9408646e2a21e16478f72ed1e95",
//	"parentHash": "00000000000f423fbccd9cdb9e410eabdf9f94145cf5b71a678b8b9616619125",
//	"version": 9,
//	"timestamp": 1578594852000
//}

type BlockHeader struct {
	RawData          BlockHeaderRawData `json:"raw_data"`
	WitnessSignature string             `json:"witness_signature"`
}
type TransactionRet struct {
	ContractRet string `json:"contractRet"`
}

type TransactionRawDataContractParamValue struct {
	Amount          uint64 `json:"amount"`
	OwnerAddress    string `json:"owner_address"`
	ToAddress       string `json:"to_address"`
	ContractAddress string `json:"contract_address"`
	Data            string `json:"data"`
}

type TransactionRawDataContractParam struct {
	Value   TransactionRawDataContractParamValue `json:"value"`
	TypeUrl string                               `json:"type_url"`
}
type TransactionRawDataContract struct {
	Parameter TransactionRawDataContractParam `json:"parameter"`
	Type      string                          `json:"type"`
}
type TransactionRawData struct {
	Contract []TransactionRawDataContract `json:"contract"`
}
type Transaction struct {
	Ret        []TransactionRet   `json:"ret"`
	Signature  []string           `json:"signature"`
	TxID       string             `json:"txID"`
	RawData    TransactionRawData `json:"raw_data"`
	RawDataHex string             `json:"raw_data_hex"`
}

package store

type GlobalStateType struct {
	CL CLType `json:"CL"`
	EL ELType `json:"EL"`
}

type CLType struct {
	Status        bool   `json:"status"`
	ValidatorAddr string `json:"validator_address"`
	Missed        string `json:"missed"`
}

type ELType struct {
	Status        bool   `json:"status"`
	Sync          bool   `json:"sync"`
	CurrentHeight uint64 `json:"current_height"`
	Peers         uint64 `json:"peers"`
	TxpoolQueued  int    `json:"txpool_queued"`
}

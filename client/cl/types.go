package cl

type Client struct {
	host string
}

type Response struct {
	Result  string      `json:"result"`
	Data    []Validator `json:"data"`
	QueryTs int         `json:"query_ts"`
}

type Validator struct {
	Addr             string `json:"addr"`
	CommitCount      uint64 `json:"commit_count"`
	FirstCommitBlock int    `json:"first_commit_block"`
	LastCommitBlock  int    `json:"last_commit_block"`
}

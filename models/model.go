package models

type SearchResult struct {
	ResultRank  int    `json:"resultRank"`
	ResultURL   string `json:"resultURL"`
	ResultTitle string `json:"resultTitle"`
	ResultDesc  string `json:"resultDesc"`
}

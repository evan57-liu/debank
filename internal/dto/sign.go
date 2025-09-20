package dto

type DebankSignDto struct {
	XApiSign  string `json:"x-api-sign"`
	XApiTs    string `json:"x-api-ts"`
	XApiNonce string `json:"x-api-nonce"`
	XApiVer   string `json:"x-api-ver"`
	Url       string `json:"url"`
}

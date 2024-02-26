package client

import (
	"ccavenue/aescbc"
	"encoding/hex"
	"encoding/json"
	"strings"
)

type PayoutSummary struct {
	IfscCode           string  `json:"ifsc_code,omitempty"`
	SettlementCurrency string  `json:"settlement_currency,omitempty"`
	SettlementDate     string  `json:"settlement_date,omitempty"`
	TransCurrency      string  `json:"trans_currency,omitempty"`
	UtrNo              string  `json:"utr_no,omitempty"`
	SettlementBank     string  `json:"settlement_bank,omitempty"`
	PayAmount          float64 `json:"pay_amount,omitempty"`
	SubAccId           string  `json:"sub_acc_Id,omitempty"`
	PayId              int64   `json:"pay_Id,omitempty"`
	AccountNo          int64   `json:"account_no,omitempty"`
}

type PayoutSummaryError struct {
	ErrorDesc string          `json:"error_desc,omitempty"`
	ErrorCode json.RawMessage `json:"error_code,omitempty"`
}

type PayoutFilter struct {
	SettlementDate string `json:"settlement_date,omitempty"`
}

func (f PayoutFilter) Encode() (string, error) {

	jsonBytes, err := json.Marshal(f)
	if err != nil {
		return "", err
	}

	encReqBytes, err := aescbc.NewCrypter().Encrypt(jsonBytes)
	if err != nil {
		return "", err
	}

	encStr := strings.ToUpper(hex.EncodeToString(encReqBytes))

	return encStr, nil
}

func (f PayoutFilter) Command() string {
	return "payoutSummary"
}

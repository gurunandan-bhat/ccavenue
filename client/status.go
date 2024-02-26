package client

import (
	"ccavenue/aescbc"
	"encoding/hex"
	"encoding/json"
	"strings"
)

type OrderValue struct {
	OrderCurrncy      string  `json:"order_currncy,omitempty"`
	OrderAmt          float64 `json:"order_amt,omitempty"`
	OrderFeePerc      float64 `json:"order_fee_perc,omitempty"`
	OrderFeePercValue float64 `json:"order_fee_perc_value,omitempty"`
	OrderFeeFlat      float64 `json:"order_fee_flat,omitempty"`
	OrderGrossAmt     float64 `json:"order_gross_amt,omitempty"`
	OrderDiscount     float64 `json:"order_discount,omitempty"`
	OrderTax          float64 `json:"order_tax,omitempty"`
	OrderBankRefNo    string  `json:"order_bank_ref_no,omitempty"`
	OrderGtwID        string  `json:"order_gtw_id,omitempty"`
	OrderBankResponse string  `json:"order_bank_response,omitempty"`
	OrderOptionType   string  `json:"order_option_type,omitempty"`
	OrderTDS          float64 `json:"order_TDS,omitempty"`
}
type Order struct {
	ReferenceNo     string `json:"reference_no,omitempty"`
	OrderNo         string `json:"order_no,omitempty"`
	OrderDateTime   string `json:"order_date_time,omitempty"`
	OrderDeviceType string `json:"order_device_type,omitempty"`
	// Billing Details
	OrderBillName    string `json:"order_bill_name,omitempty"`
	OrderBillAddress string `json:"order_bill_address,omitempty"`
	OrderBillCity    string `json:"order_bill_city,omitempty"`
	OrderBillState   string `json:"order_bill_state,omitempty"`
	OrderBillZip     string `json:"order_bill_zip,omitempty"`
	OrderBillCountry string `json:"order_bill_country,omitempty"`
	OrderBillTel     string `json:"order_bill_tel,omitempty"`
	OrderBillEmail   string `json:"order_bill_email,omitempty"`
	// Shipping Details
	OrderShipName    string `json:"order_ship_name,omitempty"`
	OrderShipAddress string `json:"order_ship_address,omitempty"`
	OrderShipCity    string `json:"order_ship_city,omitempty"`
	OrderShipState   string `json:"order_ship_state,omitempty"`
	OrderShipZip     string `json:"order_ship_zip,omitempty"`
	OrderShipCountry string `json:"order_ship_country,omitempty"`
	OrderShipTel     string `json:"order_ship_tel,omitempty"`
	OrderShipEmail   string `json:"order_ship_email,omitempty"`
	// Order Status
	OrderNotes          string  `json:"order_notes,omitempty"`
	OrderIP             string  `json:"order_ip,omitempty"`
	OrderStatus         string  `json:"order_status,omitempty"`
	OrderFraudStatus    string  `json:"order_fraud_status,omitempty"`
	OrderStatusDateTime string  `json:"order_status_date_time,omitempty"`
	OrderCaptAmt        float64 `json:"order_capt_amt,omitempty"`
	OrderCardName       string  `json:"order_card_name,omitempty"`
	// Merchant Params
	MerchantParam1 string `json:"merchant_param1,omitempty"`
	MerchantParam2 string `json:"merchant_param2,omitempty"`
	MerchantParam3 string `json:"merchant_param3,omitempty"`
	MerchantParam4 string `json:"merchant_param4,omitempty"`
	MerchantParam5 string `json:"merchant_param5,omitempty"`
}

type OrderError struct {
	ErrorDesc string `json:"error_desc,omitempty"`
	Status    int    `json:"status,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
}
type StatusFilter struct {
	OrderNo string `json:"order_no,omitempty"`
}

func (f StatusFilter) Encode() (string, error) {

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

func (f StatusFilter) Command() string {
	return "orderStatusTracker"
}

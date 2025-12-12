package models

type OrderValidateRequest struct {
	ValId       string `json:"val_id" url:"val_id"`
	StoreId     string `json:"store_id" url:"store_id"`
	StorePasswd string `json:"store_passwd" url:"store_passwd"`
	Format      string `json:"format" url:"format"`
	V           string `json:"v" url:"v"`
}

type TransactionQueryBySessionKeyRequest struct {
	Sessionkey  string `json:"sessionkey" url:"sessionkey"`
	StoreId     string `json:"store_id" url:"store_id"`
	StorePasswd string `json:"store_passwd" url:"store_passwd"`
}

type TransactionQueryByTransactionIdRequest struct {
	Sessionkey  string `json:"sessionkey" url:"sessionkey"`
	StoreId     string `json:"store_id" url:"store_id"`
	StorePasswd string `json:"store_passwd" url:"store_passwd"`
}

type InitiateRefundRequest struct {
	BankTranId   string `json:"bank_tran_id" url:"bank_tran_id"`
	StoreId      string `json:"store_id" url:"store_id"`
	StorePasswd  string `json:"store_passwd" url:"store_passwd"`
	RefundAmount string `json:"refund_amount" url:"refund_amount"`
	RefeId       string `json:"refe_id" url:"refe_id"`
	Format       string `json:"format" url:"format"`
}

type RefundQueryURLRequest struct {
	RefundRefId string `json:"refund_ref_id" url:"refund_ref_id"`
	StoreId     string `json:"store_id" url:"store_id"`
	StorePasswd string `json:"store_passwd" url:"store_passwd"`
}

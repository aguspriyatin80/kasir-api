package dto

type TransactionReport struct {
	TransactionID int    `json:"transaction_id"`
	Tanggal       string `json:"tanggal"`
	TotalAmount   int    `json:"total_amount"`
}

type ReportResponse struct {
	StartDate      string              `json:"start_date"`
	EndDate        string              `json:"end_date"`
	TotalTransaksi int                 `json:"total_transaksi"`
	TotalRevenue   int                 `json:"total_revenue"`
	Data           []TransactionReport `json:"data"`
}

package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	var (
		res *models.Transaction
	)

	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// inisialisasi subtotal -> jumlah total transaksi keseluruhan
	totalAmount := 0
	// inisialisasi modeling transactionDetails -> nanti kita insert ke db
	details := make([]models.TransactionDetail, 0)
	// loop setiap item
	for _, item := range items {
		var productName string
		var productID, price, stock int
		// get product dapet pricing
		err := tx.QueryRow("SELECT id, name, price, stock FROM products WHERE id=$1", item.ProductID).Scan(&productID, &productName, &price, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}

		if err != nil {
			return nil, err
		}

		// hitung current total = quantity * pricing
		// ditambahin ke dalam subtotal
		subtotal := item.Quantity * price
		totalAmount += subtotal

		// kurangi jumlah stok
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, productID)
		if err != nil {
			return nil, err
		}

		// item nya dimasukkin ke transactionDetails
		details = append(details, models.TransactionDetail{
			ProductID:   productID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// insert transaction
	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING ID", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// insert transaction details
	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	res = &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}

	return res, nil
}

func (repo *TransactionRepository) SalesSummaryToday() (models.RecapToday, error) {
	var totalPenjualan int
	var totalTransaksi int
	var productName string
	var quantity int

	// 1. Total Penjualan Hari Ini
	repo.db.QueryRow(`
			SELECT COALESCE(SUM(total_amount), 0)
			FROM transactions
			WHERE created_at::date = CURRENT_DATE
		`).Scan(&totalPenjualan)

	// 2. Total Transaksi Hari Ini
	repo.db.QueryRow(`
			SELECT COUNT(*)
			FROM transactions
			WHERE created_at::date = CURRENT_DATE
		`).Scan(&totalTransaksi)

	// 3. Produk Terlaris
	repo.db.QueryRow(`
			SELECT p.name, SUM(td.quantity) AS qty_terjual
			FROM transaction_details td
			JOIN transactions t ON t.id = td.transaction_id
			JOIN products p ON p.id = td.product_id
			WHERE t.created_at::date = CURRENT_DATE
			GROUP BY p.name
			ORDER BY qty_terjual DESC
			LIMIT 1
		`).Scan(&productName, &quantity)

	var produkTerlaris models.ProdukTerlaris
	produkTerlaris = models.ProdukTerlaris{Name: productName, QtyTerjual: quantity}
	response := models.RecapToday{
		TotalPenjualan: totalPenjualan,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: produkTerlaris,
	}

	return response, nil
}

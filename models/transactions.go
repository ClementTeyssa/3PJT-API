package models

import (
	"log"
	"time"

	"github.com/ClementTeyssa/3PJT-API/config"
)

type Transaction struct {
	ID          int       `json:"id" validate:"omitempty,uuid"`
	AccountFrom string    `json:"accountfrom" validate:"required"`
	AccountTo   string    `json:"accountto" validate:"required"`
	Amount      float32   `json:"amount" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transactions []Transaction

func NewTransaction(transaction *Transaction) {
	if transaction == nil {
		log.Panic(transaction)
	}
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	err := config.GetDb().QueryRow("INSERT INTO transactions (accountfrom, accountto, amount, created_at, updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id;", transaction.AccountFrom, transaction.AccountTo, transaction.Amount, transaction.CreatedAt, transaction.UpdatedAt).Scan(&transaction.ID)

	if err != nil {
		log.Panic(err)
	}
}

func FindTransactionById(id int) *Transaction {
	var transaction Transaction
	row := config.GetDb().QueryRow("SELECT * FROM transactions WHERE id = $1;", id)
	err := row.Scan(&transaction.ID, &transaction.AccountFrom, &transaction.AccountTo, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		log.Panic(err)
	}
	return &transaction
}

func CountTransactionsById(id int) int {
	rows, err := config.GetDb().Query("SELECT COUNT(*) as count FROM transactions WHERE id = $1;", id)

	if err != nil {
		log.Panic(err)
	}

	count := 0
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Panic(err)
		}
	}

	return count
}

func AllTransactions() *Transactions {
	var transactions Transactions
	rows, err := config.GetDb().Query("SELECT * FROM transactions")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.ID, &transaction.AccountFrom, &transaction.AccountTo, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			log.Panic(err)
		}
		transactions = append(transactions, transaction)
	}
	return &transactions
}

func UpdateTransaction(transaction *Transaction) {
	transaction.UpdatedAt = time.Now()
	stmt, err := config.GetDb().Prepare("UPDATE transactions SET accountfrom=$1, accountto=$2, amount=$3, updated_at=$4 WHERE id=$5;")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(transaction.AccountFrom, transaction.AccountTo, transaction.Amount, transaction.UpdatedAt)
	if err != nil {
		log.Panic(err)
	}
}

func DeleteTransactionById(id int) error {
	stmt, err := config.GetDb().Prepare("DELETE FROM transactions WHERE id=$1;")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(id)
	return err
}

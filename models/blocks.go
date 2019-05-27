package models

import (
	"log"
	"time"

	"github.com/ClementTeyssa/3PJT-API/config"
)

type Block struct {
	ID            int       `json:"id" validate:"omitempty,uuid"`
	Timestamp     int       `json:"timestamp" validate:"required"`
	TransactionID int       `json:"transactionid" validate:"required"`
	Hash          string    `json:"hash" validate:"required"`
	PrevHash      string    `json:"prevhash" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Blocks []Block

func NewBlock(block *Block) {
	if block == nil {
		log.Panic(block)
	}
	block.CreatedAt = time.Now()
	block.UpdatedAt = time.Now()
	err := config.GetDb().QueryRow("INSERT INTO blocks (timestamp, transactionid, hash, prevhash, created_at, updated_at) VALUES ($1,$2,$3,$4,$5, $6) RETURNING id;", block.Timestamp, block.TransactionID, block.Hash, block.PrevHash, block.CreatedAt, block.UpdatedAt).Scan(&block.ID)

	if err != nil {
		log.Panic(err)
	}
}

func FindBlockById(id int) *Block {
	var block Block
	row := config.GetDb().QueryRow("SELECT * FROM blocks WHERE id = $1;", id)
	err := row.Scan(&block.ID, &block.Timestamp, &block.TransactionID, &block.Hash, &block.PrevHash, &block.CreatedAt, &block.UpdatedAt)

	if err != nil {
		log.Panic(err)
	}
	return &block
}

func AllBlocks() *Blocks {
	var blocks Blocks
	rows, err := config.GetDb().Query("SELECT * FROM blocks")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var block Block
		err := rows.Scan(&block.ID, &block.Timestamp, &block.TransactionID, &block.Hash, &block.PrevHash, &block.CreatedAt, &block.UpdatedAt)
		if err != nil {
			log.Panic(err)
		}
		blocks = append(blocks, block)
	}
	return &blocks
}

func UpdateBlock(block *Block) {
	block.UpdatedAt = time.Now()
	stmt, err := config.GetDb().Prepare("UPDATE blocks SET timestamp=$1, transactionid=$2, hash=$3, prevhash=$4, updated_at=$5 WHERE id=$6;")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(block.Timestamp, block.TransactionID, block.Hash, block.PrevHash, block.CreatedAt, block.UpdatedAt)
	if err != nil {
		log.Panic(err)
	}
}

func DeleteBlockById(id int) error {
	stmt, err := config.GetDb().Prepare("DELETE FROM blocks WHERE id=$1;")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(id)
	return err
}

package seeders

import (
	"database/sql"
	"log"
	"time"
)

type Seed struct {
	db *sql.DB
}

func NewSeed(db *sql.DB) Seed {
	return Seed{
		db: db,
	}
}

func (s Seed) ExecuteSeeder() {
	s.CategorySeeder()
	s.UserSeeder()
	s.WalletSeeder()
	s.TransactionSeeder()
}

func (s Seed) UserSeeder() {
	var err error
	sql := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	_, err = s.db.Exec(sql,
		"Dummy User",
		"dummy@email.com",
		"password",
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Fatalf("error seeding user: %v", err)
	}
}

func (s Seed) CategorySeeder() {
	var err error
	sql := "INSERT INTO categories (type, name, icon, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err = s.db.Exec(sql,
		"E",
		"Food",
		"",
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Fatalf("error seeding category: %v", err)
	}
	_, err = s.db.Exec(sql,
		"I",
		"Salary",
		"",
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Fatalf("error seeding category: %v", err)
	}
}

func (s Seed) WalletSeeder() {
	var err error
	sql := "INSERT INTO wallets (name, icon, currency, balance, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err = s.db.Exec(sql,
		"My Wallet",
		"",
		"IDR",
		1000,
		"1",
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Fatalf("error seeding wallet: %v", err)
	}
}

func (s Seed) TransactionSeeder() {
	var err error
	sql := "INSERT INTO transactions (wallet_id, category_id, nominal, date_time, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err = s.db.Exec(sql,
		1,
		2,
		1000,
		time.Now(),
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Fatalf("error seeding transaction: %v", err)
	}
}

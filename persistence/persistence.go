package persistence

import (
	"database/sql"
	"fmt"
	"mfa/migration"
	"strconv"
)

func NewDB() *sql.DB {
	dbDriver := "pgx"

	db, err := sql.Open(dbDriver, "host=localhost port=5532 user=mfa_admin_user password=mfa_admin_pass dbname=mfa_demo sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("could not open db %v", err))
	}
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("unable to ping db %v", err))
	}

	// Migrate Database
	if err := migration.MigrateTo(db, dbDriver); err != nil {
		panic(fmt.Sprintf("datbase migration failed %v", err))
	}
	return db
}

type Repository interface {
	GetSharedSecret(userId string) (string, error)
	SaveSharedSecret(userId string, sharedSecret string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (r *repository) GetSharedSecret(userId string) (string, error) {
	uid, err := strconv.Atoi(userId)
	if err != nil {
		return "", err
	}
	var ss string
	query := "SELECT shared_secret FROM user_mfa_shared_secret WHERE user_id = $1"
	row := r.DB.QueryRow(query, uid)
	err = row.Scan(&ss)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (r *repository) SaveSharedSecret(userId string, sharedSecret string) error {
	uid, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}
	query := "INSERT INTO user_mfa_shared_secret (user_id, shared_secret) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET shared_secret = $3"
	_, err = r.DB.Exec(query, uid, sharedSecret, sharedSecret)
	if err != nil {
		return err
	}
	return nil
}

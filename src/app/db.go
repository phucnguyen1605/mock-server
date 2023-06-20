package app

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	QueryRow(dest interface{}, sql string, values ...interface{}) error
	Query(query string, arg ...interface{}) (*sqlx.Rows, error)
	Exec(query string, arg ...interface{}) error
}

type MySQL struct {
	url string
	db  *sqlx.DB
}

var (
	// MySQL singleton instance
	dbInstance *MySQL
	once       sync.Once
)

// newDB is factory func for creating db instance
func newDB(env map[string]string) DB {
	url, _ := env["DB_CONNECTION_STRING"]
	if url == "" {
		log.Println("No db setting")
		return nil
	}

	once.Do(func() {
		dbInstance = &MySQL{
			url: url,
		}
	})

	return dbInstance
}

func (m *MySQL) connect() error {
	if m.db == nil {
		log.Printf("Connecting db...")
		db, err := sqlx.Connect("mysql", m.url)
		if err != nil {
			log.Printf("%v", err)
			return err
		}

		m.db = db
		db.SetMaxIdleConns(10)
	}
	log.Printf("The number of established connections both in use and idle: %d", m.db.Stats().OpenConnections)
	return nil
}

func (m *MySQL) close() {
	if m.db != nil {
		m.db.Close()
		m.db = nil
		log.Printf("Closed DB connection")
	}
}

// QueryRow returns sinble object
func (m *MySQL) QueryRow(dest interface{}, sql string, values ...interface{}) error {
	if err := m.connect(); err != nil {
		return err
	}

	row := m.db.QueryRowx(sql, values...)
	err := row.StructScan(dest)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}

// Query returns array of object
func (m *MySQL) Query(query string, arg ...interface{}) (*sqlx.Rows, error) {
	if err := m.connect(); err != nil {
		return nil, err
	}

	var rows *sqlx.Rows
	var err error
	rows, err = m.db.Queryx(query, arg...)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Exec executes named query
func (m *MySQL) Exec(query string, arg ...interface{}) error {
	if err := m.connect(); err != nil {
		return err
	}

	_, err := m.db.Exec(query, arg...)
	return err
}

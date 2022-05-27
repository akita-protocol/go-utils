package db

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"git.keelakay.com/keelakay/utils/errors"
	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const pkg errors.Pkg = "db"

// Connect connects to a mysql database
func Connect() (*sqlx.DB, error) {
	const op errors.Op = "Connect"

	credentials, err := base64.StdEncoding.DecodeString(os.Getenv("DB_AUTH"))
	if err != nil {
		return nil, errors.E(pkg, op, errors.Database, err)
	}

	db, err := sqlx.Open("mysql", string(credentials))
	if err != nil {
		return nil, errors.E(pkg, op, errors.Database, err)
	}
	db.SetMaxOpenConns(10)
	return db, nil
}

type ConnectionHandle interface {
	*sqlx.DB | *sqlx.Tx
}

// Insert inserts an object into the database
func Insert[H ConnectionHandle, S DBObject](h H, s S) (int64, error) {
	const op errors.Op = "Insert"
	var (
		res    sql.Result
		err    error
		query  = ""
		keys   = ""
		values []interface{}
	)

	table := getTable(s)
	if table == "" {
		return 0, errors.E(pkg, op, errors.Database, fmt.Errorf("invalid table struct"))
	}

	params := structs.Map(s)

	for k, v := range params {
		keys += fmt.Sprintf("%s, ", k)
		values = append(values, v)
	}
	q := strings.Repeat("?, ", len(values))
	query = fmt.Sprintf("insert into %s (%s) values (%s)", table, keys[0:len(keys)-2], q[0:len(q)-2])

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return 0, errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		res, err = stmt.Exec(values...)
		if err != nil {
			return 0, errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		res, err = h.Exec(query, values...)
		if err != nil {
			return 0, errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.E(pkg, op, errors.Database, err, "Failed to Retrieve ID")
	}

	return id, nil
}

// Update updates an object in the database
func Update[H ConnectionHandle, S DBObject](h H, s S, where map[string]interface{}) (*sql.Result, error) {
	const op errors.Op = "Update"
	var (
		res    sql.Result
		err    error
		query  = ""
		keys   = ""
		w      = ""
		values []interface{}
	)

	table := getTable(s)
	if table == "" {
		return nil, errors.E(pkg, op, errors.Database, fmt.Errorf("invalid table struct"))
	}

	params := structs.Map(s)

	for k, v := range params {
		keys += fmt.Sprintf("%s=?, ", k)
		values = append(values, v)
	}

	for k, v := range where {
		w += fmt.Sprintf("%s=? and ", k)
		values = append(values, v)
	}

	query = fmt.Sprintf("update %s set %s where %s", table, keys[0:len(keys)-2], w[0:len(w)-5])

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return nil, errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		res, err = stmt.Exec(values...)
		if err != nil {
			return nil, errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		res, err = h.Exec(query, values...)
		if err != nil {
			return nil, errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return &res, nil
}

// Upsert inserts into or updates an object in the database
func Upsert[H ConnectionHandle, S DBObject](h H, s S) (int64, error) {
	const op errors.Op = "Upsert"
	var (
		res    sql.Result
		err    error
		query  = ""
		keys   = ""
		dupe   = ""
		values []interface{}
	)

	table := getTable(s)
	if table == "" {
		return 0, errors.E(pkg, op, errors.Database, fmt.Errorf("invalid table struct"))
	}

	params := structs.Map(s)
	_, id_preset := params["id"]
	if id_preset {
		return 0, errors.E(pkg, op, errors.Database, fmt.Errorf("dont use upsert with preset ids"))
	}

	for k, v := range params {
		keys += fmt.Sprintf("%s, ", k)
		dupe += fmt.Sprintf("%s = VALUES(%s), ", k, k)
		values = append(values, v)
	}

	q := strings.Repeat("?, ", len(values))

	query = fmt.Sprintf("insert into %s (%s) values (%s) on duplicate key update id=LAST_INSERT_ID(id), %s", table, keys[0:len(keys)-2], q[0:len(q)-2], dupe[0:len(dupe)-2])

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return 0, errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		res, err = stmt.Exec(values...)
		if err != nil {
			return 0, errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		res, err = h.Exec(query, values...)
		if err != nil {
			return 0, errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.E(pkg, op, errors.Database, err, "Failed to Retrieve ID")
	}

	return id, nil
}

// Delete removes an object from the database
func Delete[H ConnectionHandle, S DBObject](h H, s S) (*sql.Result, error) {
	const op errors.Op = "Delete"
	var (
		res   sql.Result
		err   error
		query = ""
	)

	table := getTable(s)
	if table == "" {
		return nil, errors.E(pkg, op, errors.Database, fmt.Errorf("invalid table struct"))
	}

	params := structs.Map(s)

	id, ok := params["id"]
	if !ok {
		return nil, errors.E(pkg, op, errors.Database, fmt.Errorf("an object id is required to delete by object"))
	}

	query = fmt.Sprintf("delete from %s where id = ?", table)

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return nil, errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		res, err = stmt.Exec(id)
		if err != nil {
			return nil, errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		res, err = h.Exec(query, id)
		if err != nil {
			return nil, errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return &res, nil
}

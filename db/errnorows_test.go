package db

import (
	"database/sql"
	stderrors "errors"
	"fmt"
	"net"
	"testing"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

func TestErrNoRows(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"sql.ErrNoRows", sql.ErrNoRows, true},
		{"wrapped sql.ErrNoRows", fmt.Errorf("get: %w", sql.ErrNoRows), true},
		{"utils not found", errors.E(errors.Op("x"), errors.DatabaseResultNotFound), true},
		{"utils other kind", errors.E(errors.Op("x"), errors.Database), false},
		{"connection error", &net.OpError{Op: "dial", Net: "tcp", Err: stderrors.New("connection refused")}, false},
		{"plain error", stderrors.New("boom"), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ErrNoRows(c.err); got != c.want {
				t.Fatalf("ErrNoRows(%v) = %v, want %v", c.err, got, c.want)
			}
		})
	}
}

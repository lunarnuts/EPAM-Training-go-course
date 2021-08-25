package handlers

import (
	"net/http"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestInsert(t *testing.T) {
	type args struct {
		p *pgxpool.Pool
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.p, tt.args.w, tt.args.r)
		})
	}
}

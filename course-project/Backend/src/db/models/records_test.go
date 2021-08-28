package records

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/db"
	"github.com/stretchr/testify/assert"
)

type mockDBConnection struct {
}

type mockRow struct {
	id            uint64
	cityName      string
	timeRequested string
	temperature   float64
}

func (this mockRow) Scan(dest ...interface{}) error {
	if len(dest) == 1 {
		Id := dest[0].(*uint64)
		*Id = this.id
		return nil
	}
	Id := dest[0].(*uint64)
	cityName := dest[1].(*string)
	timeRequested := dest[2].(*string)
	temperature := dest[3].(*float64)
	*Id = this.id
	*cityName = this.cityName
	*timeRequested = this.timeRequested
	*temperature = this.temperature
	return nil
}

func (this mockRow) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, nil
}

func (this mockRow) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	columns := []string{"id", "cityname", "timerequested", "temperature"}
	return pgxpoolmock.NewRows(columns).AddRow(0, "", "", 0.0).ToPgxRows(), nil
}

func (this mockRow) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {

	return mockRow{1, "Arkham", "batman time", 1000.0}
}
func TestSelect(t *testing.T) {
	conn := mockRow{}
	got, err := Select(conn, 1)
	want := Record{
		Id:            1,
		CityName:      "Arkham",
		TimeRequested: "batman time",
		Temperature:   1000.0,
	}
	if !assert.Equal(t, got, want) {
		t.Errorf("Select() got = %v, want = %v", got, want)
	}
	if err != nil {
		t.Errorf("Select() got = %v, want = %v", got, want)
	}
}

func TestSelectAll(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "cityname", "timerequested", "temperature"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow(uint64(1), "Arkham", "mock time", 27.0).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	mockPool2 := pgxpoolmock.NewMockPgxPool(ctrl)
	pgxRows2 := pgxpoolmock.NewRows(columns).AddRow(uint64(1), "San Andreas", "mock time", 19.0).AddRow(uint64(2), "Arkham", "mock time", 27.0).ToPgxRows()
	mockPool2.EXPECT().Query(gomock.Any(), gomock.Any()).Return(pgxRows2, nil)

	mockPool3 := pgxpoolmock.NewMockPgxPool(ctrl)
	pgxRows3 := pgxpoolmock.NewRows(columns).AddRow(uint64(0), "", "", 0.0).ToPgxRows()
	mockPool3.EXPECT().Query(gomock.Any(), gomock.Any()).Return(pgxRows3, pgx.ErrNoRows)

	type args struct {
		conn db.DBConn
	}
	tests := []struct {
		name    string
		args    args
		want    []Record
		wantErr error
	}{
		{name: "Arkham", args: args{conn: mockPool}, want: []Record{{Id: 1, CityName: "Arkham", TimeRequested: "mock time", Temperature: 27.0}}, wantErr: nil},
		{name: "San Andreas", args: args{conn: mockPool2}, want: []Record{{Id: 1, CityName: "San Andreas", TimeRequested: "mock time", Temperature: 19.0},
			{Id: 2, CityName: "Arkham", TimeRequested: "mock time", Temperature: 27.0}}, wantErr: nil},
		{name: "Empty", args: args{conn: mockPool3}, want: []Record{}, wantErr: pgx.ErrNoRows},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectAll(tt.args.conn)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("SelectAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	res := pgconn.CommandTag{byte(100)}
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(res, nil)

	mockPool2 := pgxpoolmock.NewMockPgxPool(ctrl)
	res2 := pgconn.CommandTag{byte(0)}
	mockPool2.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(res2, pgx.ErrNoRows)

	type args struct {
		conn db.DBConn
		id   uint64
		rec  Record
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "first", args: args{conn: mockPool}, wantErr: false},
		{name: "error", args: args{conn: mockPool2}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(tt.args.conn, tt.args.id, tt.args.rec); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	res := pgconn.CommandTag{byte(100)}
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(res, nil)

	mockPool2 := pgxpoolmock.NewMockPgxPool(ctrl)
	res2 := pgconn.CommandTag{byte(0)}
	mockPool2.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(res2, pgx.ErrNoRows)

	type args struct {
		conn db.DBConn
		id   uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "first", args: args{conn: mockPool}, wantErr: false},
		{name: "error", args: args{conn: mockPool2}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.conn, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	res := mockRow{
		id:            1,
		cityName:      "Arkham",
		timeRequested: "mock time",
		temperature:   0.0,
	}
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(res)

	mockPool2 := pgxpoolmock.NewMockPgxPool(ctrl)
	res2 := mockRow{
		id:            10,
		cityName:      "City of Sin",
		timeRequested: "666",
		temperature:   10000.0,
	}

	mockPool2.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(res2)

	type args struct {
		conn db.DBConn
		rec  Record
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{name: "first", args: args{conn: mockPool, rec: Record{
			Id:            1,
			CityName:      "Arkham",
			TimeRequested: "batman time",
			Temperature:   1000.0,
		}}, want: 1, wantErr: false},
		{name: "error", args: args{conn: mockPool2, rec: Record{
			Id:            10,
			CityName:      "City of Sin",
			TimeRequested: "666",
			Temperature:   10000.0,
		}}, want: 10, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Insert(tt.args.conn, tt.args.rec)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

package records

import (
	"context"
	"fmt"

	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Record struct {
	Id            int     `json:"id" db:"id"`
	CityName      string  `json:"cityName" db:"cityName"`
	TimeRequested string  `json:"timeRequested" db:"timeRequested"`
	Temperature   float64 `json:"temperature" db:"temperature"`
}

func SelectAll(p *pgxpool.Pool) ([]Record, error) {
	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return []Record{}, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		"SELECT * FROM logbook")
	if err != nil {
		log.Printf("Query failed with: %v", err)
	}
	var recs []Record
	for rows.Next() {
		var rec Record
		err = rows.Scan(&rec.Id, &rec.CityName, &rec.TimeRequested, &rec.Temperature)
		if err == pgx.ErrNoRows {
			fmt.Printf("No rows: %v", err)
			return recs, err
		}

		if err != nil {
			log.Printf("Unable to SELECT: %v", err)
			return recs, err
		}
		recs = append(recs, rec)
	}
	log.Println(recs)
	return recs, nil
}

func Select(p *pgxpool.Pool, id uint64) (Record, error) {
	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return Record{}, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"SELECT id, cityName, timeRequested, temperature FROM logbook WHERE id = $1",
		id)

	var rec Record
	err = row.Scan(&rec.Id, &rec.CityName, &rec.TimeRequested, &rec.Temperature)
	if err == pgx.ErrNoRows {
		return rec, err
	}

	if err != nil {
		log.Printf("Unable to SELECT: %v", err)
		return rec, err
	}
	return rec, nil
}

func Insert(p *pgxpool.Pool, rec Record) (uint64, error) {

	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v", err)
		return 0, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"INSERT INTO logbook (cityName, timeRequested) VALUES ($1, $2) RETURNING id",
		rec.CityName, rec.TimeRequested)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v", err)
		return 0, err
	}

	return id, nil
}

func Update(p *pgxpool.Pool, id uint64, rec Record) error {

	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v", err)
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		"UPDATE logbook SET cityName = $2, timeRequested = $3, temperature = $4 WHERE id = $1",
		id, rec.CityName, rec.TimeRequested, rec.Temperature)
	if err != nil {
		log.Printf("Unable to UPDATE: %v\n", err)
		return err
	}

	if ct.RowsAffected() == 0 {
		return err
	}
	return nil
}

func Delete(p *pgxpool.Pool, id uint64) error {

	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v", err)
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), "DELETE FROM logbook WHERE id = $1", id)
	if err != nil {
		log.Printf("Unable to DELETE: %v", err)
		return err
	}

	if ct.RowsAffected() == 0 {
		return err
	}

	return nil
}

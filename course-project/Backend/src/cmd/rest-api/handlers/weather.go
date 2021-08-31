package handlers

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/cmd/rest-api/lib"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/db"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/models"
)

func GetCurrentWeather(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	cityName := lib.CityNameFromVars(r)
	timeRequested := time.Now()
	temperature := 0.0
	rec := records.Record{
		CityName:      cityName,
		TimeRequested: timeRequested.String(),
		Temperature:   float64(temperature),
	}
	conn, err := db.AcquireConn(p)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	id, err := records.Insert(conn, rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}

	weather, err := lib.GetCurrentWeatherFromAPI(cityName)
	if err != nil {
		lib.ReturnInternalError(w, err)
	}
	rec.Temperature = weather.Temperature
	err = records.Update(conn, id, rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	rec, err = records.Select(conn, id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	lib.ReturnJSON(w, rec)
}

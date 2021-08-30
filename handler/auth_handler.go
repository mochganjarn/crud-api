package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rysmaadit/go-template/app"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/contract"
	"github.com/rysmaadit/go-template/external/gorm_client"
	"github.com/rysmaadit/go-template/service"
	log "github.com/sirupsen/logrus"
)

func GetToken(authService service.AuthServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := authService.GetToken()

		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, resp, nil)
	}
}

func ValidateToken(authService service.AuthServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := contract.NewValidateTokenRequest(r)

		if err != nil {
			log.Warning(err)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		resp, err := authService.VerifyToken(req)

		if err != nil {
			log.Error(err)
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, resp, nil)
		return
	}
}

func Movie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := "movie"

		responder.NewHttpResponse(r, w, http.StatusOK, resp, nil)
	}
}

func ShowMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramsURL := mux.Vars(r)
		//membuat koneksi ke database
		db, err := gorm_client.Connection(app.Init())
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
		}

		//migrasi struct move ke database
		db.AutoMigrate(gorm_client.Movie{})

		movie := gorm_client.Movie{}
		//get movie berdasarkan kolom slug
		result := db.First(&movie, "slug = ?", paramsURL["slug"])

		if result == nil {
			responder.NewHttpResponse(r, w, http.StatusNotFound, nil, nil)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, result.RowsAffected, nil)
	}
}

func CreateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//membuat koneksi ke database
		db, err := gorm_client.Connection(app.Init())
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
		}

		//migrasi struct move ke database
		db.AutoMigrate(gorm_client.Movie{})

		// inisialisasi input dari form params
		title := r.FormValue("title")
		slug := r.FormValue("slug")
		description := r.FormValue("description")
		duration, _ := strconv.Atoi(r.FormValue("duration"))
		image := r.FormValue("image")

		//inisialisasi struct movie sesuai input params
		movie := gorm_client.Movie{
			Title:       title,
			Slug:        slug,
			Description: description,
			Duration:    duration,
			Image:       image,
		}

		//insert data kedalam database
		db.Create(&movie)

		responder.NewHttpResponse(r, w, http.StatusCreated, movie, nil)
	}
}

func UpdateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramSlug := mux.Vars(r)
		//membuat koneksi ke database
		db, err := gorm_client.Connection(app.Init())
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
		}

		//migrasi struct move ke database
		db.AutoMigrate(gorm_client.Movie{})

		// inisialisasi input dari form params
		title := r.FormValue("title")
		slug := r.FormValue("slug")
		description := r.FormValue("description")
		duration, _ := strconv.Atoi(r.FormValue("duration"))
		image := r.FormValue("image")

		//inisialisasi struct movie sesuai input params
		movie := gorm_client.Movie{}

		//update data kedalam database
		db.Model(&movie).Where("slug = ?", paramSlug["slug"]).Updates(gorm_client.Movie{Title: title, Slug: slug,
			Description: description, Duration: duration,
			Image: image})
		responder.NewHttpResponse(r, w, http.StatusOK, movie, nil)
	}
}

func DeleteMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramsURL := mux.Vars(r)
		//membuat koneksi ke database
		db, err := gorm_client.Connection(app.Init())
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
		}

		//migrasi struct move ke database
		db.AutoMigrate(gorm_client.Movie{})

		movie := gorm_client.Movie{}

		//Delete movie berdasarkan slug
		db.Where("slug = ?", paramsURL["slug"]).Delete(&movie)
		responder.NewHttpResponse(r, w, http.StatusOK, "success", nil)
	}
}

package app

import (
	"bass/model"
	"bass/repository"
	"bass/util/validator"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func (app *App) HandleListBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := repository.ListBookings(app.db)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}

	if bookings == nil {
		fmt.Fprintf(w, "[]")
		return
	}

	dtos := bookings.ToDto()
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		return
	}
}

func (app *App) HandleCreateBooking(w http.ResponseWriter, r *http.Request) {
	form := &model.BookingForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}

	if err := app.validator.Struct(form); err != nil {
		app.logger.Warn().Err(err).Msg("")

		resp := validator.ToErrResponse(err)

		if resp == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error", "%v"}`, appErrFormErrResponseFailure)
			return
		}

		respBody, err := json.Marshal(resp)
		if err != nil {
			app.logger.Warn().Err(err).Msg("")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
			return
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(respBody)
		return
	}

	bookingModel, err := form.ToModel()

	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}

	booking, err := repository.CreateBooking(app.db, bookingModel)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataCreationFailure)
		return
	}

	app.logger.Info().Msgf("New booking created: %d", booking.ID)
	w.WriteHeader(http.StatusCreated)
}

func (app *App) HandleReadBooking(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || id == 0 {
		app.logger.Info().Msgf("can not parse ID: %v", id)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	booking, err := repository.ReadBooking(app.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}

	dto := booking.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		return
	}
}

func (app *App) HandleUpdateBooking(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || id == 0 {
		app.logger.Info().Msgf("can not parse ID: %v", id)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	form := &model.BookingForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}

	if err := app.validator.Struct(form); err != nil {
		app.logger.Warn().Err(err).Msg("")

		resp := validator.ToErrResponse(err)
		if resp == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, appErrFormErrResponseFailure)
			return
		}

		respBody, err := json.Marshal(resp)
		if err != nil {
			app.logger.Warn().Err(err).Msg("")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
			return
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(respBody)
		return
	}

	if err := app.validator.Struct(form); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, err.Error())
		return
	}

	bookingModel, err := form.ToModel()
	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}

	bookingModel.ID = uint(id)
	if err := repository.UpdateBooking(app.db, bookingModel); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataUpdateFailure)
		return
	}

	app.logger.Info().Msgf("Booking updated: %d", id)
	w.WriteHeader(http.StatusAccepted)
}

func (app *App) HandleDeleteBooking(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || id == 0 {
		app.logger.Info().Msgf("can not parse ID: %v", id)

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := repository.DeleteBooking(app.db, uint(id)); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}

	app.logger.Info().Msgf("Booking deleted: %d", id)
	w.WriteHeader(http.StatusAccepted)
}

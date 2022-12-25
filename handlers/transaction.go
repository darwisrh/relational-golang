package handlers

import (
	dto "dewetour/dto/result"
	transactiondto "dewetour/dto/transaction"
	"dewetour/models"
	"dewetour/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
)

// var path_file = "http://localhost:5000/uploads/"

type handleTransac struct {
	TransactionRepo repositories.TransactionRepo
}

func HandleTransac(TransactionRepo repositories.TransactionRepo) *handleTransac {
	return &handleTransac{TransactionRepo}
}

func (h *handleTransac) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// dataContex := r.Context().Value("dataFile")
	// filename := dataContex.(string)

	tripId, _ := strconv.Atoi(r.FormValue("tripid"))
	userid, _ := strconv.Atoi(r.FormValue("userid"))
	counterqty, _ := strconv.Atoi(r.FormValue("counterqty"))
	total, _ := strconv.Atoi(r.FormValue("total"))
	request := transactiondto.TransRequest{
		CounterQty: counterqty,
		Total:      total,
		Status:     r.FormValue("status"),
		Attachment: r.FormValue("attachment"),
		TripId:     tripId,
		UserId:     userid,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := models.TransactionModels{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		TripId:     request.TripId,
		UserId:     request.UserId,
	}

	data, err := h.TransactionRepo.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrans(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTrans(u models.TransactionModels) transactiondto.TransResponse {
	return transactiondto.TransResponse{
		ID:         u.ID,
		CounterQty: u.ID,
		Total:      u.Total,
		Status:     u.Status,
		Attachment: u.Attachment,
		TripId:     u.TripId,
		UserId:     u.UserId,
	}
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kakobotasso/watermanager/models"
)

func (e Env) GetConsumption(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumptionType := vars["consumption_type"]
	serialNumber := vars["serial"]
	var consumptionList []models.Consumption

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var response interface{}

	if err := e.DB.Where(&models.Consumption{Serial: serialNumber}).Find(&consumptionList).Error; err != nil {
		response = models.Errors{
			models.Error{
				Key:     "not_found",
				Message: "Consumption not found",
			},
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if consumptionType == "liter" {
			response = &consumptionList
		} else if consumptionType == "money" {
			response = convertLitersToPrice(consumptionList)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			response = models.Errors{
				models.Error{
					Key:     "invalid_parameter",
					Message: "Invalid type of consumption",
				},
			}
		}
		json.NewEncoder(w).Encode(response)
	}
}

func (e Env) GetConsumptionMonthly(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumptionType := vars["consumption_type"]
	serialNumber := vars["serial"]
	var consumptionList []models.Consumption

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var response interface{}

	query := fmt.Sprintf("SELECT ID, Month, Year, SUM(Liter) as liter, Serial FROM consumptions WHERE Serial = '%s' GROUP BY Month, Year;", serialNumber)
	if err := e.DB.Raw(query).Scan(&consumptionList).Error; err != nil {
		response = models.Errors{
			models.Error{
				Key:     "not_found",
				Message: "Consumption not found",
			},
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if consumptionType == "liter" {
			response = &consumptionList
		} else if consumptionType == "money" {
			response = convertLitersToPrice(consumptionList)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			response = models.Errors{
				models.Error{
					Key:     "invalid_parameter",
					Message: "Invalid type of consumption",
				},
			}
		}
		json.NewEncoder(w).Encode(response)
	}
}

func (e Env) GetEstimatedConsumption(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumptionType := vars["consumption_type"]
	serialNumber := vars["serial"]
	var consumptionList []models.Consumption

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var response interface{}

	query := fmt.Sprintf("SELECT ID, Month, Year, AVG(Liter) as liter, Serial FROM consumptions WHERE Serial = '%s' ORDER BY id DESC LIMIT 30;", serialNumber)
	if err := e.DB.Raw(query).Scan(&consumptionList).Error; err != nil {
		response = models.Errors{
			models.Error{
				Key:     "not_found",
				Message: "Consumption not found",
			},
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if consumptionType == "liter" {
			response = &consumptionList
		} else if consumptionType == "money" {
			response = convertLitersToPrice(consumptionList)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			response = models.Errors{
				models.Error{
					Key:     "invalid_parameter",
					Message: "Invalid type of consumption",
				},
			}
		}
		json.NewEncoder(w).Encode(response)
	}
}

func (e Env) GetEstimatedConsumptionMonthly(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumptionType := vars["consumption_type"]
	serialNumber := vars["serial"]
	var estimated []models.Consumption
	var consumptionList []models.Consumption

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var response interface{}

	query := fmt.Sprintf("SELECT ID, Month, Year, AVG(Liter) as liter, Serial FROM consumptions WHERE Serial = '%s' ORDER BY id DESC LIMIT 30;", serialNumber)
	if err := e.DB.Raw(query).Scan(&estimated).Error; err != nil {
		response = models.Errors{
			models.Error{
				Key:     "not_found",
				Message: "Consumption not found",
			},
		}
	} else {
		query := fmt.Sprintf("SELECT ID, Month, Year, SUM(Liter) as liter, Serial FROM consumptions WHERE Serial = '%s' GROUP BY Month, Year;", serialNumber)
		if err := e.DB.Raw(query).Scan(&consumptionList).Error; err != nil {
			response = models.Errors{
				models.Error{
					Key:     "not_found",
					Message: "Consumption not found",
				},
			}
		} else {
			w.WriteHeader(http.StatusOK)
			if consumptionType == "liter" {
				response = models.ConsumptionApp{
					Estimated:       estimated,
					ConsumptionList: consumptionList,
				}
			} else if consumptionType == "money" {
				response = models.ConsumptionApp{
					Estimated:       convertLitersToPrice(estimated),
					ConsumptionList: convertLitersToPrice(consumptionList),
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				response = models.Errors{
					models.Error{
						Key:     "invalid_parameter",
						Message: "Invalid type of consumption",
					},
				}
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func (e Env) CreateConsumption(w http.ResponseWriter, r *http.Request) {
	var consumption models.Consumption
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &consumption); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	year, month, _ := time.Now().Date()

	consumption.Month = fmt.Sprintf("%d", month)
	consumption.Year = fmt.Sprintf("%d", year)

	e.DB.Create(&consumption)

	response := &consumption
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func convertLitersToPrice(consumptionList []models.Consumption) []models.Consumption {
	var returnList []models.Consumption
	for _, consumption := range consumptionList {
		oldLiter, _ := strconv.ParseFloat(consumption.Liter, 64)
		newConsumption := models.Consumption{
			Liter:  fmt.Sprintf("R$ %[1]f", (oldLiter * 0.005)),
			Month:  consumption.Month,
			Year:   consumption.Year,
			Serial: consumption.Serial,
		}
		returnList = append(returnList, newConsumption)
	}
	return returnList
}

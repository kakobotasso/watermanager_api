package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kakobotasso/watermanager_api/models"
)

func GetConsumption(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumptionType := vars["consumption_type"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var response interface{}

	if consumptionType == "liter" {
		w.WriteHeader(http.StatusOK)
		response = models.ConsumptionList{
			models.Consumption{
				Id:    1,
				Liter: "2",
			},
			models.Consumption{
				Id:    2,
				Liter: "2",
			},
			models.Consumption{
				Id:    3,
				Liter: "3",
			},
			models.Consumption{
				Id:    4,
				Liter: "5",
			},
			models.Consumption{
				Id:    5,
				Liter: "2",
			},
			models.Consumption{
				Id:    6,
				Liter: "1",
			},
			models.Consumption{
				Id:    7,
				Liter: "3",
			},
		}
	} else if consumptionType == "money" {
		w.WriteHeader(http.StatusOK)
		response = models.ConsumptionList{
			models.Consumption{
				Id:    1,
				Liter: "R$ 0,20",
			},
			models.Consumption{
				Id:    2,
				Liter: "R$ 0,20",
			},
			models.Consumption{
				Id:    3,
				Liter: "R$ 0,30",
			},
			models.Consumption{
				Id:    4,
				Liter: "R$ 0,50",
			},
			models.Consumption{
				Id:    5,
				Liter: "R$ 0,20",
			},
			models.Consumption{
				Id:    6,
				Liter: "R$ 0,10",
			},
			models.Consumption{
				Id:    7,
				Liter: "R$ 0,30",
			},
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

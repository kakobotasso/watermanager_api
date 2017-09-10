package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kakobotasso/watermanager/models"
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
				Month: "09/2017",
			},
			models.Consumption{
				Id:    2,
				Liter: "2",
				Month: "08/2017",
			},
			models.Consumption{
				Id:    3,
				Liter: "3",
				Month: "07/2017",
			},
			models.Consumption{
				Id:    4,
				Liter: "5",
				Month: "06/2017",
			},
			models.Consumption{
				Id:    5,
				Liter: "2",
				Month: "05/2017",
			},
			models.Consumption{
				Id:    6,
				Liter: "1",
				Month: "04/2017",
			},
			models.Consumption{
				Id:    7,
				Liter: "3",
				Month: "03/2017",
			},
		}
	} else if consumptionType == "money" {
		w.WriteHeader(http.StatusOK)
		response = models.ConsumptionList{
			models.Consumption{
				Id:    1,
				Liter: "R$ 0,20",
				Month: "09/2017",
			},
			models.Consumption{
				Id:    2,
				Liter: "R$ 0,20",
				Month: "08/2017",
			},
			models.Consumption{
				Id:    3,
				Liter: "R$ 0,30",
				Month: "07/2017",
			},
			models.Consumption{
				Id:    4,
				Liter: "R$ 0,50",
				Month: "06/2017",
			},
			models.Consumption{
				Id:    5,
				Liter: "R$ 0,20",
				Month: "05/2017",
			},
			models.Consumption{
				Id:    6,
				Liter: "R$ 0,10",
				Month: "04/2017",
			},
			models.Consumption{
				Id:    7,
				Liter: "R$ 0,30",
				Month: "03/2017",
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

package main

// https://www.youtube.com/watch?v=d_L64KT3SFM

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"os/exec"
)

type statusBateria struct {
	// GAG PORQUE NOS NOMES QUE COMECAM EM MINÚSCULOS SÃO PRIVAS E O MAIÚCULO É ABERTO SLK
	Porcentagem    string  `json:"porcentagem"`
	EstaCarregando bool    `json:"estaCarregando"`
	Temperatura    float64 `json:"temperatura"`
	SaudeBateria   string  `json:"saúde"`
}

type BatteryStatus struct {
	Health      string  `json:"health"`
	Percentage  int     `json:"percentage"`
	Plugged     string  `json:"plugged"`
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"`
}

var retorno = statusBateria{

	Porcentagem:    "50%",
	EstaCarregando: false,
	Temperatura:    30.2,
}

func getStatusBateria(context *gin.Context) {
	cmd := exec.Command("termux-battery-status")

	output, err := cmd.Output()
	if err != nil {
		var ErrMessage = "Erro ao executar comando: " + err.Error()
		context.IndentedJSON(http.StatusInternalServerError, ErrMessage)
		return
	}

	var battery BatteryStatus
	err = json.Unmarshal(output, &battery)
	if err != nil {
		var ErrMessage = "Erro ao executar comando: " + err.Error()
		context.IndentedJSON(http.StatusInternalServerError, ErrMessage)
		return
	}
	retorno.EstaCarregando = battery.Plugged == "PLUGGED"
	retorno.Porcentagem = strconv.Itoa(battery.Percentage) + "%"
	retorno.Temperatura = battery.Temperature
	retorno.SaudeBateria = battery.Health
	
	context.IndentedJSON(http.StatusOK, retorno)
}

func main() {
	router := gin.Default()
	router.GET("/bateria", getStatusBateria)
	router.Run("192.168.15.230:9090")
}

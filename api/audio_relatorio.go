package api

import (
	"time"

	"github.com/totalvoice/go-client/api/model"
)

// AudioRelatorioService service
type AudioRelatorioService struct {
	client  HTTPClient
	handler Response
}

// Gerar - Relatório de mensagens de Audio
func (s AudioRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*model.AudioRelatorioResponse, error) {

	relatorio := new(model.AudioRelatorio)
	params := map[string]string{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	response := new(model.AudioRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaAudio+"/relatorio", params)
	if err != nil {
		return response, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.AudioRelatorioResponse), err
}

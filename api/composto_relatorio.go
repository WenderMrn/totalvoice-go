package api

import (
	"time"

	"github.com/totalvoice/go-client/api/model"
)

// CompostoRelatorioService service
type CompostoRelatorioService struct {
	client  HTTPClient
	handler Response
}

// Gerar - Relatório de mensagens Compostas
func (s CompostoRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*model.CompostoRelatorioResponse, error) {

	relatorio := new(model.CompostoRelatorio)
	params := map[string]string{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	response := new(model.CompostoRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaComposto+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.CompostoRelatorioResponse), err
}

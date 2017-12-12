package api

import "github.com/totalvoice/totalvoice-go/api/model"

// ContaRelatorioService service
type ContaRelatorioService struct {
	client  HTTPClient
	handler Response
}

// Gerar - Lista contas criadas por mim
func (s ContaRelatorioService) Gerar() (*model.ContaRelatorioResponse, error) {

	relatorio := new(model.ContaRelatorio)
	response := new(model.ContaRelatorioResponse)

	http, err := s.client.ListResource(relatorio, RotaConta+"/relatorio", nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.ContaRelatorioResponse), err
}

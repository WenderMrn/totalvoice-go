package totalvoice

import (
	"github.com/totalvoice/totalvoice-go/api"
)

// BaseURI - URI Base
const (
	BaseURI = "https://api2.totalvoice.com.br"
)

// TotalVoice struct
type TotalVoice struct {
	client api.HTTPClient

	Perfil      *api.PerfilService
	Audio       *api.AudioService
	Webhook     *api.WebhookService
	Saldo       *api.SaldoService
	Conta       *api.ContaService
	Composto    *api.CompostoService
	Chamada     *api.ChamadaService
	Conferencia *api.ConferenciaService
	SMS         *api.SMSService
	TTS         *api.TTSService
	Ramal       *api.RamalService
	URA         *api.URAService
	DID         *api.DIDService
}

// NewTotalVoiceClient - Cria TotalVoice struct.
func NewTotalVoiceClient(accessToken string) *TotalVoice {

	client := NewClient(accessToken, BaseURI)
	return client
}

// NewClient - Cria TotalVoice struct.
func NewClient(accessToken string, baseURI string) *TotalVoice {

	client := &Client{accessToken: accessToken, baseURI: baseURI}
	tvce := &TotalVoice{client: client}

	handler := api.Response{}

	tvce.Perfil = api.NewPerfilService(client, handler)
	tvce.Audio = api.NewAudioService(client, handler)
	tvce.Webhook = api.NewWebhookService(client, handler)
	tvce.Saldo = api.NewSaldoService(client, handler)
	tvce.Conta = api.NewContaService(client, handler)
	tvce.Composto = api.NewCompostoService(client, handler)
	tvce.Chamada = api.NewChamadaService(client, handler)
	tvce.Conferencia = api.NewConferenciaService(client, handler)
	tvce.SMS = api.NewSMSService(client, handler)
	tvce.TTS = api.NewTTSService(client, handler)
	tvce.Ramal = api.NewRamalService(client, handler)
	tvce.URA = api.NewURAService(client, handler)
	tvce.DID = api.NewDIDService(client, handler)

	return tvce
}

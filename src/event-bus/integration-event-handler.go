package eventbus

type IntegrationEventHandler interface {
	Handle(event struct{})
}

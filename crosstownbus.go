package crosstownbus

import (
	"reflect"

	eventbus "github.com/paoloposso/crosstownbus/event_bus"
	rabbitmqbus "github.com/paoloposso/crosstownbus/rabbitmq_bus"
	redisbus "github.com/paoloposso/crosstownbus/redis_bus"
)

func CreateRedisBus(event reflect.Type, uri string, password string) (eventbus.Bus, error) {
	return redisbus.CreateBus(event.Name(),
		redisbus.RedisConfig{
			Uri:      uri,
			Password: password,
		})
}

func CreateRabbitMQBus(event reflect.Type, uri string) (eventbus.Bus, error) {
	return rabbitmqbus.CreateBus(event.Name(),
		rabbitmqbus.RabbitMQOptions{
			Uri: uri,
		})
}
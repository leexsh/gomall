package emailConsume

import (
	"encoding/json"
	"gomall/app/email/infra/mq"
	"gomall/app/email/infra/notify"
	"gomall/app/email/kitex_gen/frontend/email"

	"github.com/cloudwego/kitex/server"
)

func ConsumerInit()  {
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := json.Unmarshal(msg.Data, req)
		if err != nil {
			return
		}	
		noopEamil := notify.NewNoopEmail()
		noopEamil.Send(&req)
	})
	if err != nil {
		panic(err)
	}
	server.RegisterShutdownHook(func ()  {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
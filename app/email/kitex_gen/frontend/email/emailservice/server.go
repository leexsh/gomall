// Code generated by Kitex v0.9.0. DO NOT EDIT.
package emailservice

import (
	server "github.com/cloudwego/kitex/server"
	email "gomall/app/email/kitex_gen/frontend/email"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler email.EmailService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler email.EmailService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
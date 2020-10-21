package services

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(ServicesGroup{})

type ServicesGroup struct {
	courier.EmptyOperator
}

func (ServicesGroup) Path() string {
	return "/services"
}

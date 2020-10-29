package main

import (
	"errors"
	"fmt"
)

import (
	"github.com/transaction-wg/seata-golang/pkg/context"
	"github.com/transaction-wg/seata-golang/pkg/tcc"
)

type ServiceB struct {
}

func (svc *ServiceB) Try(ctx *context.BusinessActionContext) (bool, error) {
	word := ctx.ActionContext["hello"]
	fmt.Println(word)
	fmt.Println("Service B Tried!")
	return false, errors.New("service B try Failed")
}

func (svc *ServiceB) Confirm(ctx *context.BusinessActionContext) bool {
	word := ctx.ActionContext["hello"]
	fmt.Println(word)
	fmt.Println("Service B confirmed!")
	return true
}

func (svc *ServiceB) Cancel(ctx *context.BusinessActionContext) bool {
	word := ctx.ActionContext["hello"]
	fmt.Println(word)
	fmt.Println("Service B canceled!")
	return true
}

var serviceB = &ServiceB{}

type TCCProxyServiceB struct {
	*ServiceB

	Try func(ctx *context.BusinessActionContext) (bool, error) `TccActionName:"ServiceB"`
}

func (svc *TCCProxyServiceB) GetTccService() tcc.TccService {
	return svc.ServiceB
}

var TccProxyServiceB = &TCCProxyServiceB{
	ServiceB: serviceB,
}

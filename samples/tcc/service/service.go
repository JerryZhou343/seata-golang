package main

import (
	"context"
	"log"
)

import (
	context2 "github.com/transaction-wg/seata-golang/pkg/context"
	"github.com/transaction-wg/seata-golang/pkg/tm"
)

type Service struct {
}

func (svc *Service) TCCCommitted(ctx context.Context) error {
	log.Print("xxxxxxxxxxxxxxxxxxxxxxx")
	rootContext := ctx.(*context2.RootContext)
	businessActionContextA := &context2.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	// 业务参数全部放到 ActionContext 里
	businessActionContextA.ActionContext["hello"] = "hello world,this is from BusinessActionContext A"

	businessActionContextB := &context2.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextB.ActionContext["hello"] = "hello world,this is from BusinessActionContext B"
	businessActionContextA.GetXID()
	/*
	resultA, err := serviceA.TccProxyServiceA.Try(businessActionContextA)
	fmt.Printf("result A is :%v", resultA)
	if err != nil {
		return err
	}

	resultB, err := serviceB.TccProxyServiceB.Try(businessActionContextB)
	fmt.Printf("result B is :%v", resultB)
	if err != nil {
		return err
	}
	*/
	return nil
}

func (svc *Service) TCCCanceled(ctx context.Context) error {
	log.Println("TCCCanceled....................")
	rootContext := ctx.(*context2.RootContext)
	businessActionContextA := &context2.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextA.ActionContext["hello"] = "hello world,this is from BusinessActionContext A"

	businessActionContextC := &context2.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextC.ActionContext["hello"] = "hello world,this is from BusinessActionContext C"
	/*
	resultA, err := serviceA.TccProxyServiceA.Try(businessActionContextA)
	fmt.Printf("result A is :%v", resultA)
	if err != nil {
		return err
	}

	resultC, err := serviceC.TccProxyServiceC.Try(businessActionContextC)
	fmt.Printf("result C is :%v", resultC)
	if err != nil {
		return err
	}


	 */
	return nil
}

var service = &Service{}

type ProxyService struct {
	*Service
	TCCCommitted func(ctx context.Context) error
	TCCCanceled  func(ctx context.Context) error
}

func (svc *ProxyService) GetProxyService() interface{} {
	return svc.Service
}

func (svc *ProxyService) GetMethodTransactionInfo(methodName string) *tm.TransactionInfo {
	return methodTransactionInfo[methodName]
}

var methodTransactionInfo = make(map[string]*tm.TransactionInfo)

func init() {
	methodTransactionInfo["TCCCommitted"] = &tm.TransactionInfo{
		TimeOut:     60000000,
		Name:        "TCC_TEST_COMMITTED",
		Propagation: tm.REQUIRED,
	}
	methodTransactionInfo["TCCCanceled"] = &tm.TransactionInfo{
		TimeOut:     60000000,
		Name:        "TCC_TEST_CANCELED",
		Propagation: tm.REQUIRED,
	}
}

var ProxySvc = &ProxyService{
	Service:      service,
}

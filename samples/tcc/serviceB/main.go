package main

import (
	"context"
	"fmt"
	getty "github.com/apache/dubbo-getty"
	"github.com/gin-gonic/gin"
	"github.com/transaction-wg/seata-golang/pkg"
	"github.com/transaction-wg/seata-golang/pkg/config"
	bc "github.com/transaction-wg/seata-golang/pkg/context"
	"log"
	"os"
)

import (
	"github.com/transaction-wg/seata-golang/pkg/tcc"
)

type logger struct{}

func (l logger) Info(args ...interface{}) {
	//panic("implement me")

	return
}

func (l logger) Warn(args ...interface{}) {
	//panic("implement me")

	return
}

func (l logger) Error(args ...interface{}) {
	//panic("implement me")

	return
}

func (l logger) Debug(args ...interface{}) {
	//panic("implement me")

	return
}

func (l logger) Infof(fmt string, args ...interface{}) {
	//panic("implement me")

	return
}

func (l logger) Warnf(fmt string, args ...interface{}) {
	//panic("implement me")

	return
}

func (l logger) Errorf(fmt string, args ...interface{}) {
	//panic("implement me")

	return
}

func (l logger) Debugf(fmt string, args ...interface{}) {
	//panic("implement me")
	return
}

func main() {
	r := gin.Default()
	getty.SetLogger(&logger{})
	//config.InitConfWithDefault("testService")
	config.InitConf(os.Args[1])
	pkg.NewRpcClient()
	tcc.InitTCCResourceManager()

	//tm.Implement(service.ProxySvc)
	//tcc.ImplementTCC(serviceA.TccProxyServiceA)
	tcc.ImplementTCC(TccProxyServiceB)
	//tcc.ImplementTCC(serviceC.TccProxyServiceC)

	r.GET("/try", func(c *gin.Context) {
		//service.ProxySvc.TCCCommitted(c)
		xid, find := c.GetQuery("xid")
		if find == false {
			log.Print("not found xid")
			return
		}
		log.Printf("xid:%s", xid)
		ctx := context.WithValue(context.TODO(), bc.KEY_XID, xid)
		rctx := bc.NewRootContext(ctx)
		businessActionContextB := &bc.BusinessActionContext{
			RootContext:   rctx,
			ActionContext: make(map[string]interface{}),
		}
		TccProxyServiceB.Try(businessActionContextB)
		c.JSON(200, gin.H{
			"message": "pong",
		})
		//os.Exit(-1)
	})
	r.GET("/rollback", func(c *gin.Context) {
		//service.ProxySvc.TCCCanceled(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%s", os.Args[2]))
}

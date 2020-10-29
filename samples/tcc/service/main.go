package main

import (
	"context"
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
	"github.com/transaction-wg/seata-golang/pkg/tm"
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

var _ getty.Logger = &logger{}

func main() {
	r := gin.Default()
	//getty.SetLoggerLevel(getty.LoggerLevelInfo)
	getty.SetLogger(&logger{})
	//config.InitConfWithDefault("testService")
	config.InitConf(os.Args[1])
	pkg.NewRpcClient()
	tcc.InitTCCResourceManager()

	tm.Implement(ProxySvc)
	//tcc.ImplementTCC(serviceA.TccProxyServiceA)
	//tcc.ImplementTCC(serviceB.TccProxyServiceB)
	//tcc.ImplementTCC(serviceC.TccProxyServiceC)

	r.GET("/commit", func(c *gin.Context) {
		log.Print(".................")
		if ProxySvc.TCCCommitted != nil {
			ProxySvc.TCCCommitted(c)
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/rollback", func(c *gin.Context) {
		xid, find := c.GetQuery("xid")
		if find == false {
			log.Print("not found xid")
			return
		}
		log.Printf("xid:%s", xid)
		ctx := context.WithValue(context.TODO(), bc.KEY_XID, xid)
		rctx := bc.NewRootContext(ctx)
		businessActionContext := &bc.BusinessActionContext{
			RootContext:   rctx,
			ActionContext: make(map[string]interface{}),
		}
		ProxySvc.TCCCanceled(businessActionContext)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":80")
}

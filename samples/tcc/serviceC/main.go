package main

import (
	"context"
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

func main() {
	r := gin.Default()

	//config.InitConfWithDefault("testService")
	config.InitConf(os.Args[1])
	pkg.NewRpcClient()
	tcc.InitTCCResourceManager()

	//tm.Implement(service.ProxySvc)
	//tcc.ImplementTCC(serviceA.TccProxyServiceA)
	//tcc.ImplementTCC(serviceB.TccProxyServiceB)
	tcc.ImplementTCC(TccProxyServiceC)

	r.GET("/try", func(c *gin.Context) {
		xid, find := c.GetQuery("xid")
		if find == false {
			log.Print("not found xid")
			return
		}
		log.Printf("xid:%s", xid)
		ctx := context.WithValue(context.TODO(), bc.KEY_XID, xid)
		rctx := bc.NewRootContext(ctx)
		businessActionContextC := &bc.BusinessActionContext{
			RootContext:   rctx,
			ActionContext: make(map[string]interface{}),
		}
		TccProxyServiceC.Try(businessActionContextC)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/rollback", func(c *gin.Context) {
		//service.ProxySvc.TCCCanceled(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":83")
}

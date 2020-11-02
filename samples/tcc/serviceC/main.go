package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/transaction-wg/seata-golang/pkg"
	"github.com/transaction-wg/seata-golang/pkg/config"
	bc "github.com/transaction-wg/seata-golang/pkg/context"
	"log"
	"net/http"
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
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": "pong",
		})
		os.Exit(-1)
	})
	r.GET("/rollback", func(c *gin.Context) {
		//service.ProxySvc.TCCCanceled(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%s", os.Args[2]))
}

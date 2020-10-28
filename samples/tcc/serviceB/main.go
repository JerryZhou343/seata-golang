package main

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-wg/seata-golang/pkg"
	"github.com/transaction-wg/seata-golang/pkg/config"
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
	tcc.ImplementTCC(TccProxyServiceB)
	//tcc.ImplementTCC(serviceC.TccProxyServiceC)

	r.GET("/commit", func(c *gin.Context) {
		//service.ProxySvc.TCCCommitted(c)
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
	r.Run(":82")
}

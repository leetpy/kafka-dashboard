package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetpy/kafka-dashboard/config"
	"github.com/leetpy/kafka-dashboard/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))
	r := gin.New()
	router.Load(r)
	zap.L().Info(http.ListenAndServe(viper.GetString("addr"), r).Error())
}

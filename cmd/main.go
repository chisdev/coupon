package main

import (
	"github.com/ChisTrun/carbon/pkg/config"
	_ "github.com/go-sql-driver/mysql"

	"github.com/chisdev/coupon/internal/server"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}

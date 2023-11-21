package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"oceanlearn.teach/ginessential/common"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}

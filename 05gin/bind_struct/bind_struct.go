package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a"`
}
type StructB struct{
	NestedStruct StructA
	FieldB string `form:"field_b"`
}
type StructC struct{
	NestedStructPointer *StructA 
	FieldC string `form:"field_c"`
}
type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}
func GetDataB(ctx *gin.Context) {
	var b StructB
	ctx.Bind(&b)
	ctx.JSON(http.StatusOK, gin.H{
		"a":b.NestedStruct,
		"c":b.FieldB,
	})
}
func GetDataC(ctx *gin.Context) {
	var c StructC
	ctx.Bind(&c)
	ctx.JSON(http.StatusOK, gin.H{
		"a": c.NestedStructPointer,
		"c": c.FieldC,
	})
}
func GetDataD(ctx *gin.Context) {
	var d StructD
	ctx.Bind(&d)
	ctx.JSON(http.StatusOK, gin.H{
		"x": d.NestedAnonyStruct,
		"c": d.FieldD,
	})
}
func main() {
	router:=gin.Default()
	//curl -X GET "http://localhost:8080/getb?field_a=hello&field_b=world"
	//curl "http://localhost:8080/getb?field_a=hello&field_b=world"
	router.GET("/getb",GetDataB)
	//curl -X GET "http://localhost:8080/getc?field_a=hello&field_c=world"
	//curl "http://localhost:8080/getc?field_a=hello&field_c=world"
	router.GET("/getc", GetDataC)
	//curl -X GET "http://localhost:8080/getd?field_x=hello&field_d=world"
	//curl "http://localhost:8080/getd?field_x=hello&field_d=world"
	router.GET("/getd", GetDataD)
	router.Run(":8080")
}
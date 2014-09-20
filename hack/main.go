package main

import (
	"fmt"
	"github.com/nimishzynga/goweb"
	"net/http"
)

func setupHandler() {
	goweb.MapFunc("/queryDoc", func(c *goweb.Context) {
		HandleDoc(c)
	})
}

func runExpr(doc []byte, expr []string) string {
	//exprs := []string{`{"type":"property","path":"name"}`, `{"type":"property","path":"abv"}`}
	cExpr, err := CompileN1QLExpression(expr)
	if err != nil {
	}
	secKey, err := N1QLTransform(doc, cExpr)
	if err != nil {
	}
	return string(secKey)
}

type Doc struct {
	//DocData map[string]string
	DocData string
	DocExpr []string
}

func HandleDoc(c *goweb.Context) {
	fmt.Println("dfdfd")
	if c.IsPost() || c.IsPut() {
		//NewDoc := Doc{DocData: make(map[string]string)}
		NewDoc := Doc{}
		if err := c.Fill(&NewDoc); err != nil {
			return
		}
		fmt.Println("query doc = ", NewDoc.DocData, "expr = ", NewDoc.DocExpr)
		data := runExpr([]byte(NewDoc.DocData), NewDoc.DocExpr)
		fmt.Println("map result = ", data)
		c.WriteResponse(data, 200)
	}
}

func main() {
	goweb.ConfigureDefaultFormatters()
	setupHandler()
	http.ListenAndServe(":10000", goweb.DefaultHttpHandler)
}

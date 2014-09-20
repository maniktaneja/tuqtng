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

func runWhere(doc []byte, expr string) string {
	//exprs := []string{`{"type":"property","path":"name"}`, `{"type":"property","path":"abv"}`}
	cExpr := HackCompile(expr)
    fmt.Println("compiled expr is ", cExpr)
	secKey, err := N1QLTransform(doc, cExpr)
	if err != nil {
        fmt.Println("transform failed")
	}
	return string(secKey)
}

func runExpr(doc []byte, expr []string) string {
	//exprs := []string{`{"type":"property","path":"name"}`, `{"type":"property","path":"abv"}`}
	cExpr,err := CompileN1QLExpression(expr)
	if err != nil {
        fmt.Println("compilation failed")
	}
    fmt.Println("compiled expr is ", cExpr)
	secKey, err := N1QLTransform(doc, cExpr)
	if err != nil {
        fmt.Println("transform failed", err)
	}
	return string(secKey)
}

type Expr struct {
    Where string
    Emit  []string
}

type Doc struct {
	//DocData map[string]string
	DocData string
	DocExpr Expr
}

func HandleDoc(c *goweb.Context) {
	fmt.Println("dfdfd")
	if c.IsPost() || c.IsPut() {
		//NewDoc := Doc{DocData: make(map[string]string)}
		NewDoc := Doc{}
		if err := c.Fill(&NewDoc); err != nil {
            fmt.Println(err)
			return
		}
		fmt.Println("query doc = ", NewDoc.DocData, "expr = ", NewDoc.DocExpr)
		data := runWhere([]byte(NewDoc.DocData), NewDoc.DocExpr.Where)
        if data != "" {
		    data = runExpr([]byte(NewDoc.DocData), NewDoc.DocExpr.Emit)
        } else {
            fmt.Println("runwhere failed")
        }
		fmt.Println("map result = ", data)
		c.WriteResponse(data, 200)
	}
}

func main() {
	goweb.ConfigureDefaultFormatters()
	setupHandler()
	http.ListenAndServe(":10000", goweb.DefaultHttpHandler)
}

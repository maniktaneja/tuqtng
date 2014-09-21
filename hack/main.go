package main

import (
	"encoding/json"
	"fmt"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	yaccParser "github.com/couchbaselabs/tuqtng/parser/goyacc"
	"github.com/nimishzynga/goweb"
	"net/http"
)

func setupHandler() {
	goweb.MapFunc("/queryDoc", func(c *goweb.Context) {
		HandleDoc(c)
	})
}

func runWhere(doc []byte, expr string) interface{} {
	//exprs := []string{`{"type":"property","path":"name"}`, `{"type":"property","path":"abv"}`}
	parser := yaccParser.NewN1qlParser()
	random_ast, _ := parser.Parse(expr)
	select_statement := random_ast.(*ast.SelectStatement)
	jsonm, _ := json.Marshal(select_statement.Where)
	fmt.Printf("Where clause %v\n", string(jsonm))

	a := dparval.NewValueFromBytes(doc)
	result, err := select_statement.Where.Evaluate(a)
	if err != nil {
		fmt.Println("transform failed")
	}
	return result.Value()
}

func runExpr(doc []byte, expr string) string {
	//exprs := []string{`{"type":"property","path":"name"}`, `{"type":"property","path":"abv"}`}
	cExpr, err := CompileN1QLExpression([]string{expr})
	if err != nil {
		fmt.Println("compilation failed")
	}
	fmt.Println("compiled emit expr is ", cExpr)
	secKey, err := N1QLTransform(doc, cExpr)
	if err != nil {
		fmt.Println("transform failed", err)
	}
	return string(secKey)
}

type Expr struct {
	Where string `json:"where"`
	Emit  string `json:"emit"`
}

type Doc struct {
	//DocData map[string]string
	DocData string
	DocExpr Expr
}

func HandleDoc(c *goweb.Context) {
	fmt.Println("handling")
	if c.IsPost() || c.IsPut() {
		//NewDoc := Doc{DocData: make(map[string]string)}
		NewDoc := Doc{}
		if err := c.Fill(&NewDoc); err != nil {
			fmt.Println("error is here", err)
			return
		}
		fmt.Println("query doc = ", NewDoc.DocData, "expr = ", NewDoc.DocExpr)
		where_clause := runWhere([]byte(NewDoc.DocData), NewDoc.DocExpr.Where)
		fmt.Println("where returned ", where_clause)
		emit := ""
		if where_clause == true {
			emit = runExpr([]byte(NewDoc.DocData), NewDoc.DocExpr.Emit)
		} else {
			fmt.Println("runwhere failed")
		}
		fmt.Println("map result = ", emit)
		c.WriteResponse(emit, 200)
	}
}

func main() {
	goweb.ConfigureDefaultFormatters()
	setupHandler()
	http.ListenAndServe(":10000", goweb.DefaultHttpHandler)
}

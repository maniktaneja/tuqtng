//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

// Plan is a description of a sequence of steps to produce a correct
// result for a query.

package plan

import (
	"github.com/couchbaselabs/tuqtng/ast"
)

type Plan struct {
	Root PlanElement `json:"root"`
}

type PlanChannel chan Plan

type PlanElement interface {
}

type ExpressionEvaluator struct {
	Type string `json:"type"`
}

func NewExpressionEvaluator() *ExpressionEvaluator {
	return &ExpressionEvaluator{
		Type: "expr_evaluator",
	}
}

type Scan struct {
	Type string `json:"type"`
}

func NewScan() *Scan {
	return &Scan{
		Type: "scan",
	}
}

type Fetch struct {
	Type  string      `json:"type"`
	Input PlanElement `json:"input"`
}

func NewFetch(input PlanElement) *Fetch {
	return &Fetch{
		Type:  "fetch",
		Input: input,
	}
}

type Filter struct {
	Type  string         `json:"type"`
	Input PlanElement    `json:"input"`
	Expr  ast.Expression `json:"expr"`
}

func NewFilter(input PlanElement, expr ast.Expression) *Filter {
	return &Filter{
		Type:  "filter",
		Input: input,
		Expr:  expr,
	}
}

type Order struct {
	Type  string                `json:"type"`
	Input PlanElement           `json:"input"`
	Sort  []*ast.SortExpression `json:"sort"`
}

func NewOrder(input PlanElement, sort []*ast.SortExpression) *Order {
	return &Order{
		Type:  "order",
		Input: input,
		Sort:  sort,
	}
}

type Limit struct {
	Type  string      `json:"type"`
	Input PlanElement `json:"input"`
	Val   int         `json:"value"`
}

func NewLimit(input PlanElement, limit int) *Limit {
	return &Limit{
		Type:  "limit",
		Input: input,
		Val:   limit,
	}
}

type Offset struct {
	Type  string      `json:"type"`
	Input PlanElement `json:"input"`
	Val   int         `json:"value"`
}

func NewOffset(input PlanElement, offset int) *Offset {
	return &Offset{
		Type:  "offset",
		Input: input,
		Val:   offset,
	}
}

type Projector struct {
	Type   string                   `json:"projector"`
	Input  PlanElement              `json:"input"`
	Result ast.ResultExpressionList `json:"result"`
}

func NewProjector(input PlanElement, result ast.ResultExpressionList) *Projector {
	return &Projector{
		Type:   "projector",
		Input:  input,
		Result: result,
	}
}

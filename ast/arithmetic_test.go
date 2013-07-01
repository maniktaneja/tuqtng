//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import (
	"reflect"
	"testing"
)

func TestArithmetic(t *testing.T) {

	numberSix := NewLiteralNumber(6.0)
	numberSeven := NewLiteralNumber(7.0)
	numberNegativeSeven := NewLiteralNumber(-7.0)
	stringCouchbase := NewLiteralString("Couchbase")
	stringServer := NewLiteralString("Server")
	dneProperty := NewProperty("foo")

	tests := []struct {
		input  Expression
		output interface{}
		err    error
	}{
		{NewPlusOperator(stringCouchbase, stringServer), nil, nil}, // no longer support string concatenation, uses different operator

		{NewPlusOperator(numberSeven, numberSeven), 14.0, nil},
		{NewPlusOperator(numberSeven, stringCouchbase), nil, nil},
		{NewPlusOperator(stringCouchbase, numberSeven), nil, nil},
		{NewPlusOperator(dneProperty, numberSeven), nil, &Undefined{"foo"}},
		{NewPlusOperator(numberSeven, dneProperty), nil, &Undefined{"foo"}},

		{NewSubtractOperator(numberSeven, numberSeven), 0.0, nil},
		{NewSubtractOperator(numberSeven, stringCouchbase), nil, nil},
		{NewSubtractOperator(stringCouchbase, numberSeven), nil, nil},
		{NewSubtractOperator(dneProperty, numberSeven), nil, &Undefined{"foo"}},
		{NewSubtractOperator(numberSeven, dneProperty), nil, &Undefined{"foo"}},

		{NewMultiplyOperator(numberSeven, numberSeven), 49.0, nil},
		{NewMultiplyOperator(numberSeven, stringCouchbase), nil, nil},
		{NewMultiplyOperator(stringCouchbase, numberSeven), nil, nil},
		{NewMultiplyOperator(dneProperty, numberSeven), nil, &Undefined{"foo"}},
		{NewMultiplyOperator(numberSeven, dneProperty), nil, &Undefined{"foo"}},

		{NewDivideOperator(numberSeven, numberSeven), 1.0, nil},
		{NewDivideOperator(numberSeven, stringCouchbase), nil, nil},
		{NewDivideOperator(stringCouchbase, numberSeven), nil, nil},
		{NewDivideOperator(dneProperty, numberSeven), nil, &Undefined{"foo"}},
		{NewDivideOperator(numberSeven, dneProperty), nil, &Undefined{"foo"}},

		{NewModuloOperator(numberSeven, numberSix), 1.0, nil},
		{NewModuloOperator(stringCouchbase, numberSix), nil, nil},
		{NewModuloOperator(numberSeven, stringCouchbase), nil, nil},
		{NewModuloOperator(dneProperty, numberSix), nil, &Undefined{"foo"}},
		{NewModuloOperator(numberSeven, dneProperty), nil, &Undefined{"foo"}},

		{NewChangeSignOperator(numberSeven), -7.0, nil},
		{NewChangeSignOperator(numberNegativeSeven), 7.0, nil},
		{NewChangeSignOperator(stringCouchbase), nil, nil},
		{NewChangeSignOperator(dneProperty), nil, &Undefined{"foo"}},
	}

	for _, x := range tests {
		result, err := x.input.Evaluate(nil)
		if !reflect.DeepEqual(err, x.err) {
			t.Fatalf("Expected error: %v, got %v", x.err, err)
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %t %v, got %t %v", x.output, x.output, result, result)
		}
	}

}
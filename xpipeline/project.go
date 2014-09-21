//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

import (
	"fmt"
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
)

const DEBUG_PROJECT_CHANNEL = "OP_PROJECT"

type Project struct {
	Base         *BaseOperator
	Result       ast.ResultExpressionList
	projectEmpty bool
}

func NewProject(result ast.ResultExpressionList, projectEmpty bool) *Project {
	return &Project{
		Base:         NewBaseOperator(),
		Result:       result,
		projectEmpty: projectEmpty,
	}
}

func (this *Project) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *Project) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *Project) Run(stopChannel misc.StopChannel) {
	clog.To(CHANNEL, "project operator starting")
	this.Base.RunOperator(this, stopChannel)
	clog.To(CHANNEL, "project operator finished")
}

func (this *Project) processItem(item *dparval.Value) bool {
	resultMap := map[string]interface{}{}
	for _, resultItem := range this.Result {

		val, err := this.Base.projectedValueOfResultExpression(item, resultItem)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// undefined contributes nothing to the result map
				fmt.Printf(" error %v", err)
				val = item
				//continue
			default:
				return this.Base.SendError(query.NewError(err, "unexpected error projecting expression"))
			}
		}

		if resultItem.Star {
			if val != nil {
				valval := val.Value()
				switch valval := valval.(type) {
				case map[string]interface{}:
					// then if the result was an object
					// add its contents ot the result map
					for k, v := range valval {
						resultMap[k] = v
					}
				default:
					fmt.Printf(":((")
				}
			}
		} else {
			resultMap[resultItem.As] = val
		}
	}

	if !this.projectEmpty && len(resultMap) == 0 {
		return true
	}

	// build a Value from the projection
	projection := dparval.NewValue(resultMap)

	// store the projection as an attachment on the main item
	item.SetAttachment("projection", projection)

	clog.To(DEBUG_PROJECT_CHANNEL, "projecting: %v", item)

	// write to the output
	return this.Base.SendItem(item)
}

func (this *Project) afterItems() {}

func (this *Project) SetQuery(q network.Query) {
	this.Base.SetQuery(q)
}

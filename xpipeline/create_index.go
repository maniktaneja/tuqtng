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
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
)

type CreateIndex struct {
	itemChannel           dparval.ValueChannel
	supportChannel        PipelineSupportChannel
	bucket                catalog.Bucket
	name                  string
	index_type            string
	primary               bool
	on                    ast.ExpressionList
	selectStatement       *ast.SelectStatement
	downstreamStopChannel misc.StopChannel
	query                 network.Query
}

func NewCreateIndex(bucket catalog.Bucket, name string, index_type string, primary bool, on ast.ExpressionList, selectStatement *ast.SelectStatement) *CreateIndex {
	return &CreateIndex{
		itemChannel:     make(dparval.ValueChannel),
		supportChannel:  make(PipelineSupportChannel),
		bucket:          bucket,
		name:            name,
		index_type:      index_type,
		primary:         primary,
		on:              on,
		selectStatement: selectStatement,
	}
}

func (this *CreateIndex) SetSource(source Operator) {}

func (this *CreateIndex) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *CreateIndex) Run(stopChannel misc.StopChannel) {
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	// this MUST be here so that it runs before the channels are closed
	defer this.RecoverPanic()

	indexType := catalog.IndexType(strings.ToLower(this.index_type))
	indexOn := make(catalog.IndexKey, len(this.on))
	for pos, key := range this.on {
		indexOn[pos] = key
	}

	this.downstreamStopChannel = stopChannel

	var index catalog.Index
	var err query.Error

	// select expression processing
	if this.selectStatement != nil {
		projection, _ := json.Marshal(this.selectStatement.Select)
		clog.To(CHANNEL, "Result Expression list %s ", projection)
		where := this.selectStatement.Where.String()
		where = "select * from hack where " + where
		clog.To(CHANNEL, "Where expression %s", where)
		group, _ := json.Marshal(this.selectStatement.GroupBy)
		clog.To(CHANNEL, "Group by expression %s ", group)

		clog.To(CHANNEL, "create_index (mr index) operator starting")
		index, err = this.bucket.CreateMRIndex(this.name, indexOn, string(projection), string(where), string(group))

	} else if this.primary {
		clog.To(CHANNEL, "create_index (primary) operator starting")
		index, err = this.bucket.CreatePrimaryIndex()
	} else {
		clog.To(CHANNEL, "create_index (secondary) operator starting")
		index, err = this.bucket.CreateIndex(this.name, indexOn, indexType)
	}

	if err != nil {
		this.SendError(err)
	} else {
		if index != nil {
			item := dparval.NewValue(map[string]interface{}{})
			item.SetAttachment("projection", map[string]interface{}{
				"id":   index.Id(),
				"name": index.Name(),
			})
			this.SendItem(item)
		} else {
			clog.Warn("Successfully created index, but index was nil")
		}
	}
	clog.To(CHANNEL, "create_index operator finished")
}

func (this *CreateIndex) processItem(item *dparval.Value) bool {
	return true
}

func (this *CreateIndex) afterItems() {}

func (this *CreateIndex) SendItem(item *dparval.Value) bool {
	ok := true
	for ok {
		select {
		case this.itemChannel <- item:
			return true
		case _, ok = <-this.downstreamStopChannel:
			// someone closed the stop channel
		}
	}
	return ok
}

func (this *CreateIndex) SendError(err query.Error) bool {
	ok := true
	for ok {
		select {
		case this.supportChannel <- err:
			if err.IsFatal() {
				return false
			}
			return true
		case _, ok = <-this.downstreamStopChannel:
			// someone closed the stop channel
		}
	}
	return false
}

func (this *CreateIndex) RecoverPanic() {
	r := recover()
	if r != nil {
		clog.Error(fmt.Errorf("Query Execution Panic: %v\n%s", r, debug.Stack()))
		this.SendError(query.NewError(nil, "Panic In Exeuction Pipeline"))
	}
}

func (this *CreateIndex) SetQuery(q network.Query) {
	this.query = q
}

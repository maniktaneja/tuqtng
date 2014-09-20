//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import ()

type CreateIndexStatement struct {
	Method      string           `json:"method"`
	Name        string           `json:"name"`
	ExplainOnly bool             `json:"explain"`
	Bucket      string           `json:"bucket"`
	Pool        string           `json:"pool"`
	On          ExpressionList   `json:"on"`
	Select      *SelectStatement `json: "select statement"`
	Primary     bool             `json:"primary"`
}

func NewCreateIndexStatement() *CreateIndexStatement {
	return &CreateIndexStatement{}
}

func (this *CreateIndexStatement) SetExplainOnly(only bool) {
	this.ExplainOnly = only
}

func (this *CreateIndexStatement) IsExplainOnly() bool {
	return this.ExplainOnly
}

func (this *CreateIndexStatement) VerifySemantics() error {
	if this.Select != nil {
		this.Select.VerifySemantics()
		this.Bucket = this.Select.From.Bucket
	}
	return nil
}

func (this *CreateIndexStatement) Simplify() error {
	return nil
}

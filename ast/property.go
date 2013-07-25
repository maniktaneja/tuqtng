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
	"fmt"

	"github.com/couchbaselabs/tuqtng/query"
)

type Property struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

func NewProperty(path string) *Property {
	return &Property{
		Type: "property",
		Path: path,
	}
}

func (this *Property) Evaluate(item query.Item) (query.Value, error) {
	if item == nil {
		return nil, &query.Undefined{this.Path}
	}
	rv, err := item.GetPath(this.Path)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (this *Property) String() string {
	return fmt.Sprintf("%s", this.Path)
}

func (this *Property) Validate() error {
	return nil
}

// NOTE this should only be called on leading property references, not ones deeper in a chain (ie, for a.b.c, only a, not b or c)
func (this *Property) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	// this test is not needed when there are no aliases in the from clause (expression evaluation only)
	if len(aliases) > 0 {
		for _, alias := range aliases {
			if this.Path == alias {
				return nil, nil
			}
		}
		if defaultAlias != "" {
			return NewDotMemberOperator(NewProperty(defaultAlias), this), nil
		} else {
			return nil, fmt.Errorf("Property reference %s missing qualifier bucket/alias", this.Path)
		}
	}
	return nil, nil
}

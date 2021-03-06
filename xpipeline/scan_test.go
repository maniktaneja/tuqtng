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
	"testing"

	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/catalog/mock"
	"github.com/couchbaselabs/tuqtng/misc"
)

func TestScan(t *testing.T) {

	mocksite, err := mock.NewSite("mock:items=10")
	if err != nil {
		t.Errorf("Error creating mock site")
	}
	pool, err := mocksite.PoolByName("p0")
	if err != nil {
		t.Errorf("Error accessing pool p0")
	}
	bucket, err := pool.BucketByName("b0")
	if err != nil {
		t.Errorf("Error accessing bucket b0")
	}
	index, err := bucket.IndexByName("all_docs")
	if err != nil {
		t.Errorf("Error accessing scanner all_docs")
	}

	scanIndex := index.(catalog.ScanIndex)
	scan := NewScan(bucket, scanIndex, nil, "")

	scanItemChannel, _ := scan.GetChannels()

	stopChannel := make(misc.StopChannel)
	go scan.Run(stopChannel)

	count := 0
	for item := range scanItemChannel {
		meta := item.GetAttachment("meta")
		if meta == nil {
			t.Errorf("Expected item to contain metadata")
		}
		_, ok := meta.(map[string]interface{})["id"]
		if !ok {
			t.Errorf("Expected metadata to contain an id")
		}
		count++
	}

	if count != 10 {
		t.Errorf("Expected %d items, got %d", 10, count)
	}

}

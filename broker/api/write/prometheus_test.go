// Licensed to LinDB under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. LinDB licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package write

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/lindb/lindb/mock"
	"github.com/lindb/lindb/replication"
)

func TestPrometheusWrite_Write(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		readAllFunc = ioutil.ReadAll
		ctrl.Finish()
	}()

	cm := replication.NewMockChannelManager(ctrl)
	api := NewPrometheusWriter(cm)
	// case 1: param error
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodPut,
		URL:            "/metric/prometheus",
		HandlerFunc:    api.Write,
		ExpectHTTPCode: 500,
	})
	// case 2: read request body err
	readAllFunc = func(r io.Reader) (bytes []byte, err error) {
		return nil, fmt.Errorf("err")
	}
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodPut,
		URL:            "/metric/prometheus?db=dal",
		HandlerFunc:    api.Write,
		ExpectHTTPCode: 500,
	})
	// case 3: write wal err
	input := `# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# 	TYPE go_gc_duration_seconds summary
go_gc_duration_seconds { quantile = "0.9999" } NaN
go_gc_duration_seconds_count 9
go_gc_duration_seconds_sum 90
`
	readAllFunc = func(r io.Reader) (bytes []byte, err error) {
		return []byte(input), nil
	}
	cm.EXPECT().Write(gomock.Any(), gomock.Any()).Return(errors.New("err"))
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodPut,
		URL:            "/metric/prometheus?db=dal&cluster=dal&c=1",
		HandlerFunc:    api.Write,
		ExpectHTTPCode: 500,
	})
	// case 4: write wal success
	cm.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodPut,
		URL:            "/metric/prometheus?db=dal&cluster=dal&c=1",
		HandlerFunc:    api.Write,
		ExpectHTTPCode: 204,
	})
	// case 5: parse prometheus data err
	input = "# HELP go_gc_duration_seconds A summary of the GC invocation durations"
	readAllFunc = func(r io.Reader) (bytes []byte, err error) {
		return []byte(input), nil
	}
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodPut,
		URL:            "/metric/prometheus?db=dal&cluster=dal&c=1",
		HandlerFunc:    api.Write,
		ExpectHTTPCode: 500,
	})
}

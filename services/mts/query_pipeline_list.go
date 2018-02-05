package mts

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) QueryPipelineList(request *QueryPipelineListRequest) (response *QueryPipelineListResponse, err error) {
	response = CreateQueryPipelineListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryPipelineListWithChan(request *QueryPipelineListRequest) (<-chan *QueryPipelineListResponse, <-chan error) {
	responseChan := make(chan *QueryPipelineListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPipelineList(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) QueryPipelineListWithCallback(request *QueryPipelineListRequest, callback func(response *QueryPipelineListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPipelineListResponse
		var err error
		defer close(result)
		response, err = client.QueryPipelineList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type QueryPipelineListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PipelineIds          string           `position:"Query" name:"PipelineIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type QueryPipelineListResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	NonExistPids NonExistPids `json:"NonExistPids" xml:"NonExistPids"`
	PipelineList PipelineList `json:"PipelineList" xml:"PipelineList"`
}

func CreateQueryPipelineListRequest() (request *QueryPipelineListRequest) {
	request = &QueryPipelineListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryPipelineList", "", "")
	return
}

func CreateQueryPipelineListResponse() (response *QueryPipelineListResponse) {
	response = &QueryPipelineListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

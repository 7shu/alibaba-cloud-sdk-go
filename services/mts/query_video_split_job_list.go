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

func (client *Client) QueryVideoSplitJobList(request *QueryVideoSplitJobListRequest) (response *QueryVideoSplitJobListResponse, err error) {
	response = CreateQueryVideoSplitJobListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryVideoSplitJobListWithChan(request *QueryVideoSplitJobListRequest) (<-chan *QueryVideoSplitJobListResponse, <-chan error) {
	responseChan := make(chan *QueryVideoSplitJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryVideoSplitJobList(request)
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

func (client *Client) QueryVideoSplitJobListWithCallback(request *QueryVideoSplitJobListRequest, callback func(response *QueryVideoSplitJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryVideoSplitJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryVideoSplitJobList(request)
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

type QueryVideoSplitJobListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type QueryVideoSplitJobListResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	NonExistIds NonExistIds `json:"NonExistIds" xml:"NonExistIds"`
	JobList     JobList     `json:"JobList" xml:"JobList"`
}

func CreateQueryVideoSplitJobListRequest() (request *QueryVideoSplitJobListRequest) {
	request = &QueryVideoSplitJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryVideoSplitJobList", "", "")
	return
}

func CreateQueryVideoSplitJobListResponse() (response *QueryVideoSplitJobListResponse) {
	response = &QueryVideoSplitJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

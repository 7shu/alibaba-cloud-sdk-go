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

// QuerySnapshotJobList invokes the mts.QuerySnapshotJobList API synchronously
// api document: https://help.aliyun.com/api/mts/querysnapshotjoblist.html
func (client *Client) QuerySnapshotJobList(request *QuerySnapshotJobListRequest) (response *QuerySnapshotJobListResponse, err error) {
	response = CreateQuerySnapshotJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySnapshotJobListWithChan invokes the mts.QuerySnapshotJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/querysnapshotjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySnapshotJobListWithChan(request *QuerySnapshotJobListRequest) (<-chan *QuerySnapshotJobListResponse, <-chan error) {
	responseChan := make(chan *QuerySnapshotJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySnapshotJobList(request)
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

// QuerySnapshotJobListWithCallback invokes the mts.QuerySnapshotJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/querysnapshotjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySnapshotJobListWithCallback(request *QuerySnapshotJobListRequest, callback func(response *QuerySnapshotJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySnapshotJobListResponse
		var err error
		defer close(result)
		response, err = client.QuerySnapshotJobList(request)
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

// QuerySnapshotJobListRequest is the request struct for api QuerySnapshotJobList
type QuerySnapshotJobListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SnapshotJobIds       string           `position:"Query" name:"SnapshotJobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// QuerySnapshotJobListResponse is the response struct for api QuerySnapshotJobList
type QuerySnapshotJobListResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	NonExistSnapshotJobIds NonExistSnapshotJobIds `json:"NonExistSnapshotJobIds" xml:"NonExistSnapshotJobIds"`
	SnapshotJobList        SnapshotJobList        `json:"SnapshotJobList" xml:"SnapshotJobList"`
}

// CreateQuerySnapshotJobListRequest creates a request to invoke QuerySnapshotJobList API
func CreateQuerySnapshotJobListRequest(request *QuerySnapshotJobListRequest) {
	request = &QuerySnapshotJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QuerySnapshotJobList", "mts", "openAPI")
	return
}

// CreateQuerySnapshotJobListResponse creates a response to parse from QuerySnapshotJobList response
func CreateQuerySnapshotJobListResponse() (response *QuerySnapshotJobListResponse) {
	response = &QuerySnapshotJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package ram

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

// ListUsers invokes the ram.ListUsers API synchronously
// api document: https://help.aliyun.com/api/ram/listusers.html
func (client *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
	response = CreateListUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListUsersWithChan invokes the ram.ListUsers API asynchronously
// api document: https://help.aliyun.com/api/ram/listusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUsersWithChan(request *ListUsersRequest) (<-chan *ListUsersResponse, <-chan error) {
	responseChan := make(chan *ListUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUsers(request)
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

// ListUsersWithCallback invokes the ram.ListUsers API asynchronously
// api document: https://help.aliyun.com/api/ram/listusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUsersWithCallback(request *ListUsersRequest, callback func(response *ListUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUsersResponse
		var err error
		defer close(result)
		response, err = client.ListUsers(request)
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

// ListUsersRequest is the request struct for api ListUsers
type ListUsersRequest struct {
	*requests.RpcRequest
	Marker   string           `position:"Query" name:"Marker"`
	MaxItems requests.Integer `position:"Query" name:"MaxItems"`
}

// ListUsersResponse is the response struct for api ListUsers
type ListUsersResponse struct {
	*responses.BaseResponse
	RequestId   string           `json:"RequestId" xml:"RequestId"`
	IsTruncated bool             `json:"IsTruncated" xml:"IsTruncated"`
	Marker      string           `json:"Marker" xml:"Marker"`
	Users       UsersInListUsers `json:"Users" xml:"Users"`
}

// CreateListUsersRequest creates a request to invoke ListUsers API
func CreateListUsersRequest(request *ListUsersRequest) {
	request = &ListUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListUsers", "", "")
	return
}

// CreateListUsersResponse creates a response to parse from ListUsers response
func CreateListUsersResponse() (response *ListUsersResponse) {
	response = &ListUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

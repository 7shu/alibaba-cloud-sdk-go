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

// DeleteUser invokes the ram.DeleteUser API synchronously
// api document: https://help.aliyun.com/api/ram/deleteuser.html
func (client *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
	response = CreateDeleteUserResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUserWithChan invokes the ram.DeleteUser API asynchronously
// api document: https://help.aliyun.com/api/ram/deleteuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUserWithChan(request *DeleteUserRequest) (<-chan *DeleteUserResponse, <-chan error) {
	responseChan := make(chan *DeleteUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUser(request)
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

// DeleteUserWithCallback invokes the ram.DeleteUser API asynchronously
// api document: https://help.aliyun.com/api/ram/deleteuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUserWithCallback(request *DeleteUserRequest, callback func(response *DeleteUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUserResponse
		var err error
		defer close(result)
		response, err = client.DeleteUser(request)
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

// DeleteUserRequest is the request struct for api DeleteUser
type DeleteUserRequest struct {
	*requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

// DeleteUserResponse is the response struct for api DeleteUser
type DeleteUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUserRequest creates a request to invoke DeleteUser API
func CreateDeleteUserRequest(request *DeleteUserRequest) {
	request = &DeleteUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DeleteUser", "", "")
	return
}

// CreateDeleteUserResponse creates a response to parse from DeleteUser response
func CreateDeleteUserResponse() (response *DeleteUserResponse) {
	response = &DeleteUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

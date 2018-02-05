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

func (client *Client) DeletePolicyVersion(request *DeletePolicyVersionRequest) (response *DeletePolicyVersionResponse, err error) {
	response = CreateDeletePolicyVersionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeletePolicyVersionWithChan(request *DeletePolicyVersionRequest) (<-chan *DeletePolicyVersionResponse, <-chan error) {
	responseChan := make(chan *DeletePolicyVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePolicyVersion(request)
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

func (client *Client) DeletePolicyVersionWithCallback(request *DeletePolicyVersionRequest, callback func(response *DeletePolicyVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePolicyVersionResponse
		var err error
		defer close(result)
		response, err = client.DeletePolicyVersion(request)
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

type DeletePolicyVersionRequest struct {
	*requests.RpcRequest
	PolicyName string `position:"Query" name:"PolicyName"`
	VersionId  string `position:"Query" name:"VersionId"`
}

type DeletePolicyVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeletePolicyVersionRequest() (request *DeletePolicyVersionRequest) {
	request = &DeletePolicyVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DeletePolicyVersion", "", "")
	return
}

func CreateDeletePolicyVersionResponse() (response *DeletePolicyVersionResponse) {
	response = &DeletePolicyVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package ecs

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

func (client *Client) ApplyAutoSnapshotPolicy(request *ApplyAutoSnapshotPolicyRequest) (response *ApplyAutoSnapshotPolicyResponse, err error) {
	response = CreateApplyAutoSnapshotPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ApplyAutoSnapshotPolicyWithChan(request *ApplyAutoSnapshotPolicyRequest) (<-chan *ApplyAutoSnapshotPolicyResponse, <-chan error) {
	responseChan := make(chan *ApplyAutoSnapshotPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyAutoSnapshotPolicy(request)
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

func (client *Client) ApplyAutoSnapshotPolicyWithCallback(request *ApplyAutoSnapshotPolicyRequest, callback func(response *ApplyAutoSnapshotPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyAutoSnapshotPolicyResponse
		var err error
		defer close(result)
		response, err = client.ApplyAutoSnapshotPolicy(request)
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

type ApplyAutoSnapshotPolicyRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AutoSnapshotPolicyId string           `position:"Query" name:"autoSnapshotPolicyId"`
	DiskIds              string           `position:"Query" name:"diskIds"`
}

type ApplyAutoSnapshotPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateApplyAutoSnapshotPolicyRequest() (request *ApplyAutoSnapshotPolicyRequest) {
	request = &ApplyAutoSnapshotPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ApplyAutoSnapshotPolicy", "", "")
	return
}

func CreateApplyAutoSnapshotPolicyResponse() (response *ApplyAutoSnapshotPolicyResponse) {
	response = &ApplyAutoSnapshotPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

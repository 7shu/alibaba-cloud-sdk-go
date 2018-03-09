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

// CheckAutoSnapshotPolicy invokes the ecs.CheckAutoSnapshotPolicy API synchronously
// api document: https://help.aliyun.com/api/ecs/checkautosnapshotpolicy.html
func (client *Client) CheckAutoSnapshotPolicy(request *CheckAutoSnapshotPolicyRequest) (response *CheckAutoSnapshotPolicyResponse, err error) {
	response = CreateCheckAutoSnapshotPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CheckAutoSnapshotPolicyWithChan invokes the ecs.CheckAutoSnapshotPolicy API asynchronously
// api document: https://help.aliyun.com/api/ecs/checkautosnapshotpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckAutoSnapshotPolicyWithChan(request *CheckAutoSnapshotPolicyRequest) (<-chan *CheckAutoSnapshotPolicyResponse, <-chan error) {
	responseChan := make(chan *CheckAutoSnapshotPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckAutoSnapshotPolicy(request)
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

// CheckAutoSnapshotPolicyWithCallback invokes the ecs.CheckAutoSnapshotPolicy API asynchronously
// api document: https://help.aliyun.com/api/ecs/checkautosnapshotpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckAutoSnapshotPolicyWithCallback(request *CheckAutoSnapshotPolicyRequest, callback func(response *CheckAutoSnapshotPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckAutoSnapshotPolicyResponse
		var err error
		defer close(result)
		response, err = client.CheckAutoSnapshotPolicy(request)
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

// CheckAutoSnapshotPolicyRequest is the request struct for api CheckAutoSnapshotPolicy
type CheckAutoSnapshotPolicyRequest struct {
	*requests.RpcRequest
	OwnerId                           requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount              string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId                   requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount                      string           `position:"Query" name:"OwnerAccount"`
	SystemDiskPolicyEnabled           requests.Boolean `position:"Query" name:"SystemDiskPolicyEnabled"`
	SystemDiskPolicyTimePeriod        requests.Integer `position:"Query" name:"SystemDiskPolicyTimePeriod"`
	SystemDiskPolicyRetentionDays     requests.Integer `position:"Query" name:"SystemDiskPolicyRetentionDays"`
	SystemDiskPolicyRetentionLastWeek requests.Boolean `position:"Query" name:"SystemDiskPolicyRetentionLastWeek"`
	DataDiskPolicyEnabled             requests.Boolean `position:"Query" name:"DataDiskPolicyEnabled"`
	DataDiskPolicyTimePeriod          requests.Integer `position:"Query" name:"DataDiskPolicyTimePeriod"`
	DataDiskPolicyRetentionDays       requests.Integer `position:"Query" name:"DataDiskPolicyRetentionDays"`
	DataDiskPolicyRetentionLastWeek   requests.Boolean `position:"Query" name:"DataDiskPolicyRetentionLastWeek"`
}

// CheckAutoSnapshotPolicyResponse is the response struct for api CheckAutoSnapshotPolicy
type CheckAutoSnapshotPolicyResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	AutoSnapshotOccupation int    `json:"AutoSnapshotOccupation" xml:"AutoSnapshotOccupation"`
	IsPermittedModify      string `json:"IsPermittedModify" xml:"IsPermittedModify"`
}

// CreateCheckAutoSnapshotPolicyRequest creates a request to invoke CheckAutoSnapshotPolicy API
func CreateCheckAutoSnapshotPolicyRequest(request *CheckAutoSnapshotPolicyRequest) {
	request = &CheckAutoSnapshotPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CheckAutoSnapshotPolicy", "ecs", "openAPI")
	return
}

// CreateCheckAutoSnapshotPolicyResponse creates a response to parse from CheckAutoSnapshotPolicy response
func CreateCheckAutoSnapshotPolicyResponse() (response *CheckAutoSnapshotPolicyResponse) {
	response = &CheckAutoSnapshotPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

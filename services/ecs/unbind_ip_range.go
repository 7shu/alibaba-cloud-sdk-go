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

// UnbindIpRange invokes the ecs.UnbindIpRange API synchronously
// api document: https://help.aliyun.com/api/ecs/unbindiprange.html
func (client *Client) UnbindIpRange(request *UnbindIpRangeRequest) (response *UnbindIpRangeResponse, err error) {
	response = CreateUnbindIpRangeResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindIpRangeWithChan invokes the ecs.UnbindIpRange API asynchronously
// api document: https://help.aliyun.com/api/ecs/unbindiprange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindIpRangeWithChan(request *UnbindIpRangeRequest) (<-chan *UnbindIpRangeResponse, <-chan error) {
	responseChan := make(chan *UnbindIpRangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindIpRange(request)
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

// UnbindIpRangeWithCallback invokes the ecs.UnbindIpRange API asynchronously
// api document: https://help.aliyun.com/api/ecs/unbindiprange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindIpRangeWithCallback(request *UnbindIpRangeRequest, callback func(response *UnbindIpRangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindIpRangeResponse
		var err error
		defer close(result)
		response, err = client.UnbindIpRange(request)
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

// UnbindIpRangeRequest is the request struct for api UnbindIpRange
type UnbindIpRangeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	IpAddress            string           `position:"Query" name:"IpAddress"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// UnbindIpRangeResponse is the response struct for api UnbindIpRange
type UnbindIpRangeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindIpRangeRequest creates a request to invoke UnbindIpRange API
func CreateUnbindIpRangeRequest(request *UnbindIpRangeRequest) {
	request = &UnbindIpRangeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "UnbindIpRange", "ecs", "openAPI")
	return
}

// CreateUnbindIpRangeResponse creates a response to parse from UnbindIpRange response
func CreateUnbindIpRangeResponse() (response *UnbindIpRangeResponse) {
	response = &UnbindIpRangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package aegis

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

// DescribeStrategyTarget invokes the aegis.DescribeStrategyTarget API synchronously
// api document: https://help.aliyun.com/api/aegis/describestrategytarget.html
func (client *Client) DescribeStrategyTarget(request *DescribeStrategyTargetRequest) (response *DescribeStrategyTargetResponse, err error) {
	response = CreateDescribeStrategyTargetResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStrategyTargetWithChan invokes the aegis.DescribeStrategyTarget API asynchronously
// api document: https://help.aliyun.com/api/aegis/describestrategytarget.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStrategyTargetWithChan(request *DescribeStrategyTargetRequest) (<-chan *DescribeStrategyTargetResponse, <-chan error) {
	responseChan := make(chan *DescribeStrategyTargetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStrategyTarget(request)
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

// DescribeStrategyTargetWithCallback invokes the aegis.DescribeStrategyTarget API asynchronously
// api document: https://help.aliyun.com/api/aegis/describestrategytarget.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStrategyTargetWithCallback(request *DescribeStrategyTargetRequest, callback func(response *DescribeStrategyTargetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStrategyTargetResponse
		var err error
		defer close(result)
		response, err = client.DescribeStrategyTarget(request)
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

// DescribeStrategyTargetRequest is the request struct for api DescribeStrategyTarget
type DescribeStrategyTargetRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Type            string           `position:"Query" name:"Type"`
	Config          string           `position:"Query" name:"Config"`
	Target          string           `position:"Query" name:"Target"`
}

// DescribeStrategyTargetResponse is the response struct for api DescribeStrategyTarget
type DescribeStrategyTargetResponse struct {
	*responses.BaseResponse
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	Count           int              `json:"Count" xml:"Count"`
	StrategyTargets []StrategyTarget `json:"StrategyTargets" xml:"StrategyTargets"`
}

// CreateDescribeStrategyTargetRequest creates a request to invoke DescribeStrategyTarget API
func CreateDescribeStrategyTargetRequest(request *DescribeStrategyTargetRequest) {
	request = &DescribeStrategyTargetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeStrategyTarget", "vipaegis", "openAPI")
	return
}

// CreateDescribeStrategyTargetResponse creates a response to parse from DescribeStrategyTarget response
func CreateDescribeStrategyTargetResponse() (response *DescribeStrategyTargetResponse) {
	response = &DescribeStrategyTargetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

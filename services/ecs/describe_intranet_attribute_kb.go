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

// DescribeIntranetAttributeKb invokes the ecs.DescribeIntranetAttributeKb API synchronously
// api document: https://help.aliyun.com/api/ecs/describeintranetattributekb.html
func (client *Client) DescribeIntranetAttributeKb(request *DescribeIntranetAttributeKbRequest) (response *DescribeIntranetAttributeKbResponse, err error) {
	response = CreateDescribeIntranetAttributeKbResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIntranetAttributeKbWithChan invokes the ecs.DescribeIntranetAttributeKb API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeintranetattributekb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIntranetAttributeKbWithChan(request *DescribeIntranetAttributeKbRequest) (<-chan *DescribeIntranetAttributeKbResponse, <-chan error) {
	responseChan := make(chan *DescribeIntranetAttributeKbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIntranetAttributeKb(request)
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

// DescribeIntranetAttributeKbWithCallback invokes the ecs.DescribeIntranetAttributeKb API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeintranetattributekb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIntranetAttributeKbWithCallback(request *DescribeIntranetAttributeKbRequest, callback func(response *DescribeIntranetAttributeKbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIntranetAttributeKbResponse
		var err error
		defer close(result)
		response, err = client.DescribeIntranetAttributeKb(request)
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

// DescribeIntranetAttributeKbRequest is the request struct for api DescribeIntranetAttributeKb
type DescribeIntranetAttributeKbRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// DescribeIntranetAttributeKbResponse is the response struct for api DescribeIntranetAttributeKb
type DescribeIntranetAttributeKbResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	InstanceId              string `json:"InstanceId" xml:"InstanceId"`
	VlanId                  string `json:"VlanId" xml:"VlanId"`
	IntranetIpAddress       string `json:"IntranetIpAddress" xml:"IntranetIpAddress"`
	IntranetMaxBandwidthIn  int    `json:"IntranetMaxBandwidthIn" xml:"IntranetMaxBandwidthIn"`
	IntranetMaxBandwidthOut int    `json:"IntranetMaxBandwidthOut" xml:"IntranetMaxBandwidthOut"`
}

// CreateDescribeIntranetAttributeKbRequest creates a request to invoke DescribeIntranetAttributeKb API
func CreateDescribeIntranetAttributeKbRequest(request *DescribeIntranetAttributeKbRequest) {
	request = &DescribeIntranetAttributeKbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeIntranetAttributeKb", "ecs", "openAPI")
	return
}

// CreateDescribeIntranetAttributeKbResponse creates a response to parse from DescribeIntranetAttributeKb response
func CreateDescribeIntranetAttributeKbResponse() (response *DescribeIntranetAttributeKbResponse) {
	response = &DescribeIntranetAttributeKbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

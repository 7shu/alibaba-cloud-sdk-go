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

// DescribeEipAddresses invokes the ecs.DescribeEipAddresses API synchronously
// api document: https://help.aliyun.com/api/ecs/describeeipaddresses.html
func (client *Client) DescribeEipAddresses(request *DescribeEipAddressesRequest) (response *DescribeEipAddressesResponse, err error) {
	response = CreateDescribeEipAddressesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEipAddressesWithChan invokes the ecs.DescribeEipAddresses API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeeipaddresses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEipAddressesWithChan(request *DescribeEipAddressesRequest) (<-chan *DescribeEipAddressesResponse, <-chan error) {
	responseChan := make(chan *DescribeEipAddressesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEipAddresses(request)
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

// DescribeEipAddressesWithCallback invokes the ecs.DescribeEipAddresses API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeeipaddresses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEipAddressesWithCallback(request *DescribeEipAddressesRequest, callback func(response *DescribeEipAddressesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEipAddressesResponse
		var err error
		defer close(result)
		response, err = client.DescribeEipAddresses(request)
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

// DescribeEipAddressesRequest is the request struct for api DescribeEipAddresses
type DescribeEipAddressesRequest struct {
	*requests.RpcRequest
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Status                 string           `position:"Query" name:"Status"`
	EipAddress             string           `position:"Query" name:"EipAddress"`
	AllocationId           string           `position:"Query" name:"AllocationId"`
	PageNumber             requests.Integer `position:"Query" name:"PageNumber"`
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	Filter1Key             string           `position:"Query" name:"Filter.1.Key"`
	Filter2Key             string           `position:"Query" name:"Filter.2.Key"`
	Filter1Value           string           `position:"Query" name:"Filter.1.Value"`
	Filter2Value           string           `position:"Query" name:"Filter.2.Value"`
	LockReason             string           `position:"Query" name:"LockReason"`
	AssociatedInstanceType string           `position:"Query" name:"AssociatedInstanceType"`
	AssociatedInstanceId   string           `position:"Query" name:"AssociatedInstanceId"`
	ChargeType             string           `position:"Query" name:"ChargeType"`
}

// DescribeEipAddressesResponse is the response struct for api DescribeEipAddresses
type DescribeEipAddressesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	TotalCount   int          `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int          `json:"PageNumber" xml:"PageNumber"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	EipAddresses EipAddresses `json:"EipAddresses" xml:"EipAddresses"`
}

// CreateDescribeEipAddressesRequest creates a request to invoke DescribeEipAddresses API
func CreateDescribeEipAddressesRequest(request *DescribeEipAddressesRequest) {
	request = &DescribeEipAddressesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeEipAddresses", "ecs", "openAPI")
	return
}

// CreateDescribeEipAddressesResponse creates a response to parse from DescribeEipAddresses response
func CreateDescribeEipAddressesResponse() (response *DescribeEipAddressesResponse) {
	response = &DescribeEipAddressesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

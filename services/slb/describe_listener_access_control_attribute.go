package slb

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

// DescribeListenerAccessControlAttribute invokes the slb.DescribeListenerAccessControlAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/describelisteneraccesscontrolattribute.html
func (client *Client) DescribeListenerAccessControlAttribute(request *DescribeListenerAccessControlAttributeRequest) (response *DescribeListenerAccessControlAttributeResponse, err error) {
	response = CreateDescribeListenerAccessControlAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeListenerAccessControlAttributeWithChan invokes the slb.DescribeListenerAccessControlAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describelisteneraccesscontrolattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeListenerAccessControlAttributeWithChan(request *DescribeListenerAccessControlAttributeRequest) (<-chan *DescribeListenerAccessControlAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeListenerAccessControlAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeListenerAccessControlAttribute(request)
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

// DescribeListenerAccessControlAttributeWithCallback invokes the slb.DescribeListenerAccessControlAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describelisteneraccesscontrolattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeListenerAccessControlAttributeWithCallback(request *DescribeListenerAccessControlAttributeRequest, callback func(response *DescribeListenerAccessControlAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeListenerAccessControlAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeListenerAccessControlAttribute(request)
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

// DescribeListenerAccessControlAttributeRequest is the request struct for api DescribeListenerAccessControlAttribute
type DescribeListenerAccessControlAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// DescribeListenerAccessControlAttributeResponse is the response struct for api DescribeListenerAccessControlAttribute
type DescribeListenerAccessControlAttributeResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	AccessControlStatus string `json:"AccessControlStatus" xml:"AccessControlStatus"`
	SourceItems         string `json:"SourceItems" xml:"SourceItems"`
}

// CreateDescribeListenerAccessControlAttributeRequest creates a request to invoke DescribeListenerAccessControlAttribute API
func CreateDescribeListenerAccessControlAttributeRequest(request *DescribeListenerAccessControlAttributeRequest) {
	request = &DescribeListenerAccessControlAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeListenerAccessControlAttribute", "slb", "openAPI")
	return
}

// CreateDescribeListenerAccessControlAttributeResponse creates a response to parse from DescribeListenerAccessControlAttribute response
func CreateDescribeListenerAccessControlAttributeResponse() (response *DescribeListenerAccessControlAttributeResponse) {
	response = &DescribeListenerAccessControlAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package rds

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

// DescribeTags invokes the rds.DescribeTags API synchronously
// api document: https://help.aliyun.com/api/rds/describetags.html
func (client *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
	response = CreateDescribeTagsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTagsWithChan invokes the rds.DescribeTags API asynchronously
// api document: https://help.aliyun.com/api/rds/describetags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTagsWithChan(request *DescribeTagsRequest) (<-chan *DescribeTagsResponse, <-chan error) {
	responseChan := make(chan *DescribeTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTags(request)
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

// DescribeTagsWithCallback invokes the rds.DescribeTags API asynchronously
// api document: https://help.aliyun.com/api/rds/describetags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTagsWithCallback(request *DescribeTagsRequest, callback func(response *DescribeTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeTags(request)
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

// DescribeTagsRequest is the request struct for api DescribeTags
type DescribeTagsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Tags                 string           `position:"Query" name:"Tags"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// DescribeTagsResponse is the response struct for api DescribeTags
type DescribeTagsResponse struct {
	*responses.BaseResponse
	RequestId string              `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeTags `json:"Items" xml:"Items"`
}

// CreateDescribeTagsRequest creates a request to invoke DescribeTags API
func CreateDescribeTagsRequest(request *DescribeTagsRequest) {
	request = &DescribeTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeTags", "rds", "openAPI")
	return
}

// CreateDescribeTagsResponse creates a response to parse from DescribeTags response
func CreateDescribeTagsResponse() (response *DescribeTagsResponse) {
	response = &DescribeTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

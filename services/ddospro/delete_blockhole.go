package ddospro

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

// DeleteBlockhole invokes the ddospro.DeleteBlockhole API synchronously
// api document: https://help.aliyun.com/api/ddospro/deleteblockhole.html
func (client *Client) DeleteBlockhole(request *DeleteBlockholeRequest) (response *DeleteBlockholeResponse, err error) {
	response = CreateDeleteBlockholeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBlockholeWithChan invokes the ddospro.DeleteBlockhole API asynchronously
// api document: https://help.aliyun.com/api/ddospro/deleteblockhole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBlockholeWithChan(request *DeleteBlockholeRequest) (<-chan *DeleteBlockholeResponse, <-chan error) {
	responseChan := make(chan *DeleteBlockholeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBlockhole(request)
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

// DeleteBlockholeWithCallback invokes the ddospro.DeleteBlockhole API asynchronously
// api document: https://help.aliyun.com/api/ddospro/deleteblockhole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBlockholeWithCallback(request *DeleteBlockholeRequest, callback func(response *DeleteBlockholeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBlockholeResponse
		var err error
		defer close(result)
		response, err = client.DeleteBlockhole(request)
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

// DeleteBlockholeRequest is the request struct for api DeleteBlockhole
type DeleteBlockholeRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Vip             string           `position:"Query" name:"Vip"`
}

// DeleteBlockholeResponse is the response struct for api DeleteBlockhole
type DeleteBlockholeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteBlockholeRequest creates a request to invoke DeleteBlockhole API
func CreateDeleteBlockholeRequest(request *DeleteBlockholeRequest) {
	request = &DeleteBlockholeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DeleteBlockhole", "", "")
	return
}

// CreateDeleteBlockholeResponse creates a response to parse from DeleteBlockhole response
func CreateDeleteBlockholeResponse() (response *DeleteBlockholeResponse) {
	response = &DeleteBlockholeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

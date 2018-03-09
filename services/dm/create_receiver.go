package dm

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

// CreateReceiver invokes the dm.CreateReceiver API synchronously
// api document: https://help.aliyun.com/api/dm/createreceiver.html
func (client *Client) CreateReceiver(request *CreateReceiverRequest) (response *CreateReceiverResponse, err error) {
	response = CreateCreateReceiverResponse()
	err = client.DoAction(request, response)
	return
}

// CreateReceiverWithChan invokes the dm.CreateReceiver API asynchronously
// api document: https://help.aliyun.com/api/dm/createreceiver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateReceiverWithChan(request *CreateReceiverRequest) (<-chan *CreateReceiverResponse, <-chan error) {
	responseChan := make(chan *CreateReceiverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateReceiver(request)
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

// CreateReceiverWithCallback invokes the dm.CreateReceiver API asynchronously
// api document: https://help.aliyun.com/api/dm/createreceiver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateReceiverWithCallback(request *CreateReceiverRequest, callback func(response *CreateReceiverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateReceiverResponse
		var err error
		defer close(result)
		response, err = client.CreateReceiver(request)
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

// CreateReceiverRequest is the request struct for api CreateReceiver
type CreateReceiverRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ReceiversName        string           `position:"Query" name:"ReceiversName"`
	ReceiversAlias       string           `position:"Query" name:"ReceiversAlias"`
	Desc                 string           `position:"Query" name:"Desc"`
}

// CreateReceiverResponse is the response struct for api CreateReceiver
type CreateReceiverResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateReceiverRequest creates a request to invoke CreateReceiver API
func CreateCreateReceiverRequest(request *CreateReceiverRequest) {
	request = &CreateReceiverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "CreateReceiver", "", "")
	return
}

// CreateCreateReceiverResponse creates a response to parse from CreateReceiver response
func CreateCreateReceiverResponse() (response *CreateReceiverResponse) {
	response = &CreateReceiverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

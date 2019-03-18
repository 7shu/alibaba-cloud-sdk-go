package cas

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

// CreateOrderDocument invokes the cas.CreateOrderDocument API synchronously
// api document: https://help.aliyun.com/api/cas/createorderdocument.html
func (client *Client) CreateOrderDocument(request *CreateOrderDocumentRequest) (response *CreateOrderDocumentResponse, err error) {
	response = CreateCreateOrderDocumentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrderDocumentWithChan invokes the cas.CreateOrderDocument API asynchronously
// api document: https://help.aliyun.com/api/cas/createorderdocument.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrderDocumentWithChan(request *CreateOrderDocumentRequest) (<-chan *CreateOrderDocumentResponse, <-chan error) {
	responseChan := make(chan *CreateOrderDocumentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrderDocument(request)
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

// CreateOrderDocumentWithCallback invokes the cas.CreateOrderDocument API asynchronously
// api document: https://help.aliyun.com/api/cas/createorderdocument.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrderDocumentWithCallback(request *CreateOrderDocumentRequest, callback func(response *CreateOrderDocumentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrderDocumentResponse
		var err error
		defer close(result)
		response, err = client.CreateOrderDocument(request)
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

// CreateOrderDocumentRequest is the request struct for api CreateOrderDocument
type CreateOrderDocumentRequest struct {
	*requests.RpcRequest
	OssKey       string           `position:"Query" name:"OssKey"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	OrderId      requests.Integer `position:"Query" name:"OrderId"`
	DocumentType requests.Integer `position:"Query" name:"DocumentType"`
	Lang         string           `position:"Query" name:"Lang"`
	ExtName      string           `position:"Query" name:"ExtName"`
}

// CreateOrderDocumentResponse is the response struct for api CreateOrderDocument
type CreateOrderDocumentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateOrderDocumentRequest creates a request to invoke CreateOrderDocument API
func CreateCreateOrderDocumentRequest() (request *CreateOrderDocumentRequest) {
	request = &CreateOrderDocumentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2018-08-13", "CreateOrderDocument", "cas_esign_fdd", "openAPI")
	return
}

// CreateCreateOrderDocumentResponse creates a response to parse from CreateOrderDocument response
func CreateCreateOrderDocumentResponse() (response *CreateOrderDocumentResponse) {
	response = &CreateOrderDocumentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package domain_intl

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

// SaveBatchTaskForCreatingOrderRenew invokes the domain_intl.SaveBatchTaskForCreatingOrderRenew API synchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskforcreatingorderrenew.html
func (client *Client) SaveBatchTaskForCreatingOrderRenew(request *SaveBatchTaskForCreatingOrderRenewRequest) (response *SaveBatchTaskForCreatingOrderRenewResponse, err error) {
	response = CreateSaveBatchTaskForCreatingOrderRenewResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForCreatingOrderRenewWithChan invokes the domain_intl.SaveBatchTaskForCreatingOrderRenew API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskforcreatingorderrenew.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveBatchTaskForCreatingOrderRenewWithChan(request *SaveBatchTaskForCreatingOrderRenewRequest) (<-chan *SaveBatchTaskForCreatingOrderRenewResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForCreatingOrderRenewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForCreatingOrderRenew(request)
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

// SaveBatchTaskForCreatingOrderRenewWithCallback invokes the domain_intl.SaveBatchTaskForCreatingOrderRenew API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskforcreatingorderrenew.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveBatchTaskForCreatingOrderRenewWithCallback(request *SaveBatchTaskForCreatingOrderRenewRequest, callback func(response *SaveBatchTaskForCreatingOrderRenewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForCreatingOrderRenewResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForCreatingOrderRenew(request)
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

// SaveBatchTaskForCreatingOrderRenewRequest is the request struct for api SaveBatchTaskForCreatingOrderRenew
type SaveBatchTaskForCreatingOrderRenewRequest struct {
	*requests.RpcRequest
	UserClientIp    string                                               `position:"Query" name:"UserClientIp"`
	Lang            string                                               `position:"Query" name:"Lang"`
	OrderRenewParam *[]SaveBatchTaskForCreatingOrderRenewOrderRenewParam `position:"Query" name:"OrderRenewParam"  type:"Repeated"`
}

// SaveBatchTaskForCreatingOrderRenewOrderRenewParam is a repeated param struct in SaveBatchTaskForCreatingOrderRenewRequest
type SaveBatchTaskForCreatingOrderRenewOrderRenewParam struct {
	DomainName            string `name:"DomainName"`
	CurrentExpirationDate string `name:"CurrentExpirationDate"`
	SubscriptionDuration  string `name:"SubscriptionDuration"`
}

// SaveBatchTaskForCreatingOrderRenewResponse is the response struct for api SaveBatchTaskForCreatingOrderRenew
type SaveBatchTaskForCreatingOrderRenewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForCreatingOrderRenewRequest creates a request to invoke SaveBatchTaskForCreatingOrderRenew API
func CreateSaveBatchTaskForCreatingOrderRenewRequest(request *SaveBatchTaskForCreatingOrderRenewRequest) {
	request = &SaveBatchTaskForCreatingOrderRenewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForCreatingOrderRenew", "", "")
	return
}

// CreateSaveBatchTaskForCreatingOrderRenewResponse creates a response to parse from SaveBatchTaskForCreatingOrderRenew response
func CreateSaveBatchTaskForCreatingOrderRenewResponse() (response *SaveBatchTaskForCreatingOrderRenewResponse) {
	response = &SaveBatchTaskForCreatingOrderRenewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

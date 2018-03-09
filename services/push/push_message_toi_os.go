package push

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

// PushMessageToiOS invokes the push.PushMessageToiOS API synchronously
// api document: https://help.aliyun.com/api/push/pushmessagetoios.html
func (client *Client) PushMessageToiOS(request *PushMessageToiOSRequest) (response *PushMessageToiOSResponse, err error) {
	response = CreatePushMessageToiOSResponse()
	err = client.DoAction(request, response)
	return
}

// PushMessageToiOSWithChan invokes the push.PushMessageToiOS API asynchronously
// api document: https://help.aliyun.com/api/push/pushmessagetoios.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PushMessageToiOSWithChan(request *PushMessageToiOSRequest) (<-chan *PushMessageToiOSResponse, <-chan error) {
	responseChan := make(chan *PushMessageToiOSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushMessageToiOS(request)
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

// PushMessageToiOSWithCallback invokes the push.PushMessageToiOS API asynchronously
// api document: https://help.aliyun.com/api/push/pushmessagetoios.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PushMessageToiOSWithCallback(request *PushMessageToiOSRequest, callback func(response *PushMessageToiOSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushMessageToiOSResponse
		var err error
		defer close(result)
		response, err = client.PushMessageToiOS(request)
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

// PushMessageToiOSRequest is the request struct for api PushMessageToiOS
type PushMessageToiOSRequest struct {
	*requests.RpcRequest
	AppKey      requests.Integer `position:"Query" name:"AppKey"`
	Target      string           `position:"Query" name:"Target"`
	TargetValue string           `position:"Query" name:"TargetValue"`
	Title       string           `position:"Query" name:"Title"`
	Body        string           `position:"Query" name:"Body"`
	JobKey      string           `position:"Query" name:"JobKey"`
}

// PushMessageToiOSResponse is the response struct for api PushMessageToiOS
type PushMessageToiOSResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MessageId string `json:"MessageId" xml:"MessageId"`
}

// CreatePushMessageToiOSRequest creates a request to invoke PushMessageToiOS API
func CreatePushMessageToiOSRequest(request *PushMessageToiOSRequest) {
	request = &PushMessageToiOSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "PushMessageToiOS", "", "")
	return
}

// CreatePushMessageToiOSResponse creates a response to parse from PushMessageToiOS response
func CreatePushMessageToiOSResponse() (response *PushMessageToiOSResponse) {
	response = &PushMessageToiOSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

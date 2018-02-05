package cms

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

func (client *Client) EnableActiceAlert(request *EnableActiceAlertRequest) (response *EnableActiceAlertResponse, err error) {
	response = CreateEnableActiceAlertResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) EnableActiceAlertWithChan(request *EnableActiceAlertRequest) (<-chan *EnableActiceAlertResponse, <-chan error) {
	responseChan := make(chan *EnableActiceAlertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableActiceAlert(request)
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

func (client *Client) EnableActiceAlertWithCallback(request *EnableActiceAlertRequest, callback func(response *EnableActiceAlertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableActiceAlertResponse
		var err error
		defer close(result)
		response, err = client.EnableActiceAlert(request)
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

type EnableActiceAlertRequest struct {
	*requests.RpcRequest
	Product string `position:"Query" name:"Product"`
	UserId  string `position:"Query" name:"UserId"`
}

type EnableActiceAlertResponse struct {
	*responses.BaseResponse
	Success bool   `json:"Success" xml:"Success"`
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

func CreateEnableActiceAlertRequest() (request *EnableActiceAlertRequest) {
	request = &EnableActiceAlertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "EnableActiceAlert", "", "")
	return
}

func CreateEnableActiceAlertResponse() (response *EnableActiceAlertResponse) {
	response = &EnableActiceAlertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

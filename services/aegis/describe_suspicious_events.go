package aegis

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

// DescribeSuspiciousEvents invokes the aegis.DescribeSuspiciousEvents API synchronously
// api document: https://help.aliyun.com/api/aegis/describesuspiciousevents.html
func (client *Client) DescribeSuspiciousEvents(request *DescribeSuspiciousEventsRequest) (response *DescribeSuspiciousEventsResponse, err error) {
	response = CreateDescribeSuspiciousEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSuspiciousEventsWithChan invokes the aegis.DescribeSuspiciousEvents API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesuspiciousevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspiciousEventsWithChan(request *DescribeSuspiciousEventsRequest) (<-chan *DescribeSuspiciousEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeSuspiciousEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSuspiciousEvents(request)
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

// DescribeSuspiciousEventsWithCallback invokes the aegis.DescribeSuspiciousEvents API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesuspiciousevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspiciousEventsWithCallback(request *DescribeSuspiciousEventsRequest, callback func(response *DescribeSuspiciousEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSuspiciousEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSuspiciousEvents(request)
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

// DescribeSuspiciousEventsRequest is the request struct for api DescribeSuspiciousEvents
type DescribeSuspiciousEventsRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	Uuid            string           `position:"Query" name:"Uuid"`
	Dealed          string           `position:"Query" name:"Dealed"`
	Remark          string           `position:"Query" name:"Remark"`
	Level           string           `position:"Query" name:"Level"`
	EventType       string           `position:"Query" name:"EventType"`
}

// DescribeSuspiciousEventsResponse is the response struct for api DescribeSuspiciousEvents
type DescribeSuspiciousEventsResponse struct {
	*responses.BaseResponse
	RequestId   string        `json:"RequestId" xml:"RequestId"`
	PageSize    int           `json:"PageSize" xml:"PageSize"`
	TotalCount  int           `json:"TotalCount" xml:"TotalCount"`
	CurrentPage int           `json:"CurrentPage" xml:"CurrentPage"`
	LogList     []LogListItem `json:"LogList" xml:"LogList"`
}

// CreateDescribeSuspiciousEventsRequest creates a request to invoke DescribeSuspiciousEvents API
func CreateDescribeSuspiciousEventsRequest(request *DescribeSuspiciousEventsRequest) {
	request = &DescribeSuspiciousEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeSuspiciousEvents", "vipaegis", "openAPI")
	return
}

// CreateDescribeSuspiciousEventsResponse creates a response to parse from DescribeSuspiciousEvents response
func CreateDescribeSuspiciousEventsResponse() (response *DescribeSuspiciousEventsResponse) {
	response = &DescribeSuspiciousEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

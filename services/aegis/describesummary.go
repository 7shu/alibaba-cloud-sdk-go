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

// Describesummary invokes the aegis.Describesummary API synchronously
// api document: https://help.aliyun.com/api/aegis/describesummary.html
func (client *Client) Describesummary(request *DescribesummaryRequest) (response *DescribesummaryResponse, err error) {
	response = CreateDescribesummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribesummaryWithChan invokes the aegis.Describesummary API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribesummaryWithChan(request *DescribesummaryRequest) (<-chan *DescribesummaryResponse, <-chan error) {
	responseChan := make(chan *DescribesummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Describesummary(request)
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

// DescribesummaryWithCallback invokes the aegis.Describesummary API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribesummaryWithCallback(request *DescribesummaryRequest, callback func(response *DescribesummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribesummaryResponse
		var err error
		defer close(result)
		response, err = client.Describesummary(request)
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

// DescribesummaryRequest is the request struct for api Describesummary
type DescribesummaryRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Uuids           string           `position:"Query" name:"Uuids"`
	TypeNames       string           `position:"Query" name:"TypeNames"`
	SubTypeNames    string           `position:"Query" name:"SubTypeNames"`
	RiskLevels      string           `position:"Query" name:"RiskLevels"`
	StatusList      string           `position:"Query" name:"StatusList"`
	RiskName        string           `position:"Query" name:"RiskName"`
	Dealed          string           `position:"Query" name:"Dealed"`
	StrategyId      requests.Integer `position:"Query" name:"StrategyId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
}

// DescribesummaryResponse is the response struct for api Describesummary
type DescribesummaryResponse struct {
	*responses.BaseResponse
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	Count           int              `json:"Count" xml:"Count"`
	PageSize        int              `json:"PageSize" xml:"PageSize"`
	TotalCount      int              `json:"TotalCount" xml:"TotalCount"`
	CurrentPage     int              `json:"CurrentPage" xml:"CurrentPage"`
	WarningSummarys []WarningSummary `json:"WarningSummarys" xml:"WarningSummarys"`
}

// CreateDescribesummaryRequest creates a request to invoke Describesummary API
func CreateDescribesummaryRequest(request *DescribesummaryRequest) {
	request = &DescribesummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "Describesummary", "vipaegis", "openAPI")
	return
}

// CreateDescribesummaryResponse creates a response to parse from Describesummary response
func CreateDescribesummaryResponse() (response *DescribesummaryResponse) {
	response = &DescribesummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

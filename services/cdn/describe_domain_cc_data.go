package cdn

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

func (client *Client) DescribeDomainCCData(request *DescribeDomainCCDataRequest) (response *DescribeDomainCCDataResponse, err error) {
	response = CreateDescribeDomainCCDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainCCDataWithChan(request *DescribeDomainCCDataRequest) (<-chan *DescribeDomainCCDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainCCDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainCCData(request)
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

func (client *Client) DescribeDomainCCDataWithCallback(request *DescribeDomainCCDataRequest, callback func(response *DescribeDomainCCDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainCCDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainCCData(request)
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

type DescribeDomainCCDataRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
}

type DescribeDomainCCDataResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	DomainName   string     `json:"DomainName" xml:"DomainName"`
	DataInterval string     `json:"DataInterval" xml:"DataInterval"`
	StartTime    string     `json:"StartTime" xml:"StartTime"`
	EndTime      string     `json:"EndTime" xml:"EndTime"`
	CCDataList   CCDataList `json:"CCDataList" xml:"CCDataList"`
}

func CreateDescribeDomainCCDataRequest() (request *DescribeDomainCCDataRequest) {
	request = &DescribeDomainCCDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCCData", "", "")
	return
}

func CreateDescribeDomainCCDataResponse() (response *DescribeDomainCCDataResponse) {
	response = &DescribeDomainCCDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

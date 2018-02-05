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

func (client *Client) BatchDescribeDomainBpsData(request *BatchDescribeDomainBpsDataRequest) (response *BatchDescribeDomainBpsDataResponse, err error) {
	response = CreateBatchDescribeDomainBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) BatchDescribeDomainBpsDataWithChan(request *BatchDescribeDomainBpsDataRequest) (<-chan *BatchDescribeDomainBpsDataResponse, <-chan error) {
	responseChan := make(chan *BatchDescribeDomainBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchDescribeDomainBpsData(request)
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

func (client *Client) BatchDescribeDomainBpsDataWithCallback(request *BatchDescribeDomainBpsDataRequest, callback func(response *BatchDescribeDomainBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchDescribeDomainBpsDataResponse
		var err error
		defer close(result)
		response, err = client.BatchDescribeDomainBpsData(request)
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

type BatchDescribeDomainBpsDataRequest struct {
	*requests.RpcRequest
}

type BatchDescribeDomainBpsDataResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	BpsDatas   BpsDatas `json:"BpsDatas" xml:"BpsDatas"`
}

func CreateBatchDescribeDomainBpsDataRequest() (request *BatchDescribeDomainBpsDataRequest) {
	request = &BatchDescribeDomainBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "BatchDescribeDomainBpsData", "", "")
	return
}

func CreateBatchDescribeDomainBpsDataResponse() (response *BatchDescribeDomainBpsDataResponse) {
	response = &BatchDescribeDomainBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

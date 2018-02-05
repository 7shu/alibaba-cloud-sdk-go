package ecs

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

func (client *Client) DescribeDiskMonitorData(request *DescribeDiskMonitorDataRequest) (response *DescribeDiskMonitorDataResponse, err error) {
	response = CreateDescribeDiskMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDiskMonitorDataWithChan(request *DescribeDiskMonitorDataRequest) (<-chan *DescribeDiskMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDiskMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDiskMonitorData(request)
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

func (client *Client) DescribeDiskMonitorDataWithCallback(request *DescribeDiskMonitorDataRequest, callback func(response *DescribeDiskMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDiskMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDiskMonitorData(request)
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

type DescribeDiskMonitorDataRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DiskId               string           `position:"Query" name:"DiskId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
	Period               requests.Integer `position:"Query" name:"Period"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeDiskMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	MonitorData MonitorData `json:"MonitorData" xml:"MonitorData"`
}

func CreateDescribeDiskMonitorDataRequest() (request *DescribeDiskMonitorDataRequest) {
	request = &DescribeDiskMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDiskMonitorData", "", "")
	return
}

func CreateDescribeDiskMonitorDataResponse() (response *DescribeDiskMonitorDataResponse) {
	response = &DescribeDiskMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

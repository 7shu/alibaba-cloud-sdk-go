package ddospro

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

func (client *Client) DescribeDomainQps(request *DescribeDomainQpsRequest) (response *DescribeDomainQpsResponse, err error) {
	response = CreateDescribeDomainQpsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainQpsWithChan(request *DescribeDomainQpsRequest) (<-chan *DescribeDomainQpsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainQpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainQps(request)
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

func (client *Client) DescribeDomainQpsWithCallback(request *DescribeDomainQpsRequest, callback func(response *DescribeDomainQpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainQpsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainQps(request)
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

type DescribeDomainQpsRequest struct {
	*requests.RpcRequest
	Domain          string           `position:"Query" name:"Domain"`
	StartDateMillis requests.Integer `position:"Query" name:"StartDateMillis"`
	EndDateMillis   requests.Integer `position:"Query" name:"EndDateMillis"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
}

type DescribeDomainQpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

func CreateDescribeDomainQpsRequest() (request *DescribeDomainQpsRequest) {
	request = &DescribeDomainQpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DescribeDomainQps", "", "")
	return
}

func CreateDescribeDomainQpsResponse() (response *DescribeDomainQpsResponse) {
	response = &DescribeDomainQpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

func (client *Client) QueryDnsHost(request *QueryDnsHostRequest) (response *QueryDnsHostResponse, err error) {
	response = CreateQueryDnsHostResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryDnsHostWithChan(request *QueryDnsHostRequest) (<-chan *QueryDnsHostResponse, <-chan error) {
	responseChan := make(chan *QueryDnsHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDnsHost(request)
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

func (client *Client) QueryDnsHostWithCallback(request *QueryDnsHostRequest, callback func(response *QueryDnsHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDnsHostResponse
		var err error
		defer close(result)
		response, err = client.QueryDnsHost(request)
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

type QueryDnsHostRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Lang       string `position:"Query" name:"Lang"`
}

type QueryDnsHostResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	DnsHostList []DnsHost `json:"DnsHostList" xml:"DnsHostList"`
}

func CreateQueryDnsHostRequest() (request *QueryDnsHostRequest) {
	request = &QueryDnsHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryDnsHost", "", "")
	return
}

func CreateQueryDnsHostResponse() (response *QueryDnsHostResponse) {
	response = &QueryDnsHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

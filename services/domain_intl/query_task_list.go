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

func (client *Client) QueryTaskList(request *QueryTaskListRequest) (response *QueryTaskListResponse, err error) {
	response = CreateQueryTaskListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryTaskListWithChan(request *QueryTaskListRequest) (<-chan *QueryTaskListResponse, <-chan error) {
	responseChan := make(chan *QueryTaskListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTaskList(request)
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

func (client *Client) QueryTaskListWithCallback(request *QueryTaskListRequest, callback func(response *QueryTaskListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskListResponse
		var err error
		defer close(result)
		response, err = client.QueryTaskList(request)
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

type QueryTaskListRequest struct {
	*requests.RpcRequest
	UserClientIp    string           `position:"Query" name:"UserClientIp"`
	Lang            string           `position:"Query" name:"Lang"`
	BeginCreateTime requests.Integer `position:"Query" name:"BeginCreateTime"`
	EndCreateTime   requests.Integer `position:"Query" name:"EndCreateTime"`
	PageNum         requests.Integer `position:"Query" name:"PageNum"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

type QueryTaskListResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int    `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int    `json:"CurrentPageNum" xml:"CurrentPageNum"`
	TotalPageNum   int    `json:"TotalPageNum" xml:"TotalPageNum"`
	PageSize       int    `json:"PageSize" xml:"PageSize"`
	PrePage        bool   `json:"PrePage" xml:"PrePage"`
	NextPage       bool   `json:"NextPage" xml:"NextPage"`
	Data           Data   `json:"Data" xml:"Data"`
}

func CreateQueryTaskListRequest() (request *QueryTaskListRequest) {
	request = &QueryTaskListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTaskList", "", "")
	return
}

func CreateQueryTaskListResponse() (response *QueryTaskListResponse) {
	response = &QueryTaskListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

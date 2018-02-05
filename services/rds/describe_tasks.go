package rds

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

func (client *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
	response = CreateDescribeTasksResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeTasksWithChan(request *DescribeTasksRequest) (<-chan *DescribeTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTasks(request)
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

func (client *Client) DescribeTasksWithCallback(request *DescribeTasksRequest, callback func(response *DescribeTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeTasks(request)
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

type DescribeTasksRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Status               string           `position:"Query" name:"Status"`
	TaskAction           string           `position:"Query" name:"TaskAction"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeTasksResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            Items  `json:"Items" xml:"Items"`
}

func CreateDescribeTasksRequest() (request *DescribeTasksRequest) {
	request = &DescribeTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeTasks", "", "")
	return
}

func CreateDescribeTasksResponse() (response *DescribeTasksResponse) {
	response = &DescribeTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

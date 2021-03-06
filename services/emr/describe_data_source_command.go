package emr

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

// DescribeDataSourceCommand invokes the emr.DescribeDataSourceCommand API synchronously
// api document: https://help.aliyun.com/api/emr/describedatasourcecommand.html
func (client *Client) DescribeDataSourceCommand(request *DescribeDataSourceCommandRequest) (response *DescribeDataSourceCommandResponse, err error) {
	response = CreateDescribeDataSourceCommandResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataSourceCommandWithChan invokes the emr.DescribeDataSourceCommand API asynchronously
// api document: https://help.aliyun.com/api/emr/describedatasourcecommand.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataSourceCommandWithChan(request *DescribeDataSourceCommandRequest) (<-chan *DescribeDataSourceCommandResponse, <-chan error) {
	responseChan := make(chan *DescribeDataSourceCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataSourceCommand(request)
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

// DescribeDataSourceCommandWithCallback invokes the emr.DescribeDataSourceCommand API asynchronously
// api document: https://help.aliyun.com/api/emr/describedatasourcecommand.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataSourceCommandWithCallback(request *DescribeDataSourceCommandRequest, callback func(response *DescribeDataSourceCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataSourceCommandResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataSourceCommand(request)
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

// DescribeDataSourceCommandRequest is the request struct for api DescribeDataSourceCommand
type DescribeDataSourceCommandRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Id              string           `position:"Query" name:"Id"`
}

// DescribeDataSourceCommandResponse is the response struct for api DescribeDataSourceCommand
type DescribeDataSourceCommandResponse struct {
	*responses.BaseResponse
	RequestId string                              `json:"RequestId" xml:"RequestId"`
	CommandId string                              `json:"CommandId" xml:"CommandId"`
	HostName  string                              `json:"HostName" xml:"HostName"`
	State     string                              `json:"State" xml:"State"`
	StartTime int                                 `json:"StartTime" xml:"StartTime"`
	EndTime   int                                 `json:"EndTime" xml:"EndTime"`
	Message   string                              `json:"Message" xml:"Message"`
	Data      string                              `json:"Data" xml:"Data"`
	HostList  HostListInDescribeDataSourceCommand `json:"HostList" xml:"HostList"`
}

// CreateDescribeDataSourceCommandRequest creates a request to invoke DescribeDataSourceCommand API
func CreateDescribeDataSourceCommandRequest() (request *DescribeDataSourceCommandRequest) {
	request = &DescribeDataSourceCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeDataSourceCommand", "emr", "openAPI")
	return
}

// CreateDescribeDataSourceCommandResponse creates a response to parse from DescribeDataSourceCommand response
func CreateDescribeDataSourceCommandResponse() (response *DescribeDataSourceCommandResponse) {
	response = &DescribeDataSourceCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

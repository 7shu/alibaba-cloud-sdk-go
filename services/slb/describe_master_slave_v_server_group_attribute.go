package slb

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

func (client *Client) DescribeMasterSlaveVServerGroupAttribute(request *DescribeMasterSlaveVServerGroupAttributeRequest) (response *DescribeMasterSlaveVServerGroupAttributeResponse, err error) {
	response = CreateDescribeMasterSlaveVServerGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeMasterSlaveVServerGroupAttributeWithChan(request *DescribeMasterSlaveVServerGroupAttributeRequest) (<-chan *DescribeMasterSlaveVServerGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeMasterSlaveVServerGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMasterSlaveVServerGroupAttribute(request)
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

func (client *Client) DescribeMasterSlaveVServerGroupAttributeWithCallback(request *DescribeMasterSlaveVServerGroupAttributeRequest, callback func(response *DescribeMasterSlaveVServerGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMasterSlaveVServerGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeMasterSlaveVServerGroupAttribute(request)
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

type DescribeMasterSlaveVServerGroupAttributeRequest struct {
	*requests.RpcRequest
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId               string           `position:"Query" name:"access_key_id"`
	Tags                      string           `position:"Query" name:"Tags"`
	MasterSlaveVServerGroupId string           `position:"Query" name:"MasterSlaveVServerGroupId"`
}

type DescribeMasterSlaveVServerGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId                   string                    `json:"RequestId" xml:"RequestId"`
	MasterSlaveVServerGroupId   string                    `json:"MasterSlaveVServerGroupId" xml:"MasterSlaveVServerGroupId"`
	MasterSlaveVServerGroupName string                    `json:"MasterSlaveVServerGroupName" xml:"MasterSlaveVServerGroupName"`
	MasterSlaveBackendServers   MasterSlaveBackendServers `json:"MasterSlaveBackendServers" xml:"MasterSlaveBackendServers"`
}

func CreateDescribeMasterSlaveVServerGroupAttributeRequest() (request *DescribeMasterSlaveVServerGroupAttributeRequest) {
	request = &DescribeMasterSlaveVServerGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveVServerGroupAttribute", "", "")
	return
}

func CreateDescribeMasterSlaveVServerGroupAttributeResponse() (response *DescribeMasterSlaveVServerGroupAttributeResponse) {
	response = &DescribeMasterSlaveVServerGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

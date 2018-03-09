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

// DeleteMasterSlaveVServerGroup invokes the slb.DeleteMasterSlaveVServerGroup API synchronously
// api document: https://help.aliyun.com/api/slb/deletemasterslavevservergroup.html
func (client *Client) DeleteMasterSlaveVServerGroup(request *DeleteMasterSlaveVServerGroupRequest) (response *DeleteMasterSlaveVServerGroupResponse, err error) {
	response = CreateDeleteMasterSlaveVServerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMasterSlaveVServerGroupWithChan invokes the slb.DeleteMasterSlaveVServerGroup API asynchronously
// api document: https://help.aliyun.com/api/slb/deletemasterslavevservergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMasterSlaveVServerGroupWithChan(request *DeleteMasterSlaveVServerGroupRequest) (<-chan *DeleteMasterSlaveVServerGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteMasterSlaveVServerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMasterSlaveVServerGroup(request)
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

// DeleteMasterSlaveVServerGroupWithCallback invokes the slb.DeleteMasterSlaveVServerGroup API asynchronously
// api document: https://help.aliyun.com/api/slb/deletemasterslavevservergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMasterSlaveVServerGroupWithCallback(request *DeleteMasterSlaveVServerGroupRequest, callback func(response *DeleteMasterSlaveVServerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMasterSlaveVServerGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteMasterSlaveVServerGroup(request)
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

// DeleteMasterSlaveVServerGroupRequest is the request struct for api DeleteMasterSlaveVServerGroup
type DeleteMasterSlaveVServerGroupRequest struct {
	*requests.RpcRequest
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId               string           `position:"Query" name:"access_key_id"`
	Tags                      string           `position:"Query" name:"Tags"`
	MasterSlaveVServerGroupId string           `position:"Query" name:"MasterSlaveVServerGroupId"`
}

// DeleteMasterSlaveVServerGroupResponse is the response struct for api DeleteMasterSlaveVServerGroup
type DeleteMasterSlaveVServerGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteMasterSlaveVServerGroupRequest creates a request to invoke DeleteMasterSlaveVServerGroup API
func CreateDeleteMasterSlaveVServerGroupRequest(request *DeleteMasterSlaveVServerGroupRequest) {
	request = &DeleteMasterSlaveVServerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteMasterSlaveVServerGroup", "slb", "openAPI")
	return
}

// CreateDeleteMasterSlaveVServerGroupResponse creates a response to parse from DeleteMasterSlaveVServerGroup response
func CreateDeleteMasterSlaveVServerGroupResponse() (response *DeleteMasterSlaveVServerGroupResponse) {
	response = &DeleteMasterSlaveVServerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

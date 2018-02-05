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

func (client *Client) CreateRouteEntry(request *CreateRouteEntryRequest) (response *CreateRouteEntryResponse, err error) {
	response = CreateCreateRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateRouteEntryWithChan(request *CreateRouteEntryRequest) (<-chan *CreateRouteEntryResponse, <-chan error) {
	responseChan := make(chan *CreateRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRouteEntry(request)
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

func (client *Client) CreateRouteEntryWithCallback(request *CreateRouteEntryRequest, callback func(response *CreateRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateRouteEntry(request)
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

type CreateRouteEntryRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer               `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string                         `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	RouteTableId         string                         `position:"Query" name:"RouteTableId"`
	DestinationCidrBlock string                         `position:"Query" name:"DestinationCidrBlock"`
	NextHopId            string                         `position:"Query" name:"NextHopId"`
	ClientToken          string                         `position:"Query" name:"ClientToken"`
	NextHopType          string                         `position:"Query" name:"NextHopType"`
	OwnerAccount         string                         `position:"Query" name:"OwnerAccount"`
	NextHopList          *[]CreateRouteEntryNextHopList `position:"Query" name:"NextHopList"  type:"Repeated"`
}

type CreateRouteEntryNextHopList struct {
	NextHopType string `name:"NextHopType"`
	NextHopId   string `name:"NextHopId"`
}

type CreateRouteEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateRouteEntryRequest() (request *CreateRouteEntryRequest) {
	request = &CreateRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateRouteEntry", "", "")
	return
}

func CreateCreateRouteEntryResponse() (response *CreateRouteEntryResponse) {
	response = &CreateRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

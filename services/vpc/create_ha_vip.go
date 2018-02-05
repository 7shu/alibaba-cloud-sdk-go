package vpc

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

func (client *Client) CreateHaVip(request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
	response = CreateCreateHaVipResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateHaVipWithChan(request *CreateHaVipRequest) (<-chan *CreateHaVipResponse, <-chan error) {
	responseChan := make(chan *CreateHaVipResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHaVip(request)
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

func (client *Client) CreateHaVipWithCallback(request *CreateHaVipRequest, callback func(response *CreateHaVipResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHaVipResponse
		var err error
		defer close(result)
		response, err = client.CreateHaVip(request)
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

type CreateHaVipRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	IpAddress            string           `position:"Query" name:"IpAddress"`
	Description          string           `position:"Query" name:"Description"`
}

type CreateHaVipResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HaVipId   string `json:"HaVipId" xml:"HaVipId"`
}

func CreateCreateHaVipRequest() (request *CreateHaVipRequest) {
	request = &CreateHaVipRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateHaVip", "", "")
	return
}

func CreateCreateHaVipResponse() (response *CreateHaVipResponse) {
	response = &CreateHaVipResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

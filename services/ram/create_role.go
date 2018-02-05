package ram

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

func (client *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
	response = CreateCreateRoleResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateRoleWithChan(request *CreateRoleRequest) (<-chan *CreateRoleResponse, <-chan error) {
	responseChan := make(chan *CreateRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRole(request)
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

func (client *Client) CreateRoleWithCallback(request *CreateRoleRequest, callback func(response *CreateRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRoleResponse
		var err error
		defer close(result)
		response, err = client.CreateRole(request)
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

type CreateRoleRequest struct {
	*requests.RpcRequest
	RoleName                 string `position:"Query" name:"RoleName"`
	Description              string `position:"Query" name:"Description"`
	AssumeRolePolicyDocument string `position:"Query" name:"AssumeRolePolicyDocument"`
}

type CreateRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Role      Role   `json:"Role" xml:"Role"`
}

func CreateCreateRoleRequest() (request *CreateRoleRequest) {
	request = &CreateRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "CreateRole", "", "")
	return
}

func CreateCreateRoleResponse() (response *CreateRoleResponse) {
	response = &CreateRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

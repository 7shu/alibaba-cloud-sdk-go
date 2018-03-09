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

// CreateReadOnlyDBInstance invokes the rds.CreateReadOnlyDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/createreadonlydbinstance.html
func (client *Client) CreateReadOnlyDBInstance(request *CreateReadOnlyDBInstanceRequest) (response *CreateReadOnlyDBInstanceResponse, err error) {
	response = CreateCreateReadOnlyDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateReadOnlyDBInstanceWithChan invokes the rds.CreateReadOnlyDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createreadonlydbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateReadOnlyDBInstanceWithChan(request *CreateReadOnlyDBInstanceRequest) (<-chan *CreateReadOnlyDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateReadOnlyDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateReadOnlyDBInstance(request)
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

// CreateReadOnlyDBInstanceWithCallback invokes the rds.CreateReadOnlyDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createreadonlydbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateReadOnlyDBInstanceWithCallback(request *CreateReadOnlyDBInstanceRequest, callback func(response *CreateReadOnlyDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateReadOnlyDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateReadOnlyDBInstance(request)
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

// CreateReadOnlyDBInstanceRequest is the request struct for api CreateReadOnlyDBInstance
type CreateReadOnlyDBInstanceRequest struct {
	*requests.RpcRequest
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	DBInstanceId          string           `position:"Query" name:"DBInstanceId"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
	EngineVersion         string           `position:"Query" name:"EngineVersion"`
	PayType               string           `position:"Query" name:"PayType"`
	DBInstanceDescription string           `position:"Query" name:"DBInstanceDescription"`
	InstanceNetworkType   string           `position:"Query" name:"InstanceNetworkType"`
	VPCId                 string           `position:"Query" name:"VPCId"`
	VSwitchId             string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string           `position:"Query" name:"PrivateIpAddress"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
}

// CreateReadOnlyDBInstanceResponse is the response struct for api CreateReadOnlyDBInstance
type CreateReadOnlyDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId          string `json:"OrderId" xml:"OrderId"`
	ConnectionString string `json:"ConnectionString" xml:"ConnectionString"`
	Port             string `json:"Port" xml:"Port"`
}

// CreateCreateReadOnlyDBInstanceRequest creates a request to invoke CreateReadOnlyDBInstance API
func CreateCreateReadOnlyDBInstanceRequest(request *CreateReadOnlyDBInstanceRequest) {
	request = &CreateReadOnlyDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateReadOnlyDBInstance", "rds", "openAPI")
	return
}

// CreateCreateReadOnlyDBInstanceResponse creates a response to parse from CreateReadOnlyDBInstance response
func CreateCreateReadOnlyDBInstanceResponse() (response *CreateReadOnlyDBInstanceResponse) {
	response = &CreateReadOnlyDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

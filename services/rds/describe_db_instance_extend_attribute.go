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

func (client *Client) DescribeDBInstanceExtendAttribute(request *DescribeDBInstanceExtendAttributeRequest) (response *DescribeDBInstanceExtendAttributeResponse, err error) {
	response = CreateDescribeDBInstanceExtendAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDBInstanceExtendAttributeWithChan(request *DescribeDBInstanceExtendAttributeRequest) (<-chan *DescribeDBInstanceExtendAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceExtendAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceExtendAttribute(request)
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

func (client *Client) DescribeDBInstanceExtendAttributeWithCallback(request *DescribeDBInstanceExtendAttributeRequest, callback func(response *DescribeDBInstanceExtendAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceExtendAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceExtendAttribute(request)
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

type DescribeDBInstanceExtendAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

type DescribeDBInstanceExtendAttributeResponse struct {
	*responses.BaseResponse
	RequestId                         string `json:"RequestId" xml:"RequestId"`
	CanTempUpgrade                    bool   `json:"CanTempUpgrade" xml:"CanTempUpgrade"`
	TempUpgradeTimeStart              string `json:"TempUpgradeTimeStart" xml:"TempUpgradeTimeStart"`
	TempUpgradeTimeEnd                string `json:"TempUpgradeTimeEnd" xml:"TempUpgradeTimeEnd"`
	TempUpgradeRecoveryTime           string `json:"TempUpgradeRecoveryTime" xml:"TempUpgradeRecoveryTime"`
	TempUpgradeRecoveryClass          string `json:"TempUpgradeRecoveryClass" xml:"TempUpgradeRecoveryClass"`
	TempUpgradeRecoveryCpu            int    `json:"TempUpgradeRecoveryCpu" xml:"TempUpgradeRecoveryCpu"`
	TempUpgradeRecoveryMemory         int    `json:"TempUpgradeRecoveryMemory" xml:"TempUpgradeRecoveryMemory"`
	TempUpgradeRecoveryMaxIOPS        string `json:"TempUpgradeRecoveryMaxIOPS" xml:"TempUpgradeRecoveryMaxIOPS"`
	TempUpgradeRecoveryMaxConnections string `json:"TempUpgradeRecoveryMaxConnections" xml:"TempUpgradeRecoveryMaxConnections"`
}

func CreateDescribeDBInstanceExtendAttributeRequest() (request *DescribeDBInstanceExtendAttributeRequest) {
	request = &DescribeDBInstanceExtendAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceExtendAttribute", "", "")
	return
}

func CreateDescribeDBInstanceExtendAttributeResponse() (response *DescribeDBInstanceExtendAttributeResponse) {
	response = &DescribeDBInstanceExtendAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

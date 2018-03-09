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

// PreCheckBeforeImportData invokes the rds.PreCheckBeforeImportData API synchronously
// api document: https://help.aliyun.com/api/rds/precheckbeforeimportdata.html
func (client *Client) PreCheckBeforeImportData(request *PreCheckBeforeImportDataRequest) (response *PreCheckBeforeImportDataResponse, err error) {
	response = CreatePreCheckBeforeImportDataResponse()
	err = client.DoAction(request, response)
	return
}

// PreCheckBeforeImportDataWithChan invokes the rds.PreCheckBeforeImportData API asynchronously
// api document: https://help.aliyun.com/api/rds/precheckbeforeimportdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreCheckBeforeImportDataWithChan(request *PreCheckBeforeImportDataRequest) (<-chan *PreCheckBeforeImportDataResponse, <-chan error) {
	responseChan := make(chan *PreCheckBeforeImportDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreCheckBeforeImportData(request)
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

// PreCheckBeforeImportDataWithCallback invokes the rds.PreCheckBeforeImportData API asynchronously
// api document: https://help.aliyun.com/api/rds/precheckbeforeimportdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreCheckBeforeImportDataWithCallback(request *PreCheckBeforeImportDataRequest, callback func(response *PreCheckBeforeImportDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreCheckBeforeImportDataResponse
		var err error
		defer close(result)
		response, err = client.PreCheckBeforeImportData(request)
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

// PreCheckBeforeImportDataRequest is the request struct for api PreCheckBeforeImportData
type PreCheckBeforeImportDataRequest struct {
	*requests.RpcRequest
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken            string           `position:"Query" name:"ClientToken"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	SourceDatabaseIp       string           `position:"Query" name:"SourceDatabaseIp"`
	SourceDatabasePort     string           `position:"Query" name:"SourceDatabasePort"`
	SourceDatabaseUserName string           `position:"Query" name:"SourceDatabaseUserName"`
	SourceDatabasePassword string           `position:"Query" name:"SourceDatabasePassword"`
	ImportDataType         string           `position:"Query" name:"ImportDataType"`
	SourceDatabaseDBNames  string           `position:"Query" name:"SourceDatabaseDBNames"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
}

// PreCheckBeforeImportDataResponse is the response struct for api PreCheckBeforeImportData
type PreCheckBeforeImportDataResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PreCheckId string `json:"PreCheckId" xml:"PreCheckId"`
}

// CreatePreCheckBeforeImportDataRequest creates a request to invoke PreCheckBeforeImportData API
func CreatePreCheckBeforeImportDataRequest(request *PreCheckBeforeImportDataRequest) {
	request = &PreCheckBeforeImportDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "PreCheckBeforeImportData", "rds", "openAPI")
	return
}

// CreatePreCheckBeforeImportDataResponse creates a response to parse from PreCheckBeforeImportData response
func CreatePreCheckBeforeImportDataResponse() (response *PreCheckBeforeImportDataResponse) {
	response = &PreCheckBeforeImportDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

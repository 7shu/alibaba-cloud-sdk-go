package cdn

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

// DeleteLiveAppRecordConfig invokes the cdn.DeleteLiveAppRecordConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/deleteliveapprecordconfig.html
func (client *Client) DeleteLiveAppRecordConfig(request *DeleteLiveAppRecordConfigRequest) (response *DeleteLiveAppRecordConfigResponse, err error) {
	response = CreateDeleteLiveAppRecordConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveAppRecordConfigWithChan invokes the cdn.DeleteLiveAppRecordConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/deleteliveapprecordconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveAppRecordConfigWithChan(request *DeleteLiveAppRecordConfigRequest) (<-chan *DeleteLiveAppRecordConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveAppRecordConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveAppRecordConfig(request)
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

// DeleteLiveAppRecordConfigWithCallback invokes the cdn.DeleteLiveAppRecordConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/deleteliveapprecordconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveAppRecordConfigWithCallback(request *DeleteLiveAppRecordConfigRequest, callback func(response *DeleteLiveAppRecordConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveAppRecordConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveAppRecordConfig(request)
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

// DeleteLiveAppRecordConfigRequest is the request struct for api DeleteLiveAppRecordConfig
type DeleteLiveAppRecordConfigRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
}

// DeleteLiveAppRecordConfigResponse is the response struct for api DeleteLiveAppRecordConfig
type DeleteLiveAppRecordConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveAppRecordConfigRequest creates a request to invoke DeleteLiveAppRecordConfig API
func CreateDeleteLiveAppRecordConfigRequest(request *DeleteLiveAppRecordConfigRequest) {
	request = &DeleteLiveAppRecordConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLiveAppRecordConfig", "", "")
	return
}

// CreateDeleteLiveAppRecordConfigResponse creates a response to parse from DeleteLiveAppRecordConfig response
func CreateDeleteLiveAppRecordConfigResponse() (response *DeleteLiveAppRecordConfigResponse) {
	response = &DeleteLiveAppRecordConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

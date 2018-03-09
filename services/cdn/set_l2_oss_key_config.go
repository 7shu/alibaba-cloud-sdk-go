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

// SetL2OssKeyConfig invokes the cdn.SetL2OssKeyConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setl2osskeyconfig.html
func (client *Client) SetL2OssKeyConfig(request *SetL2OssKeyConfigRequest) (response *SetL2OssKeyConfigResponse, err error) {
	response = CreateSetL2OssKeyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetL2OssKeyConfigWithChan invokes the cdn.SetL2OssKeyConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setl2osskeyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetL2OssKeyConfigWithChan(request *SetL2OssKeyConfigRequest) (<-chan *SetL2OssKeyConfigResponse, <-chan error) {
	responseChan := make(chan *SetL2OssKeyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetL2OssKeyConfig(request)
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

// SetL2OssKeyConfigWithCallback invokes the cdn.SetL2OssKeyConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setl2osskeyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetL2OssKeyConfigWithCallback(request *SetL2OssKeyConfigRequest, callback func(response *SetL2OssKeyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetL2OssKeyConfigResponse
		var err error
		defer close(result)
		response, err = client.SetL2OssKeyConfig(request)
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

// SetL2OssKeyConfigRequest is the request struct for api SetL2OssKeyConfig
type SetL2OssKeyConfigRequest struct {
	*requests.RpcRequest
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	DomainName     string           `position:"Query" name:"DomainName"`
	PrivateOssAuth string           `position:"Query" name:"PrivateOssAuth"`
}

// SetL2OssKeyConfigResponse is the response struct for api SetL2OssKeyConfig
type SetL2OssKeyConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetL2OssKeyConfigRequest creates a request to invoke SetL2OssKeyConfig API
func CreateSetL2OssKeyConfigRequest(request *SetL2OssKeyConfigRequest) {
	request = &SetL2OssKeyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetL2OssKeyConfig", "", "")
	return
}

// CreateSetL2OssKeyConfigResponse creates a response to parse from SetL2OssKeyConfig response
func CreateSetL2OssKeyConfigResponse() (response *SetL2OssKeyConfigResponse) {
	response = &SetL2OssKeyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

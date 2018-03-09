package mts

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

// UpdateTerrorismPipeline invokes the mts.UpdateTerrorismPipeline API synchronously
// api document: https://help.aliyun.com/api/mts/updateterrorismpipeline.html
func (client *Client) UpdateTerrorismPipeline(request *UpdateTerrorismPipelineRequest) (response *UpdateTerrorismPipelineResponse, err error) {
	response = CreateUpdateTerrorismPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTerrorismPipelineWithChan invokes the mts.UpdateTerrorismPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/updateterrorismpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateTerrorismPipelineWithChan(request *UpdateTerrorismPipelineRequest) (<-chan *UpdateTerrorismPipelineResponse, <-chan error) {
	responseChan := make(chan *UpdateTerrorismPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTerrorismPipeline(request)
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

// UpdateTerrorismPipelineWithCallback invokes the mts.UpdateTerrorismPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/updateterrorismpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateTerrorismPipelineWithCallback(request *UpdateTerrorismPipelineRequest, callback func(response *UpdateTerrorismPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTerrorismPipelineResponse
		var err error
		defer close(result)
		response, err = client.UpdateTerrorismPipeline(request)
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

// UpdateTerrorismPipelineRequest is the request struct for api UpdateTerrorismPipeline
type UpdateTerrorismPipelineRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Name                 string           `position:"Query" name:"Name"`
	State                string           `position:"Query" name:"State"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// UpdateTerrorismPipelineResponse is the response struct for api UpdateTerrorismPipeline
type UpdateTerrorismPipelineResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Pipeline  Pipeline `json:"Pipeline" xml:"Pipeline"`
}

// CreateUpdateTerrorismPipelineRequest creates a request to invoke UpdateTerrorismPipeline API
func CreateUpdateTerrorismPipelineRequest(request *UpdateTerrorismPipelineRequest) {
	request = &UpdateTerrorismPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateTerrorismPipeline", "mts", "openAPI")
	return
}

// CreateUpdateTerrorismPipelineResponse creates a response to parse from UpdateTerrorismPipeline response
func CreateUpdateTerrorismPipelineResponse() (response *UpdateTerrorismPipelineResponse) {
	response = &UpdateTerrorismPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

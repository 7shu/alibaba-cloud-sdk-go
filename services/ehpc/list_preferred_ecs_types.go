package ehpc

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

// ListPreferredEcsTypes invokes the ehpc.ListPreferredEcsTypes API synchronously
// api document: https://help.aliyun.com/api/ehpc/listpreferredecstypes.html
func (client *Client) ListPreferredEcsTypes(request *ListPreferredEcsTypesRequest) (response *ListPreferredEcsTypesResponse, err error) {
	response = CreateListPreferredEcsTypesResponse()
	err = client.DoAction(request, response)
	return
}

// ListPreferredEcsTypesWithChan invokes the ehpc.ListPreferredEcsTypes API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listpreferredecstypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPreferredEcsTypesWithChan(request *ListPreferredEcsTypesRequest) (<-chan *ListPreferredEcsTypesResponse, <-chan error) {
	responseChan := make(chan *ListPreferredEcsTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPreferredEcsTypes(request)
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

// ListPreferredEcsTypesWithCallback invokes the ehpc.ListPreferredEcsTypes API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listpreferredecstypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPreferredEcsTypesWithCallback(request *ListPreferredEcsTypesRequest, callback func(response *ListPreferredEcsTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPreferredEcsTypesResponse
		var err error
		defer close(result)
		response, err = client.ListPreferredEcsTypes(request)
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

// ListPreferredEcsTypesRequest is the request struct for api ListPreferredEcsTypes
type ListPreferredEcsTypesRequest struct {
	*requests.RpcRequest
	ZoneId       string `position:"Query" name:"ZoneId"`
	SpotStrategy string `position:"Query" name:"SpotStrategy"`
}

// ListPreferredEcsTypesResponse is the response struct for api ListPreferredEcsTypes
type ListPreferredEcsTypesResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	SupportSpotInstance bool   `json:"SupportSpotInstance" xml:"SupportSpotInstance"`
	Series              Series `json:"Series" xml:"Series"`
}

// CreateListPreferredEcsTypesRequest creates a request to invoke ListPreferredEcsTypes API
func CreateListPreferredEcsTypesRequest(request *ListPreferredEcsTypesRequest) {
	request = &ListPreferredEcsTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "ListPreferredEcsTypes", "ehs", "openAPI")
	return
}

// CreateListPreferredEcsTypesResponse creates a response to parse from ListPreferredEcsTypes response
func CreateListPreferredEcsTypesResponse() (response *ListPreferredEcsTypesResponse) {
	response = &ListPreferredEcsTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

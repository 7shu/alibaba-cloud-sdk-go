package cloudphoto

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

// GetFramedPhotoUrls invokes the cloudphoto.GetFramedPhotoUrls API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/getframedphotourls.html
func (client *Client) GetFramedPhotoUrls(request *GetFramedPhotoUrlsRequest) (response *GetFramedPhotoUrlsResponse, err error) {
	response = CreateGetFramedPhotoUrlsResponse()
	err = client.DoAction(request, response)
	return
}

// GetFramedPhotoUrlsWithChan invokes the cloudphoto.GetFramedPhotoUrls API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getframedphotourls.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFramedPhotoUrlsWithChan(request *GetFramedPhotoUrlsRequest) (<-chan *GetFramedPhotoUrlsResponse, <-chan error) {
	responseChan := make(chan *GetFramedPhotoUrlsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFramedPhotoUrls(request)
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

// GetFramedPhotoUrlsWithCallback invokes the cloudphoto.GetFramedPhotoUrls API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getframedphotourls.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFramedPhotoUrlsWithCallback(request *GetFramedPhotoUrlsRequest, callback func(response *GetFramedPhotoUrlsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFramedPhotoUrlsResponse
		var err error
		defer close(result)
		response, err = client.GetFramedPhotoUrls(request)
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

// GetFramedPhotoUrlsRequest is the request struct for api GetFramedPhotoUrls
type GetFramedPhotoUrlsRequest struct {
	*requests.RpcRequest
	PhotoId   *[]string `position:"Query" name:"PhotoId"  type:"Repeated"`
	FrameId   string    `position:"Query" name:"FrameId"`
	StoreName string    `position:"Query" name:"StoreName"`
	LibraryId string    `position:"Query" name:"LibraryId"`
}

// GetFramedPhotoUrlsResponse is the response struct for api GetFramedPhotoUrls
type GetFramedPhotoUrlsResponse struct {
	*responses.BaseResponse
	Code      string                      `json:"Code" xml:"Code"`
	Message   string                      `json:"Message" xml:"Message"`
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Action    string                      `json:"Action" xml:"Action"`
	Results   ResultsInGetFramedPhotoUrls `json:"Results" xml:"Results"`
}

// CreateGetFramedPhotoUrlsRequest creates a request to invoke GetFramedPhotoUrls API
func CreateGetFramedPhotoUrlsRequest(request *GetFramedPhotoUrlsRequest) {
	request = &GetFramedPhotoUrlsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetFramedPhotoUrls", "cloudphoto", "openAPI")
	return
}

// CreateGetFramedPhotoUrlsResponse creates a response to parse from GetFramedPhotoUrls response
func CreateGetFramedPhotoUrlsResponse() (response *GetFramedPhotoUrlsResponse) {
	response = &GetFramedPhotoUrlsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

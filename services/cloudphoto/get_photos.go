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

// GetPhotos invokes the cloudphoto.GetPhotos API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/getphotos.html
func (client *Client) GetPhotos(request *GetPhotosRequest) (response *GetPhotosResponse, err error) {
	response = CreateGetPhotosResponse()
	err = client.DoAction(request, response)
	return
}

// GetPhotosWithChan invokes the cloudphoto.GetPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPhotosWithChan(request *GetPhotosRequest) (<-chan *GetPhotosResponse, <-chan error) {
	responseChan := make(chan *GetPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPhotos(request)
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

// GetPhotosWithCallback invokes the cloudphoto.GetPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPhotosWithCallback(request *GetPhotosRequest, callback func(response *GetPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPhotosResponse
		var err error
		defer close(result)
		response, err = client.GetPhotos(request)
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

// GetPhotosRequest is the request struct for api GetPhotos
type GetPhotosRequest struct {
	*requests.RpcRequest
	PhotoId   *[]string `position:"Query" name:"PhotoId"  type:"Repeated"`
	StoreName string    `position:"Query" name:"StoreName"`
	LibraryId string    `position:"Query" name:"LibraryId"`
}

// GetPhotosResponse is the response struct for api GetPhotos
type GetPhotosResponse struct {
	*responses.BaseResponse
	Code      string  `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Action    string  `json:"Action" xml:"Action"`
	Photos    []Photo `json:"Photos" xml:"Photos"`
}

// CreateGetPhotosRequest creates a request to invoke GetPhotos API
func CreateGetPhotosRequest(request *GetPhotosRequest) {
	request = &GetPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotos", "cloudphoto", "openAPI")
	return
}

// CreateGetPhotosResponse creates a response to parse from GetPhotos response
func CreateGetPhotosResponse() (response *GetPhotosResponse) {
	response = &GetPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

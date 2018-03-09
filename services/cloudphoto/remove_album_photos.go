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

// RemoveAlbumPhotos invokes the cloudphoto.RemoveAlbumPhotos API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/removealbumphotos.html
func (client *Client) RemoveAlbumPhotos(request *RemoveAlbumPhotosRequest) (response *RemoveAlbumPhotosResponse, err error) {
	response = CreateRemoveAlbumPhotosResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveAlbumPhotosWithChan invokes the cloudphoto.RemoveAlbumPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/removealbumphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveAlbumPhotosWithChan(request *RemoveAlbumPhotosRequest) (<-chan *RemoveAlbumPhotosResponse, <-chan error) {
	responseChan := make(chan *RemoveAlbumPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveAlbumPhotos(request)
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

// RemoveAlbumPhotosWithCallback invokes the cloudphoto.RemoveAlbumPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/removealbumphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveAlbumPhotosWithCallback(request *RemoveAlbumPhotosRequest, callback func(response *RemoveAlbumPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveAlbumPhotosResponse
		var err error
		defer close(result)
		response, err = client.RemoveAlbumPhotos(request)
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

// RemoveAlbumPhotosRequest is the request struct for api RemoveAlbumPhotos
type RemoveAlbumPhotosRequest struct {
	*requests.RpcRequest
	AlbumId   requests.Integer `position:"Query" name:"AlbumId"`
	PhotoId   *[]string        `position:"Query" name:"PhotoId"  type:"Repeated"`
	StoreName string           `position:"Query" name:"StoreName"`
	LibraryId string           `position:"Query" name:"LibraryId"`
}

// RemoveAlbumPhotosResponse is the response struct for api RemoveAlbumPhotos
type RemoveAlbumPhotosResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Action    string   `json:"Action" xml:"Action"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateRemoveAlbumPhotosRequest creates a request to invoke RemoveAlbumPhotos API
func CreateRemoveAlbumPhotosRequest(request *RemoveAlbumPhotosRequest) {
	request = &RemoveAlbumPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "RemoveAlbumPhotos", "cloudphoto", "openAPI")
	return
}

// CreateRemoveAlbumPhotosResponse creates a response to parse from RemoveAlbumPhotos response
func CreateRemoveAlbumPhotosResponse() (response *RemoveAlbumPhotosResponse) {
	response = &RemoveAlbumPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

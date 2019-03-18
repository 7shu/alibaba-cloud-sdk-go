package cas

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

// CreateSignaturePeopleCertificate invokes the cas.CreateSignaturePeopleCertificate API synchronously
// api document: https://help.aliyun.com/api/cas/createsignaturepeoplecertificate.html
func (client *Client) CreateSignaturePeopleCertificate(request *CreateSignaturePeopleCertificateRequest) (response *CreateSignaturePeopleCertificateResponse, err error) {
	response = CreateCreateSignaturePeopleCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSignaturePeopleCertificateWithChan invokes the cas.CreateSignaturePeopleCertificate API asynchronously
// api document: https://help.aliyun.com/api/cas/createsignaturepeoplecertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSignaturePeopleCertificateWithChan(request *CreateSignaturePeopleCertificateRequest) (<-chan *CreateSignaturePeopleCertificateResponse, <-chan error) {
	responseChan := make(chan *CreateSignaturePeopleCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSignaturePeopleCertificate(request)
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

// CreateSignaturePeopleCertificateWithCallback invokes the cas.CreateSignaturePeopleCertificate API asynchronously
// api document: https://help.aliyun.com/api/cas/createsignaturepeoplecertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSignaturePeopleCertificateWithCallback(request *CreateSignaturePeopleCertificateRequest, callback func(response *CreateSignaturePeopleCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSignaturePeopleCertificateResponse
		var err error
		defer close(result)
		response, err = client.CreateSignaturePeopleCertificate(request)
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

// CreateSignaturePeopleCertificateRequest is the request struct for api CreateSignaturePeopleCertificate
type CreateSignaturePeopleCertificateRequest struct {
	*requests.RpcRequest
	PeopleName     string `position:"Query" name:"PeopleName"`
	SourceIp       string `position:"Query" name:"SourceIp"`
	Mobile         string `position:"Query" name:"Mobile"`
	IdentityNumber string `position:"Query" name:"IdentityNumber"`
	Lang           string `position:"Query" name:"Lang"`
	Email          string `position:"Query" name:"Email"`
}

// CreateSignaturePeopleCertificateResponse is the response struct for api CreateSignaturePeopleCertificate
type CreateSignaturePeopleCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateSignaturePeopleCertificateRequest creates a request to invoke CreateSignaturePeopleCertificate API
func CreateCreateSignaturePeopleCertificateRequest() (request *CreateSignaturePeopleCertificateRequest) {
	request = &CreateSignaturePeopleCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2018-08-13", "CreateSignaturePeopleCertificate", "cas_esign_fdd", "openAPI")
	return
}

// CreateCreateSignaturePeopleCertificateResponse creates a response to parse from CreateSignaturePeopleCertificate response
func CreateCreateSignaturePeopleCertificateResponse() (response *CreateSignaturePeopleCertificateResponse) {
	response = &CreateSignaturePeopleCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

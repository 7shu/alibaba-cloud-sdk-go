package domain_intl

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

func (client *Client) SaveRegistrantProfile(request *SaveRegistrantProfileRequest) (response *SaveRegistrantProfileResponse, err error) {
	response = CreateSaveRegistrantProfileResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SaveRegistrantProfileWithChan(request *SaveRegistrantProfileRequest) (<-chan *SaveRegistrantProfileResponse, <-chan error) {
	responseChan := make(chan *SaveRegistrantProfileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveRegistrantProfile(request)
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

func (client *Client) SaveRegistrantProfileWithCallback(request *SaveRegistrantProfileRequest, callback func(response *SaveRegistrantProfileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveRegistrantProfileResponse
		var err error
		defer close(result)
		response, err = client.SaveRegistrantProfile(request)
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

type SaveRegistrantProfileRequest struct {
	*requests.RpcRequest
	Country                  string           `position:"Query" name:"Country"`
	UserClientIp             string           `position:"Query" name:"UserClientIp"`
	Lang                     string           `position:"Query" name:"Lang"`
	RegistrantProfileId      requests.Integer `position:"Query" name:"RegistrantProfileId"`
	DefaultRegistrantProfile requests.Boolean `position:"Query" name:"DefaultRegistrantProfile"`
	City                     string           `position:"Query" name:"City"`
	RegistrantOrganization   string           `position:"Query" name:"RegistrantOrganization"`
	RegistrantName           string           `position:"Query" name:"RegistrantName"`
	Province                 string           `position:"Query" name:"Province"`
	Address                  string           `position:"Query" name:"Address"`
	Email                    string           `position:"Query" name:"Email"`
	PostalCode               string           `position:"Query" name:"PostalCode"`
	TelArea                  string           `position:"Query" name:"TelArea"`
	Telephone                string           `position:"Query" name:"Telephone"`
	TelExt                   string           `position:"Query" name:"TelExt"`
}

type SaveRegistrantProfileResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	RegistrantProfileId int    `json:"RegistrantProfileId" xml:"RegistrantProfileId"`
}

func CreateSaveRegistrantProfileRequest() (request *SaveRegistrantProfileRequest) {
	request = &SaveRegistrantProfileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveRegistrantProfile", "", "")
	return
}

func CreateSaveRegistrantProfileResponse() (response *SaveRegistrantProfileResponse) {
	response = &SaveRegistrantProfileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

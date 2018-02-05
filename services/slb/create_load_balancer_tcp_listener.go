package slb

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

func (client *Client) CreateLoadBalancerTCPListener(request *CreateLoadBalancerTCPListenerRequest) (response *CreateLoadBalancerTCPListenerResponse, err error) {
	response = CreateCreateLoadBalancerTCPListenerResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateLoadBalancerTCPListenerWithChan(request *CreateLoadBalancerTCPListenerRequest) (<-chan *CreateLoadBalancerTCPListenerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerTCPListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancerTCPListener(request)
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

func (client *Client) CreateLoadBalancerTCPListenerWithCallback(request *CreateLoadBalancerTCPListenerRequest, callback func(response *CreateLoadBalancerTCPListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerTCPListenerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancerTCPListener(request)
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

type CreateLoadBalancerTCPListenerRequest struct {
	*requests.RpcRequest
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId            string           `position:"Query" name:"LoadBalancerId"`
	ListenerPort              requests.Integer `position:"Query" name:"ListenerPort"`
	BackendServerPort         requests.Integer `position:"Query" name:"BackendServerPort"`
	Bandwidth                 requests.Integer `position:"Query" name:"Bandwidth"`
	Scheduler                 string           `position:"Query" name:"Scheduler"`
	PersistenceTimeout        requests.Integer `position:"Query" name:"PersistenceTimeout"`
	EstablishedTimeout        requests.Integer `position:"Query" name:"EstablishedTimeout"`
	HealthyThreshold          requests.Integer `position:"Query" name:"HealthyThreshold"`
	UnhealthyThreshold        requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthCheckConnectTimeout requests.Integer `position:"Query" name:"HealthCheckConnectTimeout"`
	HealthCheckConnectPort    requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckInterval       requests.Integer `position:"Query" name:"healthCheckInterval"`
	HealthCheckDomain         string           `position:"Query" name:"HealthCheckDomain"`
	HealthCheckURI            string           `position:"Query" name:"HealthCheckURI"`
	HealthCheckHttpCode       string           `position:"Query" name:"HealthCheckHttpCode"`
	HealthCheckType           string           `position:"Query" name:"HealthCheckType"`
	MaxConnection             requests.Integer `position:"Query" name:"MaxConnection"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId               string           `position:"Query" name:"access_key_id"`
	VServerGroupId            string           `position:"Query" name:"VServerGroupId"`
	MasterSlaveServerGroupId  string           `position:"Query" name:"MasterSlaveServerGroupId"`
	Tags                      string           `position:"Query" name:"Tags"`
}

type CreateLoadBalancerTCPListenerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateLoadBalancerTCPListenerRequest() (request *CreateLoadBalancerTCPListenerRequest) {
	request = &CreateLoadBalancerTCPListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancerTCPListener", "", "")
	return
}

func CreateCreateLoadBalancerTCPListenerResponse() (response *CreateLoadBalancerTCPListenerResponse) {
	response = &CreateLoadBalancerTCPListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

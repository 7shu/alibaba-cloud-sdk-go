package ecs

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

// Invocation is a nested struct in ecs response
type Invocation struct {
	TotalCount        int               `json:"TotalCount" xml:"TotalCount"`
	Timed             bool              `json:"Timed" xml:"Timed"`
	CommandName       string            `json:"CommandName" xml:"CommandName"`
	PageSize          int               `json:"PageSize" xml:"PageSize"`
	PageNumber        int               `json:"PageNumber" xml:"PageNumber"`
	InvokeId          string            `json:"InvokeId" xml:"InvokeId"`
	InvokeStatus      string            `json:"InvokeStatus" xml:"InvokeStatus"`
	CommandType       string            `json:"CommandType" xml:"CommandType"`
	Frequency         string            `json:"Frequency" xml:"Frequency"`
	CommandId         string            `json:"CommandId" xml:"CommandId"`
	InvokeInstances   InvokeInstances   `json:"InvokeInstances" xml:"InvokeInstances"`
	InvocationResults InvocationResults `json:"InvocationResults" xml:"InvocationResults"`
}

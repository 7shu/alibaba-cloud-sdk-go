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

type Activity struct {
	Type             string           `json:"Type" xml:"Type"`
	Message          string           `json:"Message" xml:"Message"`
	Name             string           `json:"Name" xml:"Name"`
	StartTime        string           `json:"StartTime" xml:"StartTime"`
	State            string           `json:"State" xml:"State"`
	Code             string           `json:"Code" xml:"Code"`
	JobId            string           `json:"JobId" xml:"JobId"`
	EndTime          string           `json:"EndTime" xml:"EndTime"`
	MNSMessageResult MNSMessageResult `json:"MNSMessageResult" xml:"MNSMessageResult"`
}

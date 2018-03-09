package cdn

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

// DataModule is a nested struct in cdn response
type DataModule struct {
	AccDomesticValue     string  `json:"AccDomesticValue" xml:"AccDomesticValue"`
	L2Value              string  `json:"L2Value" xml:"L2Value"`
	OverseasL2Value      string  `json:"OverseasL2Value" xml:"OverseasL2Value"`
	DynamicOverseasValue string  `json:"DynamicOverseasValue" xml:"DynamicOverseasValue"`
	L1OutBps             float64 `json:"L1OutBps" xml:"L1OutBps"`
	StaticValue          string  `json:"StaticValue" xml:"StaticValue"`
	TimeStamp            string  `json:"TimeStamp" xml:"TimeStamp"`
	Value                string  `json:"Value" xml:"Value"`
	Timestamp            string  `json:"Timestamp" xml:"Timestamp"`
	L1InnerBps           float64 `json:"L1InnerBps" xml:"L1InnerBps"`
	StaticOverseasValue  string  `json:"StaticOverseasValue" xml:"StaticOverseasValue"`
	DynamicValue         string  `json:"DynamicValue" xml:"DynamicValue"`
	AccValue             string  `json:"AccValue" xml:"AccValue"`
	AccOverseasValue     string  `json:"AccOverseasValue" xml:"AccOverseasValue"`
	DomesticL2Value      string  `json:"DomesticL2Value" xml:"DomesticL2Value"`
	L1Bps                float64 `json:"L1Bps" xml:"L1Bps"`
	DomesticValue        string  `json:"DomesticValue" xml:"DomesticValue"`
	OverseasValue        string  `json:"OverseasValue" xml:"OverseasValue"`
	StaticDomesticValue  string  `json:"StaticDomesticValue" xml:"StaticDomesticValue"`
	DynamicDomesticValue string  `json:"DynamicDomesticValue" xml:"DynamicDomesticValue"`
	DomainName           string  `json:"DomainName" xml:"DomainName"`
}

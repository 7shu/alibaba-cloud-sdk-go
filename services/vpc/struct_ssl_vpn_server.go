package vpc

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

// SslVpnServer is a nested struct in vpc response
type SslVpnServer struct {
	RegionId       string `json:"RegionId" xml:"RegionId"`
	SslVpnServerId string `json:"SslVpnServerId" xml:"SslVpnServerId"`
	VpnGatewayId   string `json:"VpnGatewayId" xml:"VpnGatewayId"`
	Name           string `json:"Name" xml:"Name"`
	LocalSubnet    string `json:"LocalSubnet" xml:"LocalSubnet"`
	ClientIpPool   string `json:"ClientIpPool" xml:"ClientIpPool"`
	CreateTime     int    `json:"CreateTime" xml:"CreateTime"`
	Cipher         string `json:"Cipher" xml:"Cipher"`
	Proto          string `json:"Proto" xml:"Proto"`
	Port           int    `json:"Port" xml:"Port"`
	Compress       bool   `json:"Compress" xml:"Compress"`
	Connections    int    `json:"Connections" xml:"Connections"`
	MaxConnections int    `json:"MaxConnections" xml:"MaxConnections"`
	InternetIp     string `json:"InternetIp" xml:"InternetIp"`
}

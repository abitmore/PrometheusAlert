package smartag

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

// Acr is a nested struct in smartag response
type Acr struct {
	AcrId           string `json:"AcrId" xml:"AcrId"`
	Description     string `json:"Description" xml:"Description"`
	Direction       string `json:"Direction" xml:"Direction"`
	SourceCidr      string `json:"SourceCidr" xml:"SourceCidr"`
	DestCidr        string `json:"DestCidr" xml:"DestCidr"`
	IpProtocol      string `json:"IpProtocol" xml:"IpProtocol"`
	SourcePortRange string `json:"SourcePortRange" xml:"SourcePortRange"`
	DestPortRange   string `json:"DestPortRange" xml:"DestPortRange"`
	Policy          string `json:"Policy" xml:"Policy"`
	Priority        int    `json:"Priority" xml:"Priority"`
	GmtCreate       int64  `json:"GmtCreate" xml:"GmtCreate"`
	AclId           string `json:"AclId" xml:"AclId"`
	Type            string `json:"Type" xml:"Type"`
}
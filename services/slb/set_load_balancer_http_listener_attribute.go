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

func (client *Client) SetLoadBalancerHTTPListenerAttribute(request *SetLoadBalancerHTTPListenerAttributeRequest) (response *SetLoadBalancerHTTPListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerHTTPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetLoadBalancerHTTPListenerAttributeWithChan(request *SetLoadBalancerHTTPListenerAttributeRequest) (<-chan *SetLoadBalancerHTTPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerHTTPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerHTTPListenerAttribute(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) SetLoadBalancerHTTPListenerAttributeWithCallback(request *SetLoadBalancerHTTPListenerAttributeRequest, callback func(response *SetLoadBalancerHTTPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerHTTPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerHTTPListenerAttribute(request)
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

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	*requests.RpcRequest
	VServerGroup           string           `position:"Query" name:"VServerGroup"`
	XForwardedFor          string           `position:"Query" name:"XForwardedFor"`
	UnhealthyThreshold     requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	Bandwidth              requests.Integer `position:"Query" name:"Bandwidth"`
	HealthCheck            string           `position:"Query" name:"HealthCheck"`
	XForwardedForSLBIP     string           `position:"Query" name:"XForwardedFor_SLBIP"`
	HealthCheckDomain      string           `position:"Query" name:"HealthCheckDomain"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	StickySession          string           `position:"Query" name:"StickySession"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	XForwardedForSLBID     string           `position:"Query" name:"XForwardedFor_SLBID"`
	Tags                   string           `position:"Query" name:"Tags"`
	HealthCheckTimeout     requests.Integer `position:"Query" name:"HealthCheckTimeout"`
	HealthCheckHttpCode    string           `position:"Query" name:"HealthCheckHttpCode"`
	Gzip                   string           `position:"Query" name:"Gzip"`
	Scheduler              string           `position:"Query" name:"Scheduler"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	XForwardedForProto     string           `position:"Query" name:"XForwardedFor_proto"`
	VServerGroupId         string           `position:"Query" name:"VServerGroupId"`
	Cookie                 string           `position:"Query" name:"Cookie"`
	HealthCheckInterval    requests.Integer `position:"Query" name:"HealthCheckInterval"`
	ListenerPort           requests.Integer `position:"Query" name:"ListenerPort"`
	HealthCheckURI         string           `position:"Query" name:"HealthCheckURI"`
	AccessKeyId            string           `position:"Query" name:"access_key_id"`
	MaxConnection          requests.Integer `position:"Query" name:"MaxConnection"`
	CookieTimeout          requests.Integer `position:"Query" name:"CookieTimeout"`
	StickySessionType      string           `position:"Query" name:"StickySessionType"`
	HealthCheckConnectPort requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	LoadBalancerId         string           `position:"Query" name:"LoadBalancerId"`
	HealthyThreshold       requests.Integer `position:"Query" name:"HealthyThreshold"`
}

type SetLoadBalancerHTTPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetLoadBalancerHTTPListenerAttributeRequest() (request *SetLoadBalancerHTTPListenerAttributeRequest) {
	request = &SetLoadBalancerHTTPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerHTTPListenerAttribute", "slb", "openAPI")
	return
}

func CreateSetLoadBalancerHTTPListenerAttributeResponse() (response *SetLoadBalancerHTTPListenerAttributeResponse) {
	response = &SetLoadBalancerHTTPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

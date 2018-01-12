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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) ModifyInstanceChargeType(request *ModifyInstanceChargeTypeRequest) (response *ModifyInstanceChargeTypeResponse, err error) {
	response = CreateModifyInstanceChargeTypeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyInstanceChargeTypeWithChan(request *ModifyInstanceChargeTypeRequest) (<-chan *ModifyInstanceChargeTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceChargeTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceChargeType(request)
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

func (client *Client) ModifyInstanceChargeTypeWithCallback(request *ModifyInstanceChargeTypeRequest, callback func(response *ModifyInstanceChargeTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceChargeTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceChargeType(request)
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

type ModifyInstanceChargeTypeRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Period               requests.Integer `position:"Query" name:"Period"`
	PeriodUnit           string           `position:"Query" name:"PeriodUnit"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	IncludeDataDisks     requests.Boolean `position:"Query" name:"IncludeDataDisks"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
}

type ModifyInstanceChargeTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

func CreateModifyInstanceChargeTypeRequest() (request *ModifyInstanceChargeTypeRequest) {
	request = &ModifyInstanceChargeTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceChargeType", "ecs", "openAPI")
	return
}

func CreateModifyInstanceChargeTypeResponse() (response *ModifyInstanceChargeTypeResponse) {
	response = &ModifyInstanceChargeTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

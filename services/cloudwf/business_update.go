package cloudwf

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

// BusinessUpdate invokes the cloudwf.BusinessUpdate API synchronously
// api document: https://help.aliyun.com/api/cloudwf/businessupdate.html
func (client *Client) BusinessUpdate(request *BusinessUpdateRequest) (response *BusinessUpdateResponse, err error) {
	response = CreateBusinessUpdateResponse()
	err = client.DoAction(request, response)
	return
}

// BusinessUpdateWithChan invokes the cloudwf.BusinessUpdate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/businessupdate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BusinessUpdateWithChan(request *BusinessUpdateRequest) (<-chan *BusinessUpdateResponse, <-chan error) {
	responseChan := make(chan *BusinessUpdateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BusinessUpdate(request)
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

// BusinessUpdateWithCallback invokes the cloudwf.BusinessUpdate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/businessupdate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BusinessUpdateWithCallback(request *BusinessUpdateRequest, callback func(response *BusinessUpdateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BusinessUpdateResponse
		var err error
		defer close(result)
		response, err = client.BusinessUpdate(request)
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

// BusinessUpdateRequest is the request struct for api BusinessUpdate
type BusinessUpdateRequest struct {
	*requests.RpcRequest
	Warn             requests.Integer `position:"Query" name:"Warn"`
	BusinessCity     string           `position:"Query" name:"BusinessCity"`
	WarnEmail        string           `position:"Query" name:"WarnEmail"`
	BusinessAddress  string           `position:"Query" name:"BusinessAddress"`
	Bid              requests.Integer `position:"Query" name:"Bid"`
	BusinessManager  string           `position:"Query" name:"BusinessManager"`
	BusinessProvince string           `position:"Query" name:"BusinessProvince"`
}

// BusinessUpdateResponse is the response struct for api BusinessUpdate
type BusinessUpdateResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateBusinessUpdateRequest creates a request to invoke BusinessUpdate API
func CreateBusinessUpdateRequest() (request *BusinessUpdateRequest) {
	request = &BusinessUpdateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessUpdate", "cloudwf", "openAPI")
	return
}

// CreateBusinessUpdateResponse creates a response to parse from BusinessUpdate response
func CreateBusinessUpdateResponse() (response *BusinessUpdateResponse) {
	response = &BusinessUpdateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
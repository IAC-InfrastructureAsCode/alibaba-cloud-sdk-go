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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) CategoryTree(request *CategoryTreeRequest) (response *CategoryTreeResponse, err error) {
	response = CreateCategoryTreeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CategoryTreeWithChan(request *CategoryTreeRequest) (<-chan *CategoryTreeResponse, <-chan error) {
	responseChan := make(chan *CategoryTreeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CategoryTree(request)
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

func (client *Client) CategoryTreeWithCallback(request *CategoryTreeRequest, callback func(response *CategoryTreeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CategoryTreeResponse
		var err error
		defer close(result)
		response, err = client.CategoryTree(request)
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

type CategoryTreeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type CategoryTreeResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	CategoryTree string `json:"CategoryTree" xml:"CategoryTree"`
}

func CreateCategoryTreeRequest() (request *CategoryTreeRequest) {
	request = &CategoryTreeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "CategoryTree", "mts", "openAPI")
	return
}

func CreateCategoryTreeResponse() (response *CategoryTreeResponse) {
	response = &CategoryTreeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

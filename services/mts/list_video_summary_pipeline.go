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

func (client *Client) ListVideoSummaryPipeline(request *ListVideoSummaryPipelineRequest) (response *ListVideoSummaryPipelineResponse, err error) {
	response = CreateListVideoSummaryPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListVideoSummaryPipelineWithChan(request *ListVideoSummaryPipelineRequest) (<-chan *ListVideoSummaryPipelineResponse, <-chan error) {
	responseChan := make(chan *ListVideoSummaryPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVideoSummaryPipeline(request)
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

func (client *Client) ListVideoSummaryPipelineWithCallback(request *ListVideoSummaryPipelineRequest, callback func(response *ListVideoSummaryPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVideoSummaryPipelineResponse
		var err error
		defer close(result)
		response, err = client.ListVideoSummaryPipeline(request)
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

type ListVideoSummaryPipelineRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	State                string           `position:"Query" name:"State"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type ListVideoSummaryPipelineResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int    `json:"PageNumber" xml:"PageNumber"`
	PageSize     int    `json:"PageSize" xml:"PageSize"`
	PipelineList struct {
		Pipeline []struct {
			Id           string `json:"Id" xml:"Id"`
			Name         string `json:"Name" xml:"Name"`
			State        string `json:"State" xml:"State"`
			Priority     string `json:"Priority" xml:"Priority"`
			NotifyConfig struct {
				Topic     string `json:"Topic" xml:"Topic"`
				QueueName string `json:"QueueName" xml:"QueueName"`
			} `json:"NotifyConfig" xml:"NotifyConfig"`
		} `json:"Pipeline" xml:"Pipeline"`
	} `json:"PipelineList" xml:"PipelineList"`
}

func CreateListVideoSummaryPipelineRequest() (request *ListVideoSummaryPipelineRequest) {
	request = &ListVideoSummaryPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListVideoSummaryPipeline", "mts", "openAPI")
	return
}

func CreateListVideoSummaryPipelineResponse() (response *ListVideoSummaryPipelineResponse) {
	response = &ListVideoSummaryPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

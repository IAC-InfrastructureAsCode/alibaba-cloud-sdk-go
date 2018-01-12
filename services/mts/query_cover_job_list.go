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

func (client *Client) QueryCoverJobList(request *QueryCoverJobListRequest) (response *QueryCoverJobListResponse, err error) {
	response = CreateQueryCoverJobListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryCoverJobListWithChan(request *QueryCoverJobListRequest) (<-chan *QueryCoverJobListResponse, <-chan error) {
	responseChan := make(chan *QueryCoverJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCoverJobList(request)
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

func (client *Client) QueryCoverJobListWithCallback(request *QueryCoverJobListRequest, callback func(response *QueryCoverJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCoverJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryCoverJobList(request)
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

type QueryCoverJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CoverJobIds          string           `position:"Query" name:"CoverJobIds"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryCoverJobListResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	NonExistIds struct {
		String []string `json:"String" xml:"String"`
	} `json:"NonExistIds" xml:"NonExistIds"`
	CoverJobList struct {
		CoverJob []struct {
			Id           string `json:"Id" xml:"Id"`
			UserData     string `json:"UserData" xml:"UserData"`
			PipelineId   string `json:"PipelineId" xml:"PipelineId"`
			State        string `json:"State" xml:"State"`
			Code         string `json:"Code" xml:"Code"`
			Message      string `json:"Message" xml:"Message"`
			CreationTime string `json:"CreationTime" xml:"CreationTime"`
			Input        struct {
				Bucket   string `json:"Bucket" xml:"Bucket"`
				Location string `json:"Location" xml:"Location"`
				Object   string `json:"Object" xml:"Object"`
			} `json:"Input" xml:"Input"`
			CoverConfig struct {
				OutputFile struct {
					Bucket   string `json:"Bucket" xml:"Bucket"`
					Location string `json:"Location" xml:"Location"`
					Object   string `json:"Object" xml:"Object"`
				} `json:"OutputFile" xml:"OutputFile"`
			} `json:"CoverConfig" xml:"CoverConfig"`
			CoverImageList struct {
				CoverImage []struct {
					Score string `json:"Score" xml:"Score"`
					Url   string `json:"Url" xml:"Url"`
					Time  string `json:"Time" xml:"Time"`
				} `json:"CoverImage" xml:"CoverImage"`
			} `json:"CoverImageList" xml:"CoverImageList"`
		} `json:"CoverJob" xml:"CoverJob"`
	} `json:"CoverJobList" xml:"CoverJobList"`
}

func CreateQueryCoverJobListRequest() (request *QueryCoverJobListRequest) {
	request = &QueryCoverJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryCoverJobList", "mts", "openAPI")
	return
}

func CreateQueryCoverJobListResponse() (response *QueryCoverJobListResponse) {
	response = &QueryCoverJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package nas

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

func (client *Client) CreateMountTarget(request *CreateMountTargetRequest) (response *CreateMountTargetResponse, err error) {
	response = CreateCreateMountTargetResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateMountTargetWithChan(request *CreateMountTargetRequest) (<-chan *CreateMountTargetResponse, <-chan error) {
	responseChan := make(chan *CreateMountTargetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMountTarget(request)
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

func (client *Client) CreateMountTargetWithCallback(request *CreateMountTargetRequest, callback func(response *CreateMountTargetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMountTargetResponse
		var err error
		defer close(result)
		response, err = client.CreateMountTarget(request)
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

type CreateMountTargetRequest struct {
	*requests.RpcRequest
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	VSwitchId       string `position:"Query" name:"VSwitchId"`
	FileSystemId    string `position:"Query" name:"FileSystemId"`
	VpcId           string `position:"Query" name:"VpcId"`
	NetworkType     string `position:"Query" name:"NetworkType"`
}

type CreateMountTargetResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	MountTargetDomain string `json:"MountTargetDomain" xml:"MountTargetDomain"`
}

func CreateCreateMountTargetRequest() (request *CreateMountTargetRequest) {
	request = &CreateMountTargetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CreateMountTarget", "nas", "openAPI")
	return
}

func CreateCreateMountTargetResponse() (response *CreateMountTargetResponse) {
	response = &CreateMountTargetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

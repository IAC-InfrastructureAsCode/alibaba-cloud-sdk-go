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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeLiveStreamRelayPushData(request *DescribeLiveStreamRelayPushDataRequest) (response *DescribeLiveStreamRelayPushDataResponse, err error) {
	response = CreateDescribeLiveStreamRelayPushDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamRelayPushDataWithChan(request *DescribeLiveStreamRelayPushDataRequest) (<-chan *DescribeLiveStreamRelayPushDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRelayPushDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRelayPushData(request)
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

func (client *Client) DescribeLiveStreamRelayPushDataWithCallback(request *DescribeLiveStreamRelayPushDataRequest, callback func(response *DescribeLiveStreamRelayPushDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRelayPushDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRelayPushData(request)
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

type DescribeLiveStreamRelayPushDataRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	RelayDomain   string           `position:"Query" name:"RelayDomain"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamRelayPushDataResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	RelayPushDetailModelList struct {
		RelayPushDetailModel []struct {
			Time          string  `json:"Time" xml:"Time"`
			Stream        string  `json:"Stream" xml:"Stream"`
			FrameRate     float64 `json:"FrameRate" xml:"FrameRate"`
			BitRate       float64 `json:"BitRate" xml:"BitRate"`
			FrameLossRate float64 `json:"FrameLossRate" xml:"FrameLossRate"`
			ServerAddr    string  `json:"ServerAddr" xml:"ServerAddr"`
			ClientAddr    string  `json:"ClientAddr" xml:"ClientAddr"`
		} `json:"RelayPushDetailModel" xml:"RelayPushDetailModel"`
	} `json:"RelayPushDetailModelList" xml:"RelayPushDetailModelList"`
}

func CreateDescribeLiveStreamRelayPushDataRequest() (request *DescribeLiveStreamRelayPushDataRequest) {
	request = &DescribeLiveStreamRelayPushDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRelayPushData", "", "")
	return
}

func CreateDescribeLiveStreamRelayPushDataResponse() (response *DescribeLiveStreamRelayPushDataResponse) {
	response = &DescribeLiveStreamRelayPushDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

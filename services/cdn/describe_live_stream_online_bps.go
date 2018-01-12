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

func (client *Client) DescribeLiveStreamOnlineBps(request *DescribeLiveStreamOnlineBpsRequest) (response *DescribeLiveStreamOnlineBpsResponse, err error) {
	response = CreateDescribeLiveStreamOnlineBpsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamOnlineBpsWithChan(request *DescribeLiveStreamOnlineBpsRequest) (<-chan *DescribeLiveStreamOnlineBpsResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamOnlineBpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamOnlineBps(request)
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

func (client *Client) DescribeLiveStreamOnlineBpsWithCallback(request *DescribeLiveStreamOnlineBpsRequest, callback func(response *DescribeLiveStreamOnlineBpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamOnlineBpsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamOnlineBps(request)
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

type DescribeLiveStreamOnlineBpsRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamOnlineBpsResponse struct {
	*responses.BaseResponse
	RequestId                string  `json:"RequestId" xml:"RequestId"`
	TotalUserNumber          int     `json:"TotalUserNumber" xml:"TotalUserNumber"`
	FlvBps                   float64 `json:"FlvBps" xml:"FlvBps"`
	HlsBps                   float64 `json:"HlsBps" xml:"HlsBps"`
	LiveStreamOnlineBpsInfos struct {
		LiveStreamOnlineBpsInfo []struct {
			StreamUrl string  `json:"StreamUrl" xml:"StreamUrl"`
			UpBps     float64 `json:"UpBps" xml:"UpBps"`
			DownBps   float64 `json:"DownBps" xml:"DownBps"`
			Time      string  `json:"Time" xml:"Time"`
		} `json:"LiveStreamOnlineBpsInfo" xml:"LiveStreamOnlineBpsInfo"`
	} `json:"LiveStreamOnlineBpsInfos" xml:"LiveStreamOnlineBpsInfos"`
}

func CreateDescribeLiveStreamOnlineBpsRequest() (request *DescribeLiveStreamOnlineBpsRequest) {
	request = &DescribeLiveStreamOnlineBpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamOnlineBps", "", "")
	return
}

func CreateDescribeLiveStreamOnlineBpsResponse() (response *DescribeLiveStreamOnlineBpsResponse) {
	response = &DescribeLiveStreamOnlineBpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

func (client *Client) ModifyInstanceAutoReleaseTime(request *ModifyInstanceAutoReleaseTimeRequest) (response *ModifyInstanceAutoReleaseTimeResponse, err error) {
	response = CreateModifyInstanceAutoReleaseTimeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyInstanceAutoReleaseTimeWithChan(request *ModifyInstanceAutoReleaseTimeRequest) (<-chan *ModifyInstanceAutoReleaseTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceAutoReleaseTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceAutoReleaseTime(request)
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

func (client *Client) ModifyInstanceAutoReleaseTimeWithCallback(request *ModifyInstanceAutoReleaseTimeRequest, callback func(response *ModifyInstanceAutoReleaseTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceAutoReleaseTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceAutoReleaseTime(request)
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

type ModifyInstanceAutoReleaseTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AutoReleaseTime      string           `position:"Query" name:"AutoReleaseTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type ModifyInstanceAutoReleaseTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyInstanceAutoReleaseTimeRequest() (request *ModifyInstanceAutoReleaseTimeRequest) {
	request = &ModifyInstanceAutoReleaseTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceAutoReleaseTime", "ecs", "openAPI")
	return
}

func CreateModifyInstanceAutoReleaseTimeResponse() (response *ModifyInstanceAutoReleaseTimeResponse) {
	response = &ModifyInstanceAutoReleaseTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

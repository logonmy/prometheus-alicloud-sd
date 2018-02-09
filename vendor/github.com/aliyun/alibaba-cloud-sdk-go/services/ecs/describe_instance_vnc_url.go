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

func (client *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
	response = CreateDescribeInstanceVncUrlResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInstanceVncUrlWithChan(request *DescribeInstanceVncUrlRequest) (<-chan *DescribeInstanceVncUrlResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceVncUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceVncUrl(request)
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

func (client *Client) DescribeInstanceVncUrlWithCallback(request *DescribeInstanceVncUrlRequest, callback func(response *DescribeInstanceVncUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceVncUrlResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceVncUrl(request)
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

type DescribeInstanceVncUrlRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeInstanceVncUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	VncUrl    string `json:"VncUrl" xml:"VncUrl"`
}

func CreateDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
	request = &DescribeInstanceVncUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceVncUrl", "ecs", "openAPI")
	return
}

func CreateDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
	response = &DescribeInstanceVncUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

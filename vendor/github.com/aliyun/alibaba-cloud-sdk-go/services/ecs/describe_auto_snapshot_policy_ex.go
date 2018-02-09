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

func (client *Client) DescribeAutoSnapshotPolicyEx(request *DescribeAutoSnapshotPolicyExRequest) (response *DescribeAutoSnapshotPolicyExResponse, err error) {
	response = CreateDescribeAutoSnapshotPolicyExResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeAutoSnapshotPolicyExWithChan(request *DescribeAutoSnapshotPolicyExRequest) (<-chan *DescribeAutoSnapshotPolicyExResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoSnapshotPolicyExResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoSnapshotPolicyEx(request)
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

func (client *Client) DescribeAutoSnapshotPolicyExWithCallback(request *DescribeAutoSnapshotPolicyExRequest, callback func(response *DescribeAutoSnapshotPolicyExResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoSnapshotPolicyExResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoSnapshotPolicyEx(request)
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

type DescribeAutoSnapshotPolicyExRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string           `position:"Query" name:"AutoSnapshotPolicyId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
}

type DescribeAutoSnapshotPolicyExResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	TotalCount           int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber           int    `json:"PageNumber" xml:"PageNumber"`
	PageSize             int    `json:"PageSize" xml:"PageSize"`
	AutoSnapshotPolicies struct {
		AutoSnapshotPolicy []struct {
			AutoSnapshotPolicyId   string `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
			RegionId               string `json:"RegionId" xml:"RegionId"`
			AutoSnapshotPolicyName string `json:"AutoSnapshotPolicyName" xml:"AutoSnapshotPolicyName"`
			TimePoints             string `json:"TimePoints" xml:"TimePoints"`
			RepeatWeekdays         string `json:"RepeatWeekdays" xml:"RepeatWeekdays"`
			RetentionDays          int    `json:"RetentionDays" xml:"RetentionDays"`
			DiskNums               int    `json:"DiskNums" xml:"DiskNums"`
			VolumeNums             int    `json:"VolumeNums" xml:"VolumeNums"`
			CreationTime           string `json:"CreationTime" xml:"CreationTime"`
			Status                 string `json:"Status" xml:"Status"`
		} `json:"AutoSnapshotPolicy" xml:"AutoSnapshotPolicy"`
	} `json:"AutoSnapshotPolicies" xml:"AutoSnapshotPolicies"`
}

func CreateDescribeAutoSnapshotPolicyExRequest() (request *DescribeAutoSnapshotPolicyExRequest) {
	request = &DescribeAutoSnapshotPolicyExRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAutoSnapshotPolicyEx", "ecs", "openAPI")
	return
}

func CreateDescribeAutoSnapshotPolicyExResponse() (response *DescribeAutoSnapshotPolicyExResponse) {
	response = &DescribeAutoSnapshotPolicyExResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

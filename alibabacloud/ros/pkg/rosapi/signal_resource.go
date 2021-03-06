package rosapi

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

// SignalResource invokes the ros.SignalResource API synchronously
// api document: https://help.aliyun.com/api/ros/signalresource.html
func (client *Client) SignalResource(request *SignalResourceRequest) (response *SignalResourceResponse, err error) {
	response = CreateSignalResourceResponse()
	err = client.DoAction(request, response)
	return
}

// SignalResourceWithChan invokes the ros.SignalResource API asynchronously
// api document: https://help.aliyun.com/api/ros/signalresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SignalResourceWithChan(request *SignalResourceRequest) (<-chan *SignalResourceResponse, <-chan error) {
	responseChan := make(chan *SignalResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SignalResource(request)
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

// SignalResourceWithCallback invokes the ros.SignalResource API asynchronously
// api document: https://help.aliyun.com/api/ros/signalresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SignalResourceWithCallback(request *SignalResourceRequest, callback func(response *SignalResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SignalResourceResponse
		var err error
		defer close(result)
		response, err = client.SignalResource(request)
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

// SignalResourceRequest is the request struct for api SignalResource
type SignalResourceRequest struct {
	*requests.RpcRequest
	Data              string `position:"Query" name:"Data"`
	ClientToken       string `position:"Query" name:"ClientToken"`
	StackId           string `position:"Query" name:"StackId"`
	LogicalResourceId string `position:"Query" name:"LogicalResourceId"`
	UniqueId          string `position:"Query" name:"UniqueId"`
	Status            string `position:"Query" name:"Status"`
}

// SignalResourceResponse is the response struct for api SignalResource
type SignalResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSignalResourceRequest creates a request to invoke SignalResource API
func CreateSignalResourceRequest() (request *SignalResourceRequest) {
	request = &SignalResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ROS", "2019-09-10", "SignalResource", "ros", "openAPI")
	return
}

// CreateSignalResourceResponse creates a response to parse from SignalResource response
func CreateSignalResourceResponse() (response *SignalResourceResponse) {
	response = &SignalResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

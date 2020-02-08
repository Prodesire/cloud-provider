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

// DeleteTemplate invokes the ros.DeleteTemplate API synchronously
// api document: https://help.aliyun.com/api/ros/deletetemplate.html
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
	response = CreateDeleteTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTemplateWithChan invokes the ros.DeleteTemplate API asynchronously
// api document: https://help.aliyun.com/api/ros/deletetemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTemplateWithChan(request *DeleteTemplateRequest) (<-chan *DeleteTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTemplate(request)
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

// DeleteTemplateWithCallback invokes the ros.DeleteTemplate API asynchronously
// api document: https://help.aliyun.com/api/ros/deletetemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTemplateWithCallback(request *DeleteTemplateRequest, callback func(response *DeleteTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteTemplate(request)
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

// DeleteTemplateRequest is the request struct for api DeleteTemplate
type DeleteTemplateRequest struct {
	*requests.RpcRequest
	TemplateId string `position:"Query" name:"TemplateId"`
}

// DeleteTemplateResponse is the response struct for api DeleteTemplate
type DeleteTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTemplateRequest creates a request to invoke DeleteTemplate API
func CreateDeleteTemplateRequest() (request *DeleteTemplateRequest) {
	request = &DeleteTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ROS", "2019-09-10", "DeleteTemplate", "ros", "openAPI")
	return
}

// CreateDeleteTemplateResponse creates a response to parse from DeleteTemplate response
func CreateDeleteTemplateResponse() (response *DeleteTemplateResponse) {
	response = &DeleteTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

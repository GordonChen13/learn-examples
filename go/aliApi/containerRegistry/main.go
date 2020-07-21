package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type Tag struct {
	Status string `json:"status"`
	Tag string `json:"tag"`
}

type Data struct {
	Page int `json:"page"`
	Total int `json:"total"`
	Tags []Tag `json:"tags"`
}

type Res struct {
	Data Data `json:"data"`
}

var repo = flag.String("repo", "", "gitlab repo names, example: tools/mars")

func main() {
	client, err := sdk.NewClientWithAccessKey(REGIONID, ACCESSKEYID, ACCESSKEYSECRET)
	if err != nil {
		printErrAndExit(err.Error())
	}

	group, project := parseRepo()

	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Scheme = "https"
	request.Domain = "cr." + REGIONID + ".aliyuncs.com"
	request.Version = "2016-06-07"
	request.PathPattern = "/repos/myunke/" + group + "-" + project + "/tags"
	request.QueryParams["PageSize"] = "100"
	body := ``
	request.Content = []byte(body)
	request.TransToAcsRequest()

	response := &responses.BaseResponse{}
	err = client.DoAction(request, response)
	if err != nil {
		printErrAndExit(err.Error())
	}

	var res Res
	err = json.Unmarshal(response.GetHttpContentBytes(), &res)
	if err != nil {
		printErrAndExit(err.Error())
	}

	// output raw json response
	//fmt.Println(response.GetHttpContentString())

	for _, tag := range res.Data.Tags  {
		fmt.Println(tag.Tag)
	}
}

func parseRepo() (group, project string) {
	flag.Parse()
	strs := strings.Split(*repo, "/")
	if len(strs) != 2{
		printErrAndExit("error: 解析repo参数出错")
	}

	group = strs[0]
	project = strs[1]
	return
}

func printErrAndExit(err string)  {
	fmt.Println(err)
	os.Exit(1)
}
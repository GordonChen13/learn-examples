#!/usr/bin/env python
# coding=utf-8
from aliyunsdkcore.acs_exception.exceptions import ClientException
from aliyunsdkcore.acs_exception.exceptions import ServerException
from aliyunsdkcore.client import AcsClient
# from aliyunsdkcr.request.v20160607 import GetImageLayerRequest
from aliyunsdkcr.request.v20160607 import GetRepoTagsRequest
# 示例执行异常时建议升级aliyun-python-sdk-core到最新版本
# 设置Client
apiClient = AcsClient('your access key id', 'your access key secret', 'cn-shenzhen')
# 构造请求
# request = GetImageLayerRequest.GetImageLayerRequest()
request = GetRepoTagsRequest.GetRepoTagsRequest()
# 设置参数
request.set_RepoNamespace("namespace")
request.set_RepoName("repo name")
# request.set_Tag("tag")
# 根据文档获取资源所在区域对应的RegionId
# 请求地址格式为cr.{regionId}.aliyuncs.com
request.set_endpoint("cr.cn-shenzhen.aliyuncs.com")
# 发起请求
try:
    response = apiClient.do_action_with_exception(request)
    print(response)
except ServerException as e:
    print(e)
except ClientException as e:
    print(e)
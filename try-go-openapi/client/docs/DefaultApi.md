# Api.DefaultApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createComment**](DefaultApi.md#createComment) | **POST** /posts/{postId}/comments | 
[**createPost**](DefaultApi.md#createPost) | **POST** /posts | 
[**deleteComment**](DefaultApi.md#deleteComment) | **DELETE** /posts/{postId}/comments/{commentId} | 
[**deletePost**](DefaultApi.md#deletePost) | **DELETE** /posts/{postId} | 
[**getComments**](DefaultApi.md#getComments) | **GET** /posts/{postId}/comments | 
[**getPost**](DefaultApi.md#getPost) | **GET** /posts/{postId} | 
[**getPosts**](DefaultApi.md#getPosts) | **GET** /posts | 
[**updatePost**](DefaultApi.md#updatePost) | **PUT** /posts/{postId} | 


<a name="createComment"></a>
# **createComment**
> Comment createComment(postId, commentRequest)



投稿にコメントを付ける

### Example
```javascript
var Api = require('api');

var apiInstance = new Api.DefaultApi();
var postId = 56; // Number | 投稿のID
var commentRequest = new Api.CommentRequest(); // CommentRequest | 投稿に付けるコメント
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.createComment(postId, commentRequest, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **Number**| 投稿のID | 
 **commentRequest** | [**CommentRequest**](CommentRequest.md)| 投稿に付けるコメント | 

### Return type

[**Comment**](Comment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="createPost"></a>
# **createPost**
> Post createPost(postRequest)



投稿を作成する

### Example
```javascript
var Api = require('api');

var apiInstance = new Api.DefaultApi();
var postRequest = new Api.PostRequest(); // PostRequest | 作成する投稿
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.createPost(postRequest, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRequest** | [**PostRequest**](PostRequest.md)| 作成する投稿 | 

### Return type

[**Post**](Post.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteComment"></a>
# **deleteComment**
> deleteComment(postId, commentId)



投稿に付いたコメントを削除する

### Example
```javascript
var Api = require('api');

var apiInstance = new Api.DefaultApi();
var postId = 56; // Number | 投稿のID
var commentId = 56; // Number | コメントのID
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteComment(postId, commentId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **Number**| 投稿のID | 
 **commentId** | **Number**| コメントのID | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="deletePost"></a>
# **deletePost**
> deletePost(postId)



投稿を削除する

### Example
```javascript
var Api = require('api');

var apiInstance = new Api.DefaultApi();
var postId = 56; // Number | 投稿のID
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deletePost(postId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **Number**| 投稿のID | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="getComments"></a>
# **getComments**
> Comments getComments(postId)



投稿に付いたコメントをすべて取得する

### Example
```javascript
var Api = require('api');

var apiInstance = new Api.DefaultApi();
var postId = 56; // Number | 投稿のID
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getComments(postId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **Number**| 投稿のID | 

### Return type

[**Comments**](Comments.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="getPost"></a>
# **getPost**
> Post getPost(postId)



投稿を取得する

### Example
```javascript
var Api = require('api');

var apiInstance = new Api.DefaultApi();
var postId = 56; // Number | 投稿のID
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getPost(postId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **Number**| 投稿のID | 

### Return type

[**Post**](Post.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="getPosts"></a>
# **getPosts**
> Posts getPosts()



投稿をすべて取得する

### Example
```javascript
var Api = require('api');

var apiInstance = new Api.DefaultApi();
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getPosts(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Posts**](Posts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updatePost"></a>
# **updatePost**
> Post updatePost(postId, postRequest)



投稿を更新する

### Example
```javascript
var Api = require('api');

var apiInstance = new Api.DefaultApi();
var postId = 56; // Number | 投稿のID
var postRequest = new Api.PostRequest(); // PostRequest | 作成する投稿
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.updatePost(postId, postRequest, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **Number**| 投稿のID | 
 **postRequest** | [**PostRequest**](PostRequest.md)| 作成する投稿 | 

### Return type

[**Post**](Post.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


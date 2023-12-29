# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [todo_service/todo_service.proto](#todo_service_todo_service-proto)
    - [AddRequest](#todo_service-AddRequest)
    - [AddResponse](#todo_service-AddResponse)
    - [DeleteRequest](#todo_service-DeleteRequest)
    - [DeleteResponse](#todo_service-DeleteResponse)
    - [GetRequest](#todo_service-GetRequest)
    - [GetResponse](#todo_service-GetResponse)
    - [Todo](#todo_service-Todo)
    - [UpdateStatusRequest](#todo_service-UpdateStatusRequest)
    - [UpdateStatusResponse](#todo_service-UpdateStatusResponse)
    - [UpdateTitleRequest](#todo_service-UpdateTitleRequest)
    - [UpdateTitleResponse](#todo_service-UpdateTitleResponse)
  
    - [TodoStatus](#todo_service-TodoStatus)
  
    - [TodoService](#todo_service-TodoService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="todo_service_todo_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## todo_service/todo_service.proto



<a name="todo_service-AddRequest"></a>

### AddRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todo | [Todo](#todo_service-Todo) |  |  |






<a name="todo_service-AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="todo_service-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="todo_service-DeleteResponse"></a>

### DeleteResponse







<a name="todo_service-GetRequest"></a>

### GetRequest







<a name="todo_service-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [Todo](#todo_service-Todo) | repeated |  |






<a name="todo_service-Todo"></a>

### Todo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| title | [string](#string) |  |  |
| status | [TodoStatus](#todo_service-TodoStatus) |  |  |
| status_name | [string](#string) |  |  |






<a name="todo_service-UpdateStatusRequest"></a>

### UpdateStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| status | [TodoStatus](#todo_service-TodoStatus) |  |  |






<a name="todo_service-UpdateStatusResponse"></a>

### UpdateStatusResponse







<a name="todo_service-UpdateTitleRequest"></a>

### UpdateTitleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| title | [string](#string) |  |  |






<a name="todo_service-UpdateTitleResponse"></a>

### UpdateTitleResponse






 


<a name="todo_service-TodoStatus"></a>

### TodoStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| TODO_STATUS_UNSPECIFIED | 0 |  |
| TODO_STATUS_DOING | 1 |  |
| TODO_STATUS_DONE | 2 |  |


 

 


<a name="todo_service-TodoService"></a>

### TodoService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#todo_service-GetRequest) | [GetResponse](#todo_service-GetResponse) |  |
| Add | [AddRequest](#todo_service-AddRequest) | [AddResponse](#todo_service-AddResponse) |  |
| UpdateStatus | [UpdateStatusRequest](#todo_service-UpdateStatusRequest) | [UpdateStatusResponse](#todo_service-UpdateStatusResponse) |  |
| UpdateTitle | [UpdateTitleRequest](#todo_service-UpdateTitleRequest) | [UpdateTitleResponse](#todo_service-UpdateTitleResponse) |  |
| Delete | [DeleteRequest](#todo_service-DeleteRequest) | [DeleteResponse](#todo_service-DeleteResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |


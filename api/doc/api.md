# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [authority_service.proto](#authority_service-proto)
    - [AuthorityService](#auth-AuthorityService)
  
- [model.proto](#model-proto)
    - [Action](#auth-Action)
    - [Role](#auth-Role)
  
    - [Permission](#auth-Permission)
  
- [role_service.proto](#role_service-proto)
    - [CreateRoleBindingRequest](#auth-CreateRoleBindingRequest)
    - [DeleteRoleBindingRequest](#auth-DeleteRoleBindingRequest)
    - [DeleteRoleRequest](#auth-DeleteRoleRequest)
    - [GetRoleRequest](#auth-GetRoleRequest)
  
    - [RoleService](#auth-RoleService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="authority_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## authority_service.proto


 

 

 


<a name="auth-AuthorityService"></a>

### AuthorityService
Authorization API for public use

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AuthorizeAction | [Action](#auth-Action) | [.google.protobuf.Empty](#google-protobuf-Empty) | AuthorizeAction checks if the user is authorized to perform the action |

 



<a name="model-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model.proto



<a name="auth-Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | user identifier |
| org_id | [string](#string) |  | organization identifier |
| permission | [Permission](#auth-Permission) |  | permission to check |






<a name="auth-Role"></a>

### Role



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_id | [string](#string) |  | role identifier |
| permissions | [Permission](#auth-Permission) | repeated | list of permissions |





 


<a name="auth-Permission"></a>

### Permission


| Name | Number | Description |
| ---- | ------ | ----------- |
| PERM_NOOP | 0 | Indicates no value, will lead to an error if set |
| PERM_MANAGE_USERS | 1 | Add or remove other users to an organization |
| PERM_MODIFY_USER_PERMISSIONS | 2 | Modify a user’s permission in an organization |
| PERM_CREATE_PROJECT | 3 | Create a new project in an organization |
| PERM_DELETE_PROJECT | 4 | Remove a project from an organization |
| PERM_DEPLOY_PROJECT | 5 | Deploy a project to a specific environment |
| PERM_MANAGE_ENVIRONMENTS | 6 | Modify environments definitions for a project |
| PERM_READ_LOGS | 7 | Read logs for a project |
| PERM_MODIFY_LOGS | 8 | Edit logs for a project |
| PERM_AUDIT_LOGS | 9 | Audit logs and sensitive data for a project |


 

 

 



<a name="role_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## role_service.proto



<a name="auth-CreateRoleBindingRequest"></a>

### CreateRoleBindingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_id | [string](#string) |  | role identifier |
| user_id | [string](#string) |  | user identifier |
| org_id | [string](#string) |  | organization identifier |






<a name="auth-DeleteRoleBindingRequest"></a>

### DeleteRoleBindingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_id | [string](#string) |  | role identifier |
| user_id | [string](#string) |  | user identifier |






<a name="auth-DeleteRoleRequest"></a>

### DeleteRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id is the unique identifier of the role. |






<a name="auth-GetRoleRequest"></a>

### GetRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id is the unique identifier of the role. |





 

 

 


<a name="auth-RoleService"></a>

### RoleService
Administrator API for role management.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetRole | [GetRoleRequest](#auth-GetRoleRequest) | [Role](#auth-Role) | returns existing role |
| CreateRole | [Role](#auth-Role) | [Role](#auth-Role) | creates a new role |
| UpdateRole | [Role](#auth-Role) | [Role](#auth-Role) | updates existing role |
| DeleteRole | [DeleteRoleRequest](#auth-DeleteRoleRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | deletes existing role |
| CreateRoleBinding | [CreateRoleBindingRequest](#auth-CreateRoleBindingRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | binds role to user in a scope of organization |
| DeleteRoleBinding | [DeleteRoleBindingRequest](#auth-DeleteRoleBindingRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | unbinds role from user in all scopes |

 



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


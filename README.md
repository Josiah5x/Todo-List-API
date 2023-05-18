## Todo-List-API
 Todo-List-API-With-Golang MVC 

 Welcome to our project! This is a Todo list App made in Golang using MVC pattern. It allows users to quickly and easily Add,Update Delete. With this App, users can effortlessly manage their tasks and events, helping them stay organized and on top of their to-do list. We hope you enjoy and we look forward to your contributions!

------------------------------------------------------------------

## User Section
>
- To access our todo you must be Identify as a potential User
- Login if register
- A Token will be given valid for 24h
- Below are the Steps to us

------------------------------------------------------------------


#### Creating new

<details>
 <summary><code>POST</code> <code><b>/v1/user</b></code> <code>(Creating new user)</code></summary>

##### #Request
``` javascript
{
    "userid":"",
    "firstname":"nungrie",
    "lastname":"galadima",
    "email":"nungrie0519@gmail.com",
    "username": "nung",
    "password": "3534"
}
```

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | None      |  required | object (JSON)   | N/A  |


##### Responses

``` javascript
User Successfully Added
```

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8080/v1/user/

> ```

</details>
<details>
 <summary><code>POST</code> <code><b>/v1/user/login</b></code> <code>(User login)</code></summary>

##### #Request
``` javascript
{
    "username": "nung",
    "password": "3534"
}
```

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | None      |  required | object (JSON)   | N/A  |


##### Responses

``` javascript
User Successfully Login
```

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8080/v1/user/

> ```

</details>

------------------------------------------------------

#### List of existing users 

<details>
 <summary><code>GET</code> <code><b>/v1/user/all</b></code> <code>(gets all User)</code></summary>

##### Parameters

> None

##### Responses

``` javascript
 {
    "_id": "5ab6ca49-fa35-4049-8b0d-f6f35db4bc85",
    "firstname": "nungrie",
    "lastname": "galadima",
    "email": "nungrie0519@gmail.com",
    "username": "nung",
    "password": "3534",
    "create_at": "2023-04-19T15:49:50.982+01:00",
    "update_at": "0001-01-01T00:00:00Z"
    .....
  }                                     
```

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/v1/user/all
> ```

</details>

<details>
 <summary><code>GET</code> <code><b>/v1/user/some-unique-uuid-string</b></code> <code>(gets a specific User by some-unique-uuid-string)</code></summary>

##### Parameters

> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `stub_numeric_id` |  required | int ($int64)   | The specific stub numeric some-unique-uuid-string        |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `text/plain;charset=UTF-8`       string                                                         |
> | `400`         | `application/json`                | `{"code":"400","message":"Bad Request"}`                            |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/v1/user/some-specific-unique-uuid-string
> ```

</details>


------------------------------------------------------------------------------------------


#### Update existing User


<details>
  <summary><code>PUT</code> <code><b>/v1/user/some-specific-unique-uuid-string</b></code> <code>(updates user by some-unique-uuid-string)</code></summary>

##### #Request
``` javascript
{
    "_id": "5ab6ca49-fa35-4049-8b0d-f6f35db4bc85",
    "firstname": "nungrie",
    "lastname": "galadima",
    "email": "nungrie0519@gmail.com",
    "username": "nung",
    "password": "3534",
}
```

> | name   |  type      | data type      | description                                                  |
> |--------|------------|----------------|--------------------------------------------------------------|
> | `uuid` |  required  | string         | The specific unique idendifier                  |

##### Responses

> Successfully Updated

##### Example cURL

> ```javascript
>  curl -X PUT -H "Content-Type: application/json" --data @put.json http://localhost:8080/v1/user/some-unique-uuid-string
> ```

</details>

------------------------------------------------------------------------------------------

#### Update user by Patching 

###### #Change Password
<details>
 <summary><code>PATCH</code> <code><b>/v1/user/changepassword/some-unique-uuid-string</b></code> <code>(updating a password by specific user by patching )</code></summary>

##### #Request

``` javascript
{
  "password": "1234"
}
```

##### Responses

``` javascript
Password Successfully Updated 
```

##### Example cURL

> ```javascript
>  curl -X PATCH -H "Content-Type: application/json" http://localhost:8080/v1/user/changepassword/some-unique-uuid-string
> ```

</details>

------------------------------------------------------------------------------------------

#### Deleting existing User

<details>
  <summary><code>DELETE</code> <code><b>/v1/user/delete/some-unique-uuid-string</b></code> <code>(Delete user by some-unique-uuid-string)</code></summary>

##### Parameters

> | name   |  type      | data type      | description                                          |
> |--------|------------|----------------|------------------------------------------------------|
> | `uuid` |  required  | string         | The specific stub unique idendifier                  |


##### Responses

``` javascript
Successfully Deleted
```

##### Example cURL

> ```javascript
>  curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/v1/user/delete/some-unique-uuid-string
> ```

</details>

------------------------------------------------------------------------------------------

## Todos Section
- User Section must be satisfy before accessing the todo endpoint.
------------------------------------------------------------------

#### Creating new Todo

<details>
 <summary><code>POST</code> <code><b>/v1/todo</b></code> <code>(Creating new todo)</code></summary>

##### Body(--data @post)
``` javascript
 {
  "name": "DevOP",
  "description": "Basic of DevOp",
  "deadline": "12h"
}
```

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | None      |  required | object (JSON)   | N/A  |


##### Responses

``` javascript
{
  "_id": "218317a5-696d-4379-a178-9e3c9585deaa",
  "userid": "c6d92ec5-5f1e-40ea-9ac3-c6aaf3b5813d",
  "name": "DevOP",
  "description": "Basic of DevOp",
  "status": "Undone",
  "deadline": "17 May 23 05:23AM",
  "CreatedAt": "2023-05-16T17:23:37.271105926+01:00",
  "UpdatedAt": "0001-01-01T00:00:00Z"
}
```

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8080/v1/todos/

> ```

</details>

------------------------------------------------------------------------------------------

#### Listing existing Todo

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>(gets all todos)</code></summary>

##### Parameters

> None

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `text/plain;charset=UTF-8`       string                                                         |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/v1/todos/
> ```

</details>

<details>
 <summary><code>GET</code> <code><b>/some-unique-uuid</b></code> <code>(gets todo by some-unique-uuid)</code></summary>

##### Parameters

> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `uuid` |  required | string   | T some-unique-uuid string        |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `text/plain;charset=UTF-8`               string                                                         |
> | `400`         | `application/json`                | `{"code":"400","message":"Bad Request"}`                            |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/v1/todos/
> ```

</details>


------------------------------------------------------------------------------------------


#### Updating existing Todo

<details>
  <summary><code>PUT</code> <code><b>/v1/todos/some-unique-uuid</b></code> <code>(updates todos by some-unique-uuid)</code></summary>
 
##### Request
 ``` javascript
 {
  "name": "DevOP",
  "description": "Basic of DevOp",
  "deadline": "12h"
}
```

##### Parameters

> | name   |  type      | data type      | description                                                  |
> |--------|------------|----------------|--------------------------------------------------------------|
> | `uuid` |  required  | string         | The specific proxy config unique idendifier                  |

##### Responses

> updated successfully

##### Example cURL

> ```javascript
>  curl -X PUT -H "Content-Type: application/json" --data @put.json http://localhost:8080/v1/todos/some-unique-uuid-string
> ```

</details>

----------------------------------------------

#### Update Todos by Patching 

###### #Mark Done/Undone / #Deadline
<details>
 <summary><code>PATCH</code> <code><b>/v1/todos/mark/some-unique-uuid</b></code> <code>(updating a specific todos by patching Done or Undone)</code></summary>

  ##### Request
 
 ``` javascript
 {
 "status": "Done"
 }
 
 ```
 
##### Parameters

> | name   |  type      | data type      | description                                                  |
> |--------|------------|----------------|--------------------------------------------------------------|
> | `uuid` |  required  | string         | The specific unique idendifier                  |


##### Responses

 > Update successfully                                                  |

##### Example cURL

> ```javascript
>  curl -X PATCH -H "Content-Type: application/json" http://localhost:8080/v1/todos/some-unique-uuid
> ```

</details>

<details>
 <summary><code>PATCH</code> <code><b>/v1/todos/deadline/some-unique-uuid</b></code> <code>(update the todos by patching deadline(eg.Date) )</code></summary>

 ##### Request
 
 ``` javascript
 {
 "deadline": "24h"
 }
 
 ```
 
##### Parameters

> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `uuid` |  required | string  | The specific unique-uuid string       |

##### Responses

 > Update successfully

##### Example cURL

> ```javascript
>  curl -X PATCH -H "Content-Type: application/json" http://localhost:8080/v1/todos/some-unique-uuid string
> ```

</details>

------------------------------------------------------------------------------------------

#### Deleting existing Todo

<details>
  <summary><code>DELETE</code> <code><b>/v1/todos/some-unique-uuid</b></code> <code>(Delete todos by some-unique-uuid)</code></summary>

##### Parameters

> | name   |  type      | data type      | description                                          |
> |--------|------------|----------------|------------------------------------------------------|
> | `uuid` |  required  | string         |    The specific unique-uuid string                  |


##### Responses

 > Deleted successfully
 
##### Example cURL

> ```javascript
>  curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/v1/todos/some-unique-uuid
> ```

</details>


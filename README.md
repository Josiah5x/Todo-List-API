## Todo-List-API
 Todo-List-API-With-Golang MVC 

 Welcome to our project! This is a Todo list App made in Golang using MVC pattern and Databas(mongoDB). It allows users to quickly and easily Add,Update Delete. With this App, users can effortlessly manage their tasks and events, helping them stay organized and on top of their to-do list. We hope you enjoy and we look forward to your contributions!

------------------------------------------------------------------

## User Section

#### Register :
* Require Firstname, Lastname Email Username, password
* Return appropriate response

#### Login :
* Require email and password as the post body.
* Authenticate the user through provided credentials.
* If the user is verified and the email/password is correct, respond with a JWT token, otherwise show the appropriate response


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

#### Create ToDo :
* Only accessible for authenticated users
2 fields are required: name, description and deadline of the ToDo item
* On success save the item to the database
* Return appropriate response


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
Todo created successfully 
```

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8080/v1/todos/

> ```

</details>

------------------------------------------------------------------------------------------

#### View ToDo :
* Only accessible for authenticated users
* Respond with an object containing the title and description of the requested ToDo

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
#### Update Todos by Patching :
###### #Status Done/Undone / #Deadline
* Only accessible for authenticated users
* 1 fields are required: Status and Deadline of the ToDo item
* On success save an updated item to the database
* Return appropriate respon

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
 <summary><code>PATCH</code> <code><b>/v1/todos/deadline/some-unique-uuid</b></code> <code>(update the todos by patching deadline(eg 24h) )</code></summary>

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

#### Delete ToDo :
* Only accessible for authenticated users.
* On success delete the item from the database.
* Return appropriate response.

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


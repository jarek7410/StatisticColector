# veb servise for collecting statistic

## requariments

postgrass database with default:

 - dbHost = "localhost"
 - dbPost = 5432
 - dbUser = "user"
 - dbName = "statsAndUsers"
 - dbPass = "example"

to change in mian.go

## endpoints:
all default at http. u can make it safe for server wiht certbot and reverseproxy

GET x.x.x.x/  - helthchek
  repose:
  - 200 - ok
  - 500  - Internal Server Error
  - 503 - Service Unavailable
GET x.x.x.x/tea      - coffee (i need one, future page with paypal for me for one dolar coffee)
  - 418 - I'm a teapot
GET x.x.x.x/coffee   - coffee
  - 418 - I'm a teapot

v1:
  - POST x.x.x.x/user/ - create user
    - {
	      "username": string,
	      "password": string
      }
    - 200 - ok
    - 400 - some error (more information in body)
  - GET x.x.x.x./user - get all users (not safe for now), list of all users (id,createdad,updatedat,dalatedat,username,passward)
    - 200 - ok
    - 500 - server error
  - GET x.x.x.x/user/:id - get user by id
    - 200 - ok
    - 404 - user not found
    - 500 - server error
  - POST x.x.x.x/stat/ - crate new stat
    - {
      	"name":string,
      	"value": loat64,
      	"userId": uint 
      }
    - 200 - ok
    - 400 - bad request
    - 404 - user not fournd
  - GET x.x.x.x/stat/user/:id get stats by user
    - 501 - not implemented
  - GET x.x.x.x/stat/:id get stat by its id
    - 501 - not imlemented

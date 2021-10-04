# API Gateway Example

This repo is to test the ability of Go (GoLang) as an API Gateway between the
Front-end and the Micro-services as the Back-end.

### Basic Structure
```text
- root
    | - micro
        | - microservice
            | - # all microservice files for a specific micro-service
        | - microtest
            | - # files for testing micro-services
    | - src
        | - # all api-getway files
    main.go 
```
In this example code, the `micro/microtest` is a go package to simplify the 
testing of microservices in the API gateway setting. 
```text
+----+    Exchange   +---------------+    Exchange    +-----------------+
| FE |  <----------> |  API Gateway  |  <---------->  |  Micro-Service  |
+----+               +---------------+        \       +-----------------+--+
                                                ------>  |  Micro-Service  |
                                                         +-----------------+
```
The basic assumed structure of an API Gateway in the Micro-Services 
Architecture where a single API gateway serves a single front-end. The API
gateway communicated to one or more micro-service.  

```text
                           API Gateway
   Exchange   +------------------------------------+    Exchange    +-----------------+
 <----------> |  MUX  |  Logic  |  MS Integration  |  <---------->  |  Micro-Service  |
              +------------------------------------+                +-----------------+
                     
```
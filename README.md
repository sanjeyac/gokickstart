# Go Kickstart

Go kickstart is a simple repository that I'm managing to a have a fast setup for Go web application
with my most common use cases using practices that I found useful and easy to maintain.

## Server

`server.go` contains all the dependencies that are normally available on most frameworks like routing, properties, http mangement, security, etc.

## Routes

`routes.go` contains all the mapping between http's routes and controllers in a rails

## Controllers

`controllers.go` will be the bridge between routing and business logic and middleware

## Application properties

Like most application, having configurations and constants in a file make them easier to manage, this is a de-fact standard in Spring-Boot with application.yaml
A `ServerConfig` struct will contain the values of properties.yaml

## Todo next

Integrate common security managers ( basic auth, session based and JWT )

## References

* [Mat Ryers after 8 years](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html)
* [Mat Ryers after 8 years / Talk](https://www.youtube.com/watch?v=8TLiGHJTlig)
* [Spring Boot](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/)
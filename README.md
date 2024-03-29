# Go Kickstart

Go kickstart is a simple repository that I'm managing to a have a fast setup for Go web application
with my most common use cases using practices that I found useful and easy to maintain.

## How to run

clone repository and use ```go run .```

## Server

`server.go` contains all the dependencies that are normally available on most frameworks like routing, properties, http mangement, security, etc.

## Routes

`routes.go` contains all the mapping between http's routes and controllers in a rails, while method mapping has been implemented in a Spring Boot style with GetMapping, PostMapping, PutMapping with a RequestMapping parent.

```
s.GetMapping("/", s.HandleIndex())
```

## Controllers

`controllers.go` will be the bridge between routing and business logic and middleware

## Static assets

static assets will be exposed on /static route, this is a good place for stylesheets, javascripts, images and so on.

## Application properties

Like most application, having configurations and constants in a file make them easier to manage, this is a de-facto standard in Spring-Boot with application.yaml, so 
a `ServerConfig` struct will contain the values of properties.yaml

## Todo next

Template engine
Integrate common security methods ( basic auth, session based and JWT )

## References

* [Mat Ryers after 8 years](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html)
* [Mat Ryers after 8 years / Talk](https://www.youtube.com/watch?v=8TLiGHJTlig)
* [Spring Boot](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/)
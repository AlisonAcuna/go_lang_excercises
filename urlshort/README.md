Jon Calhouns Go Excercise: Url Shortener

https://courses.calhoun.io/lessons/les_goph_04

The goal of this exercise is to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs we would look for any incoming web requests with the path /dogs and redirect them.

Jon's excercises are brilliant prompts to get into golang specific problems but they aren't designed to be highly testable.

I am going to rework the prompt slightly for a more testable challenge.


URL shorteners have two core functionalities(from https://www.milanjovanovic.tech/blog/how-to-build-a-url-shortener-with-dotnet):

Generating a unique code for a given URL
Redirecting users who access the short link to the original URL

NS: Generate test for code exists and figure out how to mock UrlCodes actually having something in it

go test -coverprofile=c.out
go tool cover -html="c.out"
# Paniko

Just a small project to demonstrate how I would build a web application using the functional programming approach in
golang, which means not using any side effect all.

## What was the benefit of using the functional programming approach?

It's made testing a lot more easier as I did not have to worry about a init function getting executed by accident, 
I don't want to accidently read the config file while running the test.  Instead of having an init function, I write 
boot functions that get execute by the
[main function](https://github.com/CJ-Jackson/paniko/blob/master/cmd/paniko/main.go) and does get executed while running the
test.

I like a fact that I can use lamba (closures) to link a controller to the router
(Eg [Here](https://github.com/CJ-Jackson/paniko/blob/master/paniko/www/home.boot.go) and
[Here](https://github.com/CJ-Jackson/paniko/blob/master/paniko/www/errors/error.boot.go)), I get to control what going
on between the router and the controller, you can check if the user is logged or not before you can execute the
controller or even check the csrf token at that point, it's quite nice really.

For testing I found that a combination of creating a custom data type inside a test function (avoiding any namespace
issues), using a lambda with multi-value return and using lambda for each test case, make it's the experience so much
cleaner and easier to read. [Link](https://github.com/CJ-Jackson/paniko/blob/master/paniko/security/user_test.go)

## Would I use FP for future projects?

Hell yeah, I find it's a lot cleaner and less clunky than the OOP approach, the physiology is just a lot better.  It's
awesome that I don't have to worry to much about encapsulation, as the controller and the lambda is sandbox to the
function, that created the closure.  I like the fact that golang struct can be either reference or non-reference, the
latter is great for controller that way I only have to call the
[constructor](https://github.com/CJ-Jackson/paniko/blob/master/paniko/www/home.go#L15)
once, as it's pass by value, every time I call a method it's create a new copy of the struct, it's basically immutable;
therefore avoiding anti-patterns like singletons.

## What does paniko do?

It's automatically docdrops, if I don't report to the system.

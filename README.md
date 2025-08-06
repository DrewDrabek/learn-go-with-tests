# learn-go-with-tests
Used to store learn go with tests items

# Table of contents
- [Hello World](#hello-world)




## Hello World

This was the first lesson that I did as part of the learn go with tests.

I already had go installed so I started with this. 

This lesson went over the strucutre for the main package and what a main func is.

It talked about how to do imports, function structure, as well as how to do basic tests.

This talks about what the go mod init does. How it will store essential information about my project and if I was to ever distrubute the program then this is where dependency information would go and so on.

Testing information:

```text
Writing a test is just like writing a function, with a few rules
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- To use the *testing.T type, you need to import "testing", like we did with fmt in the other file
```

Formatting Information:

While there is more this is really what I care about and I can use most often

```text
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
	(floating-point infinities and NaNs print as ±Inf and NaN)
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value
```

Items I learned here:

- If statements
- Function syntax
- Named returns
- Switch statements
- Unit tests 
- Variable creation
- Const creation
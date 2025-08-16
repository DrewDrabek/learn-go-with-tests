# learn-go-with-tests
Used to store learn go with tests items

# Table of contents
- [Hello World](#hello-world)
- [Integers](#integers)
- [Iterations](#iteration)




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

## Integers

This was the second lesson that I learned here and is mostly about integers and aditional testing practice. 

There were not a lot of additional things here that I leaned but I did also learn about the example functions that can written in the tests. These
are for tests that are mostly for distributed items where we can give example functions that run and pass as the code is currently. These run when 
the tests are called.

Nothing else here this was quick.


## Iteration

This lesson was mainly about how iteration works in go

There are not while, do, until or any other keywords other than `for` in golang. This can do all iteration and makes it easy to work with.

In this lesson I also learned about benchmarking in go lang and the importance of the standard library.

in order to do benchmarking you can use the testing.B object. This will test it on your system and give it its average run time.

The syntax for the for loop is:

```text

for i := 0 ; i < 5; i ++ {
	
}
```

## Arrays/Slices

Learned about the following types:

- Arrays - these are lists that have a fixed length and can only be compared to lists of the same kind 

Decleration : [5]int{1,2,3,4,5}

- Slices - these are lists that do not have a predetermined size. This is what I am going to use most often.

Deceleration: []int{}

- We can also use the make function with slices and other things like maps. It
allows you to specify the default size of the underlying list and the memory allocation. This
can be beneificial for creating slices and other items.

Decleration: sums := make([]int, someCounter)

You can also slice a slice. This is something like [1:] this will take everything starting
at the first character in the slice and go on and Vice versa if it was the other way.

The `len` function is used to get the length of the tool.

I also learned about test tool coverage and reflect.DeepEqual to compare items that might no be normally comparable.
But be careful when using this as its not type safe. 

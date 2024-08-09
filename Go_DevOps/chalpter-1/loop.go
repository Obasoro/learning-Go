// Looping in go

/*. Another comment system in Go  */


for i := 0; i < 10; i++ {
	fmt.Println(i)
}

// i+ is the post statement, it says increment i by 1

var i int
for ;i < 10; i++ {
	fmt.Println(i)
}
fmt.Println("i's final value: ", i)

// Remove the post statement
var i int
for i < 10 {
	i++
}
b := true
for b { // This will loop forever
	fmt.Println("Hello")
}

// Creating an infinite loop

for {
	fmt.Println("hello World")
}

for {
	if err := doSomething(); err != nil {
		break
	}
	fmt.Println("Keep going")
}

for i := 0; i < 10; i++ {
	if i % 2 == 0 { // Only 0 for even numbers
		continue
	}
	fmt.Println("odd number: ", i)
}

// Conditional Command
// if/else block
// switch block

if  x > 2 {
	fmt.Println("x is greater than  2")
}

if err := somePointFunction(); err != nil {
	fmt.Println(err)
}

/*.  else */

if condition {
	function1()
}else {
	function2()
}

if v, err := somePointFunction(); err != nil {
	return err
}else {
	fmt.Println(v)
	return nil
}

// Cleaner Code

v, err := somePointFunction
if err != nil {
	return err
}
fmt.Println(v)
return nil
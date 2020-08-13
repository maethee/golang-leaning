// Interface and Type assertion

interface type is dynamic type

More precisely, if T is not an interface type, x.(T) asserts that the dynamic type of x is identical to the type T. In this case, T must implement the (interface) type of x; otherwise the type assertion is invalid since it is not possible for x to store a value of type T. If T is an interface type, x.(T) asserts that the dynamic type of x implements the interface T.

x will operan type to type that we assign x.(T)

lib to help us to convert int to string
"strconv"

when we useing v, ok := x(T)

no runtime panic occurs and v is the zero value

Type switch

A type switch compares types rather than values. It is otherwise similar to an expression switch. It is marked by a special switch expression that has the form of a type assertion using the reserved word type rather than an actual type:

switch i := x.(type) { we got exactly type from SwitchGuard
case nil:
printString("x is nil") // type of i is type of x (interface{})
case int:
printInt(i) // type of i is int
case float64:
printFloat64(i) // type of i is float64
case func(int) float64:
printFunction(i) // type of i is func(int) float64
case bool, string:
printString("type is bool or string") // type of i is type of x (interface{})
default:
printString("don't know the type") // type of i is type of x (interface{})
}
when we using x.(T) it is Type switch form we can not use fallthrough because can match exactly one.
The "fallthrough" statement is not permitted in a type switch.

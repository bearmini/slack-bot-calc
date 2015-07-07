package main

import (
	"fmt"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/marcmak/calc/calc"
)

const HELP = `
Please specify an expression.
Supported operators: + - * / ^ ( )
Supported functions: Floor Ceil Cbrt Sqrt Sin Cos Tan Asin Acos Atan Ln Abs
Supported constants: E Pi Phi Sqrt2 SqrtE SqrtPi SqrtPhi
Examples:
  1+2*3/4  # => 2.5
  9^2+Tan(Pi-Sqrt2)  # => 74.6658808329578
`
func main() {
	m := martini.Classic()
	m.Get("/", top)
	m.Post("/", top)
	m.Run()
}

func process(text string) (result string, err error) {
	defer func() {
		if e := recover(); e != nil {
			result = ""
			err = e.(error)
		}
	}()

	answerf64 := calc.Solve(text)
	var x int = int(answerf64)
	return fmt.Sprintf("%%v = %v\n%%d = %d, %%b = %b, %%o = %o, %%x = %x, %%U = %U ", answerf64, x, x, x, x, x), nil
}

func top(w http.ResponseWriter, r *http.Request) (int, string) {
	text := r.FormValue("text")
	if len(text) == 0 {
		return 200, HELP
	}

	result, err := process(text)
	if err != nil {
		return 500, err.Error();
	}

	return 200, result
}

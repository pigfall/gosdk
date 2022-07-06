package output

import (
	"fmt"
	"os"
	"github.com/fatih/color"
)

func Err(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

func Errf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func ErrfWithRedColor(format string,args ...interface{}){
	color.New(color.FgRed).Fprintf(os.Stderr,format,args...)
}

func ErrWithRedColor(args ...interface{}){
	color.New(color.FgRed).Fprint(os.Stderr,args...)
}

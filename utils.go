package errors

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

type Trace struct {
	FileName     string
	FunctionName string
	Line         int
}

func getTrace() *Trace {
	// Skip 1 level to get the info of the caller of foo
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Error retrieving caller information")
		return nil
	}

	// Retrieve the details of the function
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fmt.Println("Error retrieving function information")
		return nil
	}

	// Extracting the short function name (without package path)
	shortFnName := fn.Name()
	lastIndex := strings.LastIndex(shortFnName, "/")
	if lastIndex >= 0 {
		shortFnName = shortFnName[lastIndex+1:]
	}
	dotIndex := strings.Index(shortFnName, ".")
	if dotIndex >= 0 {
		shortFnName = shortFnName[dotIndex+1:]
	}

	// Trimming the file path to get a cleaner output (optional)
	shortFile := file
	if len(file) > 0 {
		shortFileParts := strings.Split(file, "/src/")
		if len(shortFileParts) > 1 {
			shortFile = "/src/" + shortFileParts[1]
		}
	}

	log.Printf("[%s] %s:%d\n", shortFile, shortFnName, line)

	// Output: [main.go] main.main:10

	resp := &Trace{
		FileName:     shortFile,
		FunctionName: shortFnName,
		Line:         line,
	}

	return resp
}

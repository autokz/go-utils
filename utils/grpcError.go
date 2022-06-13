package utils

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const ErrMsgSeparator = "; "
const stackTracePrefixMessage = "stack trace: "

type WrapperError struct {
	Code       uint32
	Message    string
	StackTrace []string
}

func (w *WrapperError) Error() string {
	stackTrace := stackTracePrefixMessage + strings.Join(w.StackTrace, ", ")

	return w.Message + ErrMsgSeparator + stackTrace
}

func GrpcErrorWrapper(err error) error {
	parsedErr := ParseGrpcError(err)
	return GrpcError(parsedErr.Code, parsedErr.Message)
}

func GrpcError(code uint32, message string) error {
	// Parse nested wrappers...
	parsedErr := ParseGrpcError(errors.New(message))

	// Creating stack...
	stack := "unknown"

	_, file, no, ok := runtime.Caller(1)
	//details := runtime.FuncForPC(pc)
	//if ok && details != nil {
	if ok {
		//funcName := strings.Split(details.Name(), ".")
		filePath := file
		codeLine := strconv.Itoa(no)

		//stack = filePath + "::" + funcName[len(funcName)-1] + "#" + codeLine
		stack = filePath + "#" + codeLine
	}

	// Pushing stack to stackTrace...
	parsedErr.StackTrace = append([]string{stack}, parsedErr.StackTrace...)

	return status.Error(codes.Code(code), parsedErr.Error())
}

func ParseGrpcError(err error) *WrapperError {
	var code uint32
	var msg string
	var stackTrace []string

	splitErr := strings.Split(err.Error(), ErrMsgSeparator)

	// Code
	code = 9999 // Fallback
	codeString := FindInStringByRange(splitErr[0], " code = ", " desc =")

	var reQuantity = regexp.MustCompile(`[^\d.]`)
	codeString = reQuantity.ReplaceAllString(codeString, "")

	convertedCode, _ := strconv.Atoi(codeString)
	code = uint32(convertedCode)

	// Msg
	msg = FindInStringByRange(splitErr[0], " desc =", ";")
	if msg == "" {
		msg = err.Error()
	}

	// StackTrace
	if len(splitErr) > 1 {
		stringStackTrace := strings.Split(splitErr[1], stackTracePrefixMessage)
		if len(stringStackTrace) > 1 {
			stackTrace = strings.Split(stringStackTrace[1], ", ")
		}
	}

	return &WrapperError{
		Code:       code,
		Message:    strings.TrimSpace(msg),
		StackTrace: stackTrace,
	}
}

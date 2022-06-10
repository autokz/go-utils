# go-utils:

### GrpcError:
```
err1 := utils.GrpcError(100, "One")
err2 := utils.GrpcError(200, err1.Error())
err3 := utils.GrpcError(300, err2.Error())

fmt.Print(err3) 
// Output: rpc error: code = Code(300) desc = One; stack trace: /.../handle.go#3, /.../handle.go#2, /.../handle.go#1

```

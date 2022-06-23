# go-utils:

### GrpcError:
```go
err1 := utils.GrpcError(100, "One")
err2 := utils.GrpcError(200, err1.Error())
err3 := utils.GrpcError(300, err2.Error())
// and you may:
err4 := utils.GrpcErrorWrapper(er3)


fmt.Print(err3) 
// Output: rpc error: code = Code(300) desc = One; stack trace: /.../handle.go#3, /.../handle.go#2, /.../handle.go#1

fmt.Print(err4) 
// Output: rpc error: code = Code(300) desc = One; stack trace: /.../handle.go#4, /.../handle.go#3, /.../handle.go#2, /.../handle.go#1

```


### Pagination:

```protobuf
message Pagination {
  uint32 page = 1;
  int32 limit = 2;
  uint32 offset = 3;
}
```

```go
func Handle(ctx context.Context, request *pb.GetAll) (*pb.Items, error) {
	if request.Pagination != nil {
		documentReq.Pagination = &documentPb.PaginationRequest{
			Page:   request.Pagination.Page,
			Limit:  request.Pagination.Limit,
			Offset: request.Pagination.Offset,
		}
	}

	countAll, err := rep.GetCountAll(ctx)
	if err != nil {
		return nil, err
	}

	pager := utils.CreatePagination(req.Pagination, countAll)

	items, err := rep.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	//...
}
```
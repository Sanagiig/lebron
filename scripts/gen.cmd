
@REM rpc
goctl rpc protoc ./apps/order/rpc/order.proto --go_out=./apps/order/rpc --go-grpc_out=./apps/order/rpc --zrpc_out=./apps/order/rpc --style=go_zero
goctl rpc protoc ./apps/product/rpc/product.proto --go_out=./apps/product/rpc --go-grpc_out=./apps/product/rpc --zrpc_out=./apps/product/rpc --style=go_zero
goctl rpc protoc ./apps/reply/rpc/reply.proto --go_out=./apps/reply/rpc --go-grpc_out=./apps/reply/rpc --zrpc_out=./apps/reply/rpc --style=go_zero


@REM SQL gen
@REM product
goctl model mysql datasource --url "root:lwj@1993@tcp(127.0.0.1:3306)/lebron" --table "product,product_operation,category" --dir="./apps/product/rpc/model" --style=go_zero
goctl model mysql datasource --url "root:lwj@1993@tcp(127.0.0.1:3306)/lebron" --table "reply" --dir="./apps/reply/rpc/model" --style=go_zero

https://www.kancloud.cn/linimbus/go-micro/529025


<!-- pb -->
 protoc \
-I/usr/local/include \
-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--proto_path=. \
--go_out=. \
 proto/hello_http/hello_http.proto


 <!-- gw -->
protoc -I . \
-I /usr/local/include \
-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--proto_path=. \
--go_out=. \
--grpc-gateway_out=logtostderr=true:.  \
proto/hello_http/hello_http.proto


protoc -I. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis --grpc-gateway_out=logtostderr=true:. gateway.proto
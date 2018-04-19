#!/usr/bin/env sh

# run protobuf compiler on dir up
cd ..

echo "building bitbucket model proto files"
protoc --go_out=plugins=grpc:models/bitbucket/pb -I=models/bitbucket/ \
  -I$GOPATH/src \
  models/bitbucket/*.proto

echo "building root model proto files"
protoc --go_out=plugins=grpc:models/pb/ -I=models \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  --grpc-gateway_out=logtostderr=true:models/pb \
  --swagger_out=logtostderr=true:models/pb \
  models/*.proto

echo "injecting custom tags"
# inject our custom tags into build protobuf
protoc-go-inject-tag -input=models/pb/guideocelot.pb.go
protoc-go-inject-tag -input=models/pb/build.pb.go

# then we have to run go get in the stub directory cause grpc gateway ¯\_(ツ)_/¯ does this even work
cd models/pb
go get .
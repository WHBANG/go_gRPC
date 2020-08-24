# go_gRPC

1.安装Protobuf编译器protoc。<br>
    查看protoc --version
    
2.获取插件支持库 <br>
    // gRPC运行时接口编解码支持库 <br>
    go get -u github.com/golang/protobuf/proto <br>
    // 从 Proto文件(gRPC接口描述文件) 生成 go文件 的编译器插件  <br>
    go get -u github.com/golang/protobuf/protoc-gen-go <br>
    
3.写好 .proto 描述文件定义 RPC 的接口，然后用 protoc（带 gRPC 插件）基于 .proto 模板自动生成客户端和服务端的接口代码。 <br>
  //protoc --go_out=plugins=grpc:{输出目录} {.proto文件} 例子：protoc --go_out=plugins=grpc:./test/ ./test.proto <br>
  
4.服务端 <br>
  1 创建监听套接字：lis, err := net.Listen("tcp", port) <br>
  2 创建服务端：grpc.NewServer() <br>
  3 注册服务：pb.RegisterHelloServiceServer() <br>
  4 启动服务端：s.Serve(lis)  <br>
  
5.客户端<br>
  1 创建连接 conn,err := net.Dial(address,grpc.WithInsecure()) <br>
  2 创建客户端 c := pb.NewHelloServiceClient(conn) <br>
  3 调用接口  c.SayHello(context.Background(),&pb.HelloRequest{})   <br>

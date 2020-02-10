# Remark 
protoc version 3.9.2

``` bash 
Packages\protoc\bin\protoc.exe -I=.\Assets\Protos -I=Packages\protoc\include --csharp_out=.\Assets\Protos .\Assets\Protos\GameCtl.proto --grpc_out=.\Assets\Protos --plugin=protoc-gen-grpc=%proto_path%\bin\grpc_csharp_plugin.exe 

```
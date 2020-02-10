# Unlight-Zwei
Whole project maintain version controlling


```
flatc -g Fbs\skillb.fbs
```

flatc -g myschema.fbs





```cmd 
protoc -I=.\Protos -I=%proto_path%\include -I=%proto_path%\googleapis --csharp_out=.\Protos .\Protos\*.proto --grpc_out=.\Protos --plugin=protoc-gen-grpc=%proto_path%\bin\grpc_csharp_plugin.exe 
```
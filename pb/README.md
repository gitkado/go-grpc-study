```sh
# protoファイルからGo言語ベースのコードを自動生成
docker exec -it go-grpc-study_api_1 protoc --go_out=plugins=grpc:. pb/service.proto
```

[qiita: .proto ファイルからserver、client,interface等のコードを生成する](https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3#proto-%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%8B%E3%82%89serverclientinterface%E7%AD%89%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E7%94%9F%E6%88%90%E3%81%99%E3%82%8B)

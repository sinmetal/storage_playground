# storage_playground

Google Cloud Storage [Customer-supplied encryption keys](https://cloud.google.com/storage/docs/encryption#customer-supplied) Sample

## Usage

```
go run *.go --cmd printkey --bucket hoge --object hello.txt
go run *.go --cmd upload --bucket hoge --object hello.txt
go run *.go --cmd download --bucket hoge --object hello.txt
```
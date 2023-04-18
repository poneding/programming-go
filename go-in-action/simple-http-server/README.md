# simple-http-server

## Run

```bash
cd ./in-action/simple-http-server
go run main.go
```

## Test

```bash
curl localhost:8080      
<pre>
<a href="hello.log">hello.log</a>
</pre>

curl localhost:8080/hello
Hello World!

curl "localhost:8080/form?name=jaychou&address=china"
Post request successful!
Name: jaychou
Address: china
```

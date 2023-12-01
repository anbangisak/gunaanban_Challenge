# coding challenge

To build server use below command:
```
go build
```

To run server use below command:
```
./gunaanban_Challenge.exe -e development
```

To verify dashboard page use below command or url in browser:
```
curl -v http://localhost:8080/dashboard
curl -v https://localhost:8200/dashboard
```

To build docker image:
```
cd deploy/docker;
docker build -t "challenge:0.0.1" .
```
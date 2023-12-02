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
./build.sh
docker run -it -d -p 0.0.0.0:8200:8200 -p 0.0.0.0:8080:8080 challenge:0.0.2
```

To deploy docker image in kubernetes:
```
cd deploy/bundle
ansible-playbook -i hosts deploy.yml
```

To run json transformer:
```
cd challenge2
go run .\jsonTransformer.go
```

To run HackerRank credit card number validation:
```
cd additional_challenge
go run .\regexHackerRank.go
python .\regexHackerRank.py
```
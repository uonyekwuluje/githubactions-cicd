# Simple Hello World Java App

## Maven Compile
```
mvn compile install package
```

## Create Docker Image
```
docker build -t helloworldapp .
```

## Test App
Run the command below:
```
docker run -p 5000:8080 helloworldapp 
```
Open your browser and type
```
http://localhost:5000
```

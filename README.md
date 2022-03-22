# Coin market api-integration
1. Public endpoint for ad-hoc request
2. AWS CloudWatch to make a scheduler
3. Postgres as data storage

##  Architecture Diagram
![Architecture Diagram](https://github.com/Aibier/coinmarket-service/blob/main/diagram.png)
##  Run with Docker

1. **Build**
```shell script
make build
docker build . -t api-rest 
```
2. **Run**
```shell script
make run
```
3. **Test**
```shell script
make rest
```
4. **Scan**
```shell script
make scan
```
5. To deploy (1)
```shell script
1. GOOS=linux GOARCH=amd64 go build -o main main.go
2. zip -r ./main.zip  *
3. Upload to labmda function
```
5. To deploy (2) via aws-cli
```shell script
1. Build Image: docker build -t coinmarket-[version] .
2. Create RCR Repo: aws ecr create-repository --repository-name coinmarket-[version] --image-scanning-configuration scanOnPush=true
3. Tag image to repo: docker tag  coinmarket-[version]:latest [ecr_repo_id].dkr.ecr.ap-southeast-1.amazonaws.com/coinmarket-[version]:latest
4. Login to ECR: aws ecr get-login-password | docker login --username AWS --password-stdin [ecr_repo_id].dkr.ecr.ap-southeast-1.amazonaws.com
5. Push Image to ECR: docker push [ecr_repo_id].dkr.ecr.ap-southeast-1.amazonaws.com/coinmarket-[version]:latest
6. To invoke: aws lambda invoke --function-name coinmarket  output.txt 
```
6. Add env variables
```shell script
export COIN_MARKET_TOKEN=
export COIN_MARKET_URL=
export SEARCH_PATH=
export COINS=btc,eth,sol,xtz
```

## Notes
1. Create postgres database and user in your local env.
2. Make sure you setup file watchers in your idea.
3. Before make a git push , please make sure you run ```make scan``` and ``` make test```
#### Contributed by: Aspire Trust and Security Team
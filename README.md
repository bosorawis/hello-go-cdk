# Hello Go CDK!

This is a simple API gateway serverless Go project using CDK

## Design

Two main parts are 
 - [Go App](./app) - this is where all the application code goes
 - [CDK infrastructure](./infrastructure) - this is where the CDK stack code goes

## Requirements
- [Go 1.16+](https://golang.org/dl/)
- [node 14](https://nodejs.org/en/download/)

## Deploying a sample app

```bash
npm install
npx cdk bootstrap # <<- optional if this has already been done
make all
npx cdk deploy
```
## Clean up
```bash
npx cdk destroy
```

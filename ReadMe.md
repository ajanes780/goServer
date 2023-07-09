# Go Blog Server 

## Serves markdown files as html

### Usage

```bash
// starts server on port 8080 and will watch for changes in server files
air

// to watch for changes in scss files
sass --watch styles:static/css

// to build for production
make run 

```

## Deploying Articles 

```go
// to deploy articles
go run deploy.go

```
- read markdown files from `articles` directory
- parse markdown files to html
- images should be saved to S3 
- create Articles type from parsed markdown
- save Articles to database

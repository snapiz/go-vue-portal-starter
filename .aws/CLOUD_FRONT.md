## Settings
default root object: index.html

## Origins

### Lambda
origin: api entrypoint
origin path: /Prod
protocol: HTTPS only

### Frontend
origin: select s3-bucket

## Behaviors
auth/*: auth-api, Allow all HTTP methods, all cookies
graphql: main-api, Allow all HTTP methods, all cookies
{sevice name}/graphql: main-api, Allow all HTTP methods, all cookies

> #apply for all services
{sevice name}/*.js*: s3-bucket-of-service, compress, use lambda edge (edges.js)
{sevice name}/locales/*: s3-bucket-of-service, compress, use lambda edge (edges.js)
{sevice name}/assets/*: s3-bucket-of-service, compress

Default(*): s3-portal, compress, redirect HTTPS

#Errors pages

403: /index.html 200
404: /index.html 200
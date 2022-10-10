### POST Message

Form
```
curl -X POST http://localhost:3000/sendMessage \
  --data-urlencode 'message=Hello'  \
  --data-urlencode 'from=mimoc'  \
  --data-urlencode 'to=mu' \
  --data-urlencode 'token=1234'
```

Json
```
curl -X POST http://localhost:3000/sendMessage \
  -H 'content-type: application/json' \
  -d '{ "message": "Hello!", "from": "mimoc", "token": "1234" }'
```

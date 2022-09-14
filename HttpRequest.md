### POST Message

```
curl -X POST http://localhost:3000/sendMessage \
  --data-urlencode 'message=Hello'  \
  --data-urlencode 'from=mimoc'  \
  --data-urlencode 'to=mu' \
  --data-urlencode 'token=123'
```


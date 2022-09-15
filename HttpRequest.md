### POST Message

Form
```
curl -X POST http://localhost:3000/sendMessage \
  --data-urlencode 'message=Hello'  \
  --data-urlencode 'from=mimoc'  \
  --data-urlencode 'to=mu' \
  --data-urlencode 'token=123'
```

Json
```
curl -X POST http://localhost:8080/postjson \
  -H 'content-type: application/json' \
  -d '{ "field_str": "文字だ", "field_int": 12, "field_bool": true }'
```

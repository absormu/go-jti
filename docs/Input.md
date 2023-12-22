## INPUT

### GET URL: /providers
* REQUEST
```
curl --location 'http://localhost:9670/go-jti/api/v1/providers' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yYWx2b3JkMDdAZ21haWwuY29tIiwiZXhwIjoxNzAzMzA0MTgzLCJuYW1lIjoiTXVoYW1hZCBVbGlsIEFic29yIiwidWlkIjoiY20yZ2t0dTNudDRrc2czNjM4azAiLCJ1c2VyX2lkIjoxfQ.UrVXh3VCkRv2_vDvuS0WvY7qgrwqVDaqimxN9xtyeyE'
```
* RESPONSE
```
{
      "status": 0,
      "message": {
            "Indonesian": "Sukses",
            "English": "Success"
      },
      "data": [
            {
                  "id": 1,
                  "code": "",
                  "name": "TELKOMSEL"
            },
            {
                  "id": 2,
                  "code": "",
                  "name": "INDOSAT"
            },
            {
                  "id": 3,
                  "code": "",
                  "name": "XL"
            },
            {
                  "id": 4,
                  "code": "",
                  "name": "TRI"
            },
            {
                  "id": 5,
                  "code": "",
                  "name": "SMARTFEN"
            }
      ]
}
```
### POST URL: /number-phone
* REQUEST
```
curl --location 'http://localhost:9670/go-jti/api/v1/number-phone' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yYWx2b3JkMDdAZ21haWwuY29tIiwiZXhwIjoxNzAzMzA0MTgzLCJuYW1lIjoiTXVoYW1hZCBVbGlsIEFic29yIiwidWlkIjoiY20yZ2t0dTNudDRrc2czNjM4azAiLCJ1c2VyX2lkIjoxfQ.UrVXh3VCkRv2_vDvuS0WvY7qgrwqVDaqimxN9xtyeyE' \
--data '{
      "number": "081291808445", 
      "provider": {
            "id": 1
      }
}'
```
* RESPONSE
```
{
      "status": 0,
      "message": {
            "Indonesian": "Sukses",
            "English": "Success"
      }
}
```
### GET POST: /auto-number-phones
 * REQUEST
```
curl --location --request POST 'http://localhost:9670/go-jti/api/v1/auto-number-phones' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yYWx2b3JkMDdAZ21haWwuY29tIiwiZXhwIjoxNzAzMzAxMjgyLCJuYW1lIjoiTXVoYW1hZCBVbGlsIEFic29yIiwidWlkIjoiY20yZnU4bTNudDRsMzY0N2R2Y2ciLCJ1c2VyX2lkIjoxfQ.qWeY-4oy8tZkBUYKsFEp8yOn7FtoiK7s4w6qfeaUu1I'
```
* RESPONSE
```
{
      "status": 0,
      "message": {
            "Indonesian": "Sukses",
            "English": "Success"
      }
}
```
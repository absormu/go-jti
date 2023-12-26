## INPUT

### GET URL: /providers
* REQUEST
```
curl --location 'http://localhost:9670/go-jti/api/v1/providers' \
--header 'Authorization: Bearer '
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
--header 'Authorization: Bearer ' \
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
--header 'Authorization: Bearer '
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

## OUTPUT
### GET URL: /number-phones
* REQUEST
```
curl --location 'http://localhost:9670/go-jti/api/v1/number-phones?type=2' \
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
                "number": "085703218798",
                "type": 2,
                "provider": {
                        "id": 2,
                        "code": "0857",
                        "name": "INDOSAT"
                },
                "is_deleted": 0,
                "created_at": "2023-12-22 11:19:58",
                "created_by": "Muhamad Ulil Absor",
                "modified_at": "",
                "modified_by": ""
            } 
        ]
    }
```
### GET URL: /number-phone/:id
* REQUEST
```
curl --location 'http://localhost:9670/go-jti/api/v1/number-phone/1' \
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
        {
            "id": 1,
            "number": "081291808448",
            "type": 2,
            "provider": {
                  "id": 1,
                  "code": "",
                  "name": "TELKOMSEL"
            },
            "is_deleted": 0,
            "created_at": "2023-12-21 20:42:24",
            "created_by": "System",
            "modified_at": "2023-12-22 11:19:05",
            "modified_by": "Muhamad Ulil Absor"
        }
    }
```

### PUT URL: /number-phone/:id
* REQUEST
```
curl --location --request PUT 'http://localhost:9670/go-jti/api/v1/number-phone/1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer ' \
--data '{
      "number": "081291808448", 
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
```
    {
      "status": 7,
      "message": {
            "Indonesian": "Data tidak tersedia",
            "English": "Data not found"
      }
    }
```
### DELETE URL: /number-phone/:id
* REQUEST
```
curl --location --request DELETE 'http://localhost:9670/go-jti/api/v1/number-phone/1' \
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
```
    {
      "status": 7,
      "message": {
            "Indonesian": "Data tidak tersedia",
            "English": "Data not found"
      }
    }
```
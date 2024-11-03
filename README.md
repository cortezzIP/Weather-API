# Weather-API
## Documentation
### Get Current Weather

This endpoint retrieves the current weather for a specific location.

#### Request

- Method: GET
    
- URL: `localhost:8080/api/weather/current`
    
- Query Parameters:
    
    - location (string, required): The location for which the weather is to be retrieved, e.g. "Moscow"
        

#### Response

- Status: 200
    
- Content-Type: application/json

- Body:
```json
{
    "location": {
        "name": "Moscow",
        "country": "Russia"
    },
    "current": {
        "condition": {
            "text": "Partly cloudy"
        },
        "temp_c": 0.2
    }
}
```

## Architecture visualization
![Weather API architecture visualisation](https://github.com/cortezzIP/Weather-API/blob/main/docs/assets/images/weather-api-architecture-visualisation.png)

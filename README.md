# Description

<p>This is the backend service for wishlist saving app. I am build a endpoint for client to have experience of saving money from your own wishlist. This endpoint was build to make it anybody easier to manage their own money to buy something what they want. This backend service has 2 features that are (login and register) user and (Create Read Update and Delete) wishlist. I use Google Cloud VM Instance for deployment and using CloudSQL for database service. Using some clean architecture on the project structure and using CI/CD with github action was made easier for deployment for this project. 
</p>

# Documentation

## User

### Sign Up

- Method: POST
- Url: /v1/signup

```json
{
  "name": "string",
  "username": "string",
  "email": "string",
  "password": "string"
}
```

### Sign In

- Method: POST
- Url: /v1/login

```json
{
  "username": "string",
  "password": "string"
}
```

## Wishlist

### Get All Wishlists

- Method: GET
- Url: /v1/wishlist
- Header: Bearer token

```json
{
  "message": "string",
  "data": [
    {
      "WishlistId": "string",
      "WhislistName": "string",
      "TargetMoney": "int",
      "TargetMonth": "int",
      "CreatedAt": "string",
      "UpdatedAt": "string",
      "DeletedAt": null,
      "IsFinish": "string",
      "UserId": "string",
      "BalanceId": {
        "BalanceId": "string",
        "AmmountMoney": "int",
        "ExceedMoney": "int",
        "CountSave": "int",
        "CreatedAt": "string",
        "UpdatedAt": "string",
        "WishlistId": "string",
        "HistoryBalances": [
          {
            "HistoryBalanceId": "string",
            "SavingMoney": "int",
            "Status": "string",
            "CreatedAt": "string",
            "BalanceIdHistory": "string"
          }
        ]
      }
    }
  ]
}
```

### Get Wishlist

- Method: GET
- Url: /v1/wishlist/:wishlistid
- Header: Bearer token

```json
{
  "message": "string",
  "data": {
    "WishlistId": "e1b8cfed-61e8-44fa-968c-68c4754436f5",
    "WhislistName": "Sepatu Converse",
    "TargetMoney": 500000,
    "TargetMonth": 4,
    "CreatedAt": "2022-11-07T08:55:12.933Z",
    "UpdatedAt": "2022-11-07T08:55:12.933Z",
    "DeletedAt": null,
    "IsFinish": "onprogress",
    "UserId": "f94ccab1-fbfe-4573-89d2-e0b91b18e242",
    "BalanceId": {
      "BalanceId": "993ff2b3-e978-4d36-a0c5-fe31b344b265",
      "AmmountMoney": 0,
      "ExceedMoney": 0,
      "CountSave": 0,
      "CreatedAt": "2022-11-07T08:55:12.945Z",
      "UpdatedAt": "2022-11-07T08:55:12.945Z",
      "WishlistId": "e1b8cfed-61e8-44fa-968c-68c4754436f5",
      "HistoryBalances": [
        {
          "HistoryBalanceId": "64843456-3797-49c3-bbf5-b8442234f640",
          "SavingMoney": 0,
          "Status": "success",
          "CreatedAt": "2022-11-07T08:55:12.957Z",
          "BalanceIdHistory": "993ff2b3-e978-4d36-a0c5-fe31b344b265"
        }
      ]
    }
  }
}
```

### Get Recommendation From Wishlist

- Method: GET
- Url: /v1/wishlist/recommend/:wishlistid
- Header: Bearer token

```json
{
  "message": "string",
  "data": [
    {
      "WishlistId": "string",
      "Name": "string",
      "Insufficient": "int",
      "CountRecommend": "int",
      "ResponseRecommend": "string"
    }
  ]
}
```

### Create Wishlist

- Method: POST
- Url: /v1/wishlist
- Header: Bearer token

```json
{
  "wishlistname": "string",
  "targetmoney": "int",
  "targetmonth": "int"
}
```

### Edit Wishlist

- Method: PUT
- Url: /v1/wishlist/:wishlistid
- Header: Bearer token

```json
{
  "wishlistname": "string",
  "targetmoney": "int",
  "targetmonth": "int"
}
```

### Saving Money to Balance Wishlist

- Method: PUT
- Url: /v1/wishlist/balance/:wishlistid
- Header: Bearer token

```json
{
  "savingmoney": "int"
}
```

### Delete Wishlist

- Method: DELETE
- Url: /v1/wishlist/:wishlistid
- Header: Bearer token

```json
{
  "message": "string"
}
```

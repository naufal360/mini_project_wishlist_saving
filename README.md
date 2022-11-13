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

### Login

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
- Url: /v1/wishlists
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
- Url: /v1/wishlists/:wishlistid
- Header: Bearer token

```json
{
  "message": "string",
  "data": {
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
}
```

### Get Recommendation From Wishlist

- Method: GET
- Url: /v1/wishlists/:wishlistid/recommend
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
- Url: /v1/wishlists
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
- Url: /v1/wishlists/:wishlistid
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
- Url: /v1/wishlists/:wishlistid/balances
- Header: Bearer token

```json
{
  "savingmoney": "int"
}
```

### Delete Wishlist

- Method: DELETE
- Url: /v1/wishlists/:wishlistid
- Header: Bearer token

```json
{
  "message": "string"
}
```

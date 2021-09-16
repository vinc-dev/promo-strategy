# Promo

To run unit tests and build the code:  `./script/builder.sh`

Run on Local : `go run ./cmd/main.go`

HTTP Route: `POST: http://localhost:8080/checkout`

This code built with some assumptions:
* Assuming that promotions can be applied more than 1 time.
* Assuming that bonus products need to be included in scanned products.
* Assuming Product Names is unique and can be also used as identifier along side with SKU.
* Repository supposed to be used to connect with Database but instead be used to from global variable that acted as master data.

GQL Schema:

```
type Query {
    carts: [Cart]
    cart(id: String!): Cart
}

type Mutation {
    addToCart(payload: AddToCartPayload!): SuccessResponse!
    checkout(cartId: String!): SuccessResponse!
}

type Cart {
    id: String
    sku: String
    qty: Number
}

type AddToCartPayload {
    sku: String!
    qty: Number!
}

type SuccessResponse {
    success: Boolean!
}
```
Fulfillment
===========
CreateFulfillment

```go
```

CompleteFulfilment

```go
s := goshopify.New("https://%s.myshopify.com")
creds := &goshopify.Credentials{ StoreId: "some-store-id", OauthToken: "some-oauth-token" }

f := s.CompleteFulfillment("some-order-id", creds, nil)
fmt.Printf("Fulfillment %v for order #%v has been marked complete", f.Id, f.OrderId)
```

GetFulfilments

```go
s := goshopify.New("https://%s.myshopify.com")
creds := &goshopify.Credentials{ StoreId: "some-store-id", OauthToken: "some-oauth-token" }

fulfillments := s.GetFulfillments("some-order-id", creds, nil)
for _, f := range fulfillments{
	fmt.Printf("Company: %v\nStatus: %v\n", f.TrackingCompany, f.Status)
}
```

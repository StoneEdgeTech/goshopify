Customer
=======
GetCustomers

```go
s := goshopify.New("https://%s.myshopify.com")
creds := &Credentials{ StoreId: "some-store-id", OauthToken: "some-oauth-token" }

custs := s.GetCustomers(creds, nil)
for _, cust := range custs {
  fmt.Printf( "Email: %v\n", cust.Email)
}
```

GetCustomersCount

```go
s := goshopify.New("https://%s.myshopify.com")
creds := &Credentials{ StoreId: "some-store-id", OauthToken: "some-oauth-token" }

c := s.GetCustomersCount(creds, nil)
fmt.Printf("Number of Customers: %v\n", c.Count)
```

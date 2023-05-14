# Stripe Golang API

Created an api using the golang's gin framework for payment processing through stripe. Easy to integrate with the application.

## URL
API end point 
```
http://localhost:8080/charge
```
## API Body
API body parameters are below

```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "amount": 1000,
  "currency": "usd",
  "desc": "Test payment",
  "line1": "Primary address",
  "line2": "Secondary address ",
  "city": "Islamabad",
  "state": "ICT",
  "postal_code": "45400",
  "country": "Pakistan",
  "card":{
    "number":"4242424242424242",
    "exp_month": "12",
    "exp_year": "2024",
    "cvc": "123"
  }
}
```
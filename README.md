# Mock Payment Processor

Mock Payment Processor written in GO for simulate the lifecycle of a transaction.

<br>

## Usage

From the repository root directory, generate the static HTML export of the Next.js
app, and build the Go binary:

```sh
$ cd nextjs
$ npm install
$ npm run export
$ cd ..
$ go build .
```

Then run the binary:

```sh
$ ./main

2022/04/01 14:55:38 Starting HTTP server at http://localhost:8080 ...
```

## API Docs

Documentation of the API

<br>
<hr />

### Healthcheck
Informs about the API's condition
> Request Method: `GET` <br>
> Request URL: `/healthcheck`

<br>
<hr />

### Login
Retrieve the access token
> Request Method: `POST` <br>
> Request URL: `/login`

| Name | Type | Description | Required |
| ------ | ------ | ------ | ------ |
| username | `string` | Name of the user | `true` |

<br>
<hr />

### Get Transaction
Retrieve a transaction by ID
> Request Method: `GET` <br>
> Request URL: `/api/transaction/{id}`

<br>
<hr />

### Delete Transaction
Delete a transaction by ID
> Request Method: `DELETE` <br>
> Request URL: `/api/transaction/{id}`

<br>
<hr />

### Settle Transaction
Set the transaction status to settled
> Request Method: `POST` <br>
> Request URL: `/api/transaction/{id}/settle`

<br>
<hr />

### Reject Transaction
Set the transaction status to rejected
> Request Method: `POST` <br>
> Request URL: `/api/transaction/{id}/reject`

<br>
<hr />

### Get All
Retrieve all transactions
> Request Method: `GET` <br>
> Request URL: `/api/transactions`

<br>
<hr />

### Delete All
Delete all transactions
> Request Method: `DELETE` <br>
> Request URL: `/api/transactions`

<br>
<hr />

### Credit Card
Create a credit card transaction
> Request Method: `POST` <br>
> Request URL: `/api/transaction/creditcard`

| Plugin | Type | Description | Required |
| ------ | ------ | ------ | ------ |
| amount | `uint64` | Amount | `true` |
| payment_method | `object` | Method of the payment | `true` |
| billing_address | `object` | Address of costumer | `true` |

<br>
<hr />

### ACH
Create an ach transaction
> Request Method: `POST` <br>
> Request URL: `/api/transaction/ach`

| Plugin | Type | Description | Required |
| ------ | ------ | ------ | ------ |
| amount | `uint64` | Amount | `true` |
| payment_method | `object` | Method of the payment | `true` |
| billing_address | `object` | Address of costumer | `true` |

<br>
<hr />

### Status codes
<br>

| Message | Code |
| ------ | ------ |
| Success | `100` |
| Limit exceeded | `201` |
| Card blocked | `202` |
| Daily limit exceeded | `203` |
| Fraud detected | `204` |
| Error occured | `206` |

<br>
<hr />

### Example transaction
<br>

```json
{
    "amount": 1000,
    "payment_method": {
        "credit_card": {
            "card_number": "4111111111111111",
            "holder_name": "John Doe",
            "exp_date": "05/25",
            "cvc": "444"
        }
    },
    "billing_address": {
        "first_name": "John",
        "last_name": "Doe",
        "postal_code": "1111",
        "city": "Szeged",
        "address_line_1": "Example street 69.",
        "email": "example@github.com",
        "phone": "5555555555"
    }
}
```

<br>
<hr />

### Example response
<br>

```json
{
    "status": "success",
    "status_code": 100,
    "message": "success",
    "data": {
        "id": "c91khre49b3jg1lbbje0",
        "status": "pending_settlement",
        "payment_method_type": "creditcard",
        "payment_method": {
            "credit_card": {
                "card_number": "411111******1111",
                "holder_name": "John Doe",
                "exp_date": "05/25",
                "cvc": "444"
            }
        },
        "amount": 1000,
        "billing_address": {
            "first_name": "John",
            "last_name": "Doe",
            "postal_code": "1111",
            "city": "Szeged",
            "address_line_1": "Example street 69.",
            "email": "example@github.com",
            "phone": "5555555555"
        },
        "created_at": "2022-03-29T19:59:09.589402+02:00",
        "updated_at": "2022-03-29T19:59:09.589402+02:00"
    }
}
```

# Mock Payment Processor
<hr />

Mock Payment Processor written in GO for simulate the lifecycle of a transaction.

## Usage
<hr />

In the <code>root</code> directory: <br>
<code>go run main.go</code>

In the <code>nextjs</code> directory: <br>
<code>npm install</code> <br>
<code>npm run dev</code>

## API Docs
<hr />

<details>
<summary> Healthcheck </summary>

*Informs about the API's condition*

> Request Method: `GET` <br>
> Request URL: `/healthcheck`
</details>

<details>
<summary> Login </summary>

*Retrieve the access token*

> Request Method: `GET` <br>
> Request URL: `/login`
</details>

<details>
<summary> Get Transaction </summary>

*Retrieve a transaction by ID*

> Request Method: `GET` <br>
> Request URL: `/api/transaction/{id}`
</details>

<details>
<summary> Delete Transaction </summary>

*Delete a transaction by ID*

> Request Method: `DELETE` <br>
> Request URL: `/api/transaction/{id}`
</details>

<details>
<summary> Settle Transaction </summary>

*Set the transaction status to settled*

> Request Method: `POST` <br>
> Request URL: `/api/transaction/{id}/settle`
</details>

<details>
<summary> Reject Transaction </summary>

*Set the transaction status to rejected*

> Request Method: `POST` <br>
> Request URL: `/api/transaction/{id}/reject`
</details>

<details>
<summary> Get All Transaction </summary>

*Retrieve all transactions*

> Request Method: `GET` <br>
> Request URL: `/api/transactions`
</details>

<details>
<summary> Delete All Transaction </summary>

*Delete all transactions*

> Request Method: `DELETE` <br>
> Request URL: `/api/transactions`
</details>

<details>
<summary> Credit Card Transaction </summary>

*Create a credit card transaction*

> Request Method: `POST` <br>
> Request URL: `/api/transaction/creditcard`

| Plugin          | Type     | Description           | Required |
| --------------- | -------- | --------------------- | -------- |
| amount          | `uint64` | Amount                | `true`   |
| payment_method  | `object` | Method of the payment | `true`   |
| billing_address | `object` | Address of costumer   | `true`   |

*Example request*

```json
{
  "amount": int,
  "payment_method": {
    "credit_card": {
      "card_number": string,
      "holder_name": string,
      "exp_date": string,
      "cvc": string
    }
  },
  "billing_address": {
    "first_name": string,
    "last_name": string,
    "postal_code": string,
    "city": string,
    "address_line_1": string,
    "email": string,
    "phone": string
  }
}
```
</details>

<details>
<summary> ACH Transaction </summary>

*Create an ach transaction*

> Request Method: `POST` <br>
> Request URL: `/api/transaction/ach`

| Plugin          | Type     | Description           | Required |
| --------------- | -------- | --------------------- | -------- |
| amount          | `uint64` | Amount                | `true`   |
| payment_method  | `object` | Method of the payment | `true`   |
| billing_address | `object` | Address of costumer   | `true`   |

*Example request*

```json
{
  "amount": int,
  "payment_method": {
    "ach": {
      "account_type": string,
      "account_number": string,
      "routing_number": string,
      "sec_code": string
    }
  },
  "billing_address": {
    "first_name": string,
    "last_name": string,
    "postal_code": string,
    "city": string,
    "address_line_1": string,
    "email": string,
    "phone": string
  }
}
```
</details>

<details>
<summary> Apple Pay Transaction </summary>

*Create an applepay  transaction*

> Request Method: `POST` <br>
> Request URL: `/api/transaction/applepay`

| Plugin          | Type     | Description           | Required |
| --------------- | -------- | --------------------- | -------- |
| amount          | `uint64` | Amount                | `true`   |
| payment_method  | `object` | Method of the payment | `true`   |
| billing_address | `object` | Address of costumer   | `true`   |

*Example request*

```json
{
  "amount": int,
  "payment_method": {
    "apple_pay": {
      "payment_token": {
        "identifier": string,
        "payment_data": string
      }
    }
  },
  "billing_address": {
    "first_name": string,
    "last_name": string,
    "postal_code": string,
    "city": string,
    "address_line_1": string,
    "email": string,
    "phone": string
  }
}
```
</details>

<details>
<summary> Google Pay Transaction </summary>

*Create an googlepay  transaction*

> Request Method: `POST` <br>
> Request URL: `/api/transaction/googlepay`

| Plugin          | Type     | Description           | Required |
| --------------- | -------- | --------------------- | -------- |
| amount          | `uint64` | Amount                | `true`   |
| payment_method  | `object` | Method of the payment | `true`   |
| billing_address | `object` | Address of costumer   | `true`   |

*Example request*

```json
{
  "amount": int,
  "payment_method": {
    "google_pay": {
      "encrypted_payment": {
        "payment_id": string,
        "payment_data": string
      }
    }
  },
  "billing_address": {
    "first_name": string,
    "last_name": string,
    "postal_code": string,
    "city": string,
    "address_line_1": string,
    "email": string,
    "phone": string
  }
}
```
</details>

<hr />

### Status codes

| Message              | Code  |
| -------------------- | ----- |
| Success              | `100` |
| Limit exceeded       | `201` |
| Card blocked         | `202` |
| Daily limit exceeded | `203` |
| Fraud detected       | `204` |
| Error occured        | `206` |

<hr />

### Special Cards

| Error              | Card Number  |
| -------------------- | ----- |
| Limit exceeded       | `4455444455551111` |
| Card blocked         | `0000000000000000` |
| Daily limit exceeded | `7755444455551111` |
| Fraud detected       | `8888888888888888` |

<hr />

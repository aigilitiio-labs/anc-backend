# ANC OMS Input API (anc-oms-input-api)

The `anc-oms-input-api` application handles order requests, stores the data in the `orders` collection, and exposes two endpoints: `CreateOrder` and `CancelOrder`.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Endpoints](#endpoints)
- [Docker](#docker)
- [License](#license)

## Prerequisites

- Go 1.20 or later
- MongoDB
- Docker (for containerization)

## Installation

1. Clone the repository:

   ```sh
   git clone    https://github.com/aigilitiio-labs/anc-backend/anc-oms-input-api.git
   cd anchorage-anc-input-api
   ```

2. Download the dependencies:

   ```sh
   go mod download
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

### Usage

Make sure you have MongoDB running and the environment variables set up. Then, you can start the server and interact with it through the exposed endpoints.

### Environment Variables

Create a .env file in the root directory with the following variables:

```
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=orders
BIND_ADDRESS=:9090
```

### Endpoints

#### Create Order

- **URL**: /orders
- **Method**: POST
- **Request Body**:

```json
{
	"client_order_id": "93de1f59-9827-4ee2-949f-3fdfd8d9b88b",
	"product_id": "BTC-USD",
	"side": "SELL",
	"order_configuration": {
		"market_market_ioc": {
			"base_size": "100"
		}
	}
}
```

- **Response**:

```json
{
	"message": "Order created successfully",
	"order": {
		"client_order_id": "93de1f59-9827-4ee2-949f-3fdfd8d9b88b",
		"product_id": "BTC-USD",
		"side": "SELL",
		"order_configuration": {
			"market_market_ioc": {
				"base_size": "100"
			}
		},
		"status": "New"
	}
}
```

#### Cancel Order

- **URL**: /orders
- **Method**: DELETE
- **Request Body**:

```json
{
	"client_order_id": "93de1f59-9827-4ee2-949f-3fdfd8d9b88b"
}
```

- **Response**:

```json
{
	"message": "Order cancelled successfully",
	"order": {
		"client_order_id": "93de1f59-9827-4ee2-949f-3fdfd8d9b88b",
		"status": "Cancelled"
	}
}
```

### Docker

1. #### Build and Run
   ```sh
   docker build -t anc-oms-input-api .
   ```
2. #### Run the Docker container:
   ```sh
   docker run -d --name anc-oms-input-api -p 9090:9090 --env-file .env anc-oms-input-api
   ```

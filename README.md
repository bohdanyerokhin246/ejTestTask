
# ejTestTask runing
To run project, run the following command

**Clone project**
```shell
  git clone https://github.com/bohdanyerokhin246/ejTestTask.git
```
**Go to project dir**
```shell
  cd ejTestTask
```
**Run project**
```shell
   docker-compose up --build   
```

# API Documentation

## Overview
This API provides endpoints for managing buyer, seller, product and order.<br> All requests require an **Authorization** header with a valid Bearer token.

# **Login**

## **POST /login**
- **Description**: Retrieves token by user info
- **Request Body**:
    ```json
    {
      "login": "admin",
      "password": "admin"
    }
    ```
- **Response**:
    ```json
    {
        "Token": "eyJhbGciOiJIUzergregInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzM4NDIzMDgyfQ.8189Qhn7LgIzE_smmTgiMIE6oF_7KMJJYZX6xaLSUeY"
    }
    ```


# **Buyer**

## **POST /buyer**
- **Description**: Creates a new buyer
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "name": "Heorhiy",
      "phone": "380950000000"
    }
    ```
- **Response**:
    ```json
    {
        "id": 4
    }
    ```
---
## **GET /buyers**
- **Description**: Retrieves buyers
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Response**:
    ```json
    [
      {
        "ID": 1,
        "name": "Bohdan",
        "phone": "380990000000"
      },
      {
        "ID": 2,
        "name": "Ihor",
        "phone": "380970000000"
      }
    ]
    ```
---
## **GET /buyer**
- **Description**: Retrieves buyer by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
  ```json
    {
      "ID":2
    }
  ```
- **Response**:
    ```json
    {
      "ID": 2,
      "name": "Ihor",
      "phone": "380970000000"
    }
    ```
---
## **PUT /buyer**
- **Description**: Updates buyer data by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 4,
      "name": "Ihnat",
      "phone": "380930000000"
    }
    ```
- **Response**:
    ```json
    {
      "message": "Buyer 4 updated"
    }
    ```
---
## **DELETE /buyer**
- **Description**: Deletes a buyer by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID":4
    }
    ```
- **Response**:
    ```json
    {
      "message": "Buyer 4 deleted"
    }
    ```
---
# **Seller**

## **POST /seller**
- **Description**: Creates a new seller
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "name": "Anna",
      "phone": "380950000001"
    }
    ```
- **Response**:
    ```json
    {
        "id": 5
    }
    ```

---

## **GET /sellers**
- **Description**: Retrieves all sellers
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Response**:
    ```json
    [
      {
        "ID": 1,
        "name": "Oleh",
        "phone": "380990000001"
      },
      {
        "ID": 2,
        "name": "Maksym",
        "phone": "380970000001"
      }
    ]
    ```

---

## **GET /seller**
- **Description**: Retrieves seller by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 2
    }
    ```
- **Response**:
    ```json
    {
      "ID": 2,
      "name": "Maksym",
      "phone": "380970000001"
    }
    ```

---

## **PUT /seller**
- **Description**: Updates seller data by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 5,
      "name": "Sergiy",
      "phone": "380930000001"
    }
    ```
- **Response**:
    ```json
    {
      "message": "Seller 5 updated"
    }
    ```

---

## **DELETE /seller**
- **Description**: Deletes a seller by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 5
    }
    ```
- **Response**:
    ```json
    {
      "message": "Seller 5 deleted"
    }
    ```

---
# **Product**
## **POST /product**
- **Description**: Creates a new product
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
        "name": "Iphone",
        "description": "This is Apple smartphone",
        "price": 350,
        "quantity": 20,
        "sellerID": 2
    }
    ```
- **Response**:
    ```json
    {
        "id": 5
    }
    ```
---

## **GET /products**
- **Description**: Retrieves all products
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Response**:
    ```json
    [
      {
        "ID": 2,
        "name": "Iphone",
        "description": "This is Apple smartphone",
        "price": 350,
        "quantity": 20,
        "sellerID": 2
      },
      {
        "ID": 3,
        "name": "Laptop",
        "description": "This is new laptop",
        "price": 500,
        "quantity": 10,
        "sellerID": 2
      }
    ]
    ```

---

## **GET /product**
- **Description**: Retrieves product by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 5
    }
    ```
- **Response**:
    ```json
    {
      "ID": 3,
      "name": "Laptop",
      "description": "This is new laptop",
      "price": 500,
      "quantity": 10,
      "sellerID": 2
    }
    ```
---

## **PUT /product**
- **Description**: Updates product data by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 3,
      "name": "Laptop 2",
      "description": "This is new laptop",
      "price": 500,
      "quantity": 10,
      "sellerID": 2
    }
    ```
- **Response**:
    ```json
    {
      "message": "Product 5 updated"
    }
    ```

---
## **DELETE /product**
- **Description**: Deletes product by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 5
    }
    ```
- **Response**:
    ```json
    {
      "message": "Product 5 deleted"
    }
    ```
---

---
# **Order**
## **POST /order**
- **Description**: Creates a new order
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "buyerID": 1,
      "products":  
       [
          {"ID": 1, "quantity": 5},
          {"ID": 4, "quantity": 1}
       ]
    }
    ```
- **Response**:
    ```json
    {
        "id": 5
    }
    ```
---

## **GET /orders**
- **Description**: Retrieves all orders
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Response**:
    ```json
    [
      {
        "ID": 4,
        "buyerID": 1,
        "price": 10050,
        "products": [
            {
                "ID": 1,
                "name": "Apple",
                "description": "This is green apple",
                "price": 10,
                "quantity": 5
            },
            {
                "ID": 4,
                "name": "BMW",
                "description": "This is fast car",
                "price": 10000,
                "quantity": 1
            }
        ],
        "createdAt": "2025-01-31T15:18:14.240965Z"
      },
      {
        "ID": 5,
        "buyerID": 1,
        "price": 10050,
        "products": [
            {
                "ID": 1,
                "name": "Apple",
                "description": "This is green apple",
                "price": 10,
                "quantity": 5
            },
            {
                "ID": 4,
                "name": "BMW",
                "description": "This is fast car",
                "price": 10000,
                "quantity": 1
            }
        ],
        "createdAt": "2025-01-31T15:18:14.240965Z"
      }
    ]
    ```

---

## **GET /order**
- **Description**: Retrieves order by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 5
    }
    ```
- **Response**:
    ```json
    {
      "ID": 5,
      "buyerID": 1,
      "price": 10050,
      "products":  
      [
        {
            "ID": 1,
            "name": "Apple",
            "description": "This is green apple",
            "price": 10,
            "quantity": 5
        },
        {
            "ID": 4,
            "name": "BMW",
            "description": "This is fast car",
            "price": 10000,
            "quantity": 1
        }
      ],
      "createdAt": "2025-01-31T15:18:14.240965Z"
    }
    ```

---

## **PUT /order**
- **Description**: Updates order data by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 1,
      "buyerID": 1,
      "products": 
      [
        {
            "ID": 12,
            "quantity": 10
        },
        {
            "ID": 41,
            "quantity": 12
        }
      ]
    }
    ```
- **Response**:
    ```json
    {
      "message": "Order 5 updated"
    }
    ```

---
## **DELETE /order**
- **Description**: Deletes order by ID
- **Request Header**:
    ```
    Authorization: Bearer {token}
    ```
- **Request Body**:
    ```json
    {
      "ID": 5
    }
    ```
- **Response**:
    ```json
    {
      "message": "Order 5 deleted"
    }
    ```
---


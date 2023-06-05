## GraphQL API Example

- All the query have pagination, you can use `page` and `limit` as input to get data

- Use `login` first for authentication

```graphql
mutation login {
    User {
        login(input: {
            email: "admin@techvify.com"
            password: "admin@123"
        }) {
            userId
            token
        }
    }
}
```

### Create User
  Create user specify `accessID` to `2` for User or `3` for Admin

- Create User need authentication or API will return `access denied`

- Add below section in header

```json
{
  "Authorization": "token_admin_here"
}
```

```graphql
mutation createUser {
    User {
        createUser(input:{
            email:"john.techvify@gmail.com"
            password: "123@abc123" # min 8 characters
            accessID: 3 # 2: User, 3: Admin
        })
    }
}
```

### Register Customer

When customer want to register, they need to provide `email`, `password`, `customerName`, `address`, `phoneNumber`, `identifyNumber`, `dateOfBirth`

```graphql
mutation registerCustomer {
    User {
        registerCustomer(input: {
            email: "john.techvify@gmail.com",
            password: "123@abc123",  # min 8 characters
            customerName: "John Tran",
            address:"Ha Noi",
            phoneNumber: "0195364958", # min 10 characters
            identifyNumber: "7vc12531",
            dateOfBirth: "2000-01-02T15:04:05Z"
        }) {
            data {
                id
                customerID
            }
            token
        }
    }
}
```

### Query User

Query User using for Admin

- Query User need authentication or API will return `access denied`
- Add below section in header

```json
{
  "Authorization": "token_admin_here"
}
```

```graphql
query queryUser{
    User {
        users(page: 1, limit: 5) {
            data{
                id
                email
                password
                customer{
                    id
                    name
                    email
                    address
                    phoneNumber
                    identifyNumber
                    dateOfBirth
                    memberCode
                    createdAt
                    updatedAt
                }
                accessType {
                    id
                    name
                    createdAt
                    updatedAt
                }
                createdAt
                updatedAt
            }
            page
            limit
            total
        }
    }
}
```

### Login

Login for everyone

```graphql
mutation login {
    User {
        login(input: {
            email: "john.techvify@gmail.com"
            password: "123@abc123"
        }) {
            userId
            token
        }
    }
}
```

### Customer Change Password
For customer change password, they need to provide `oldPassword` and `newPassword`

- changePassword need authentication or API will return `access denied`

- Add below section in header

```json
{
  "Authorization": "token_user_here"
}
```

```graphql
mutation changePassword {
    Customer {
        changePassword(input: {
            oldPassword: "123123123",
            newPassword: "321321321"
        })
    }
}
```

### Update Customer

For Customer Update their information, they need to provide `email`, `name`, `address`, `phoneNumber`, `identifyNumber`, `dateOfBirth`

- update customer need authentication or API will return `access denied`

- Add below section in header

```json
{
  "Authorization": "token_user_here"
}
```

```graphql
mutation updateCustomer {
    Customer {
        updateCustomer(input: {
            email: "john.techvify@gmail.com",
            name: "any",
            address: "any",
            phoneNumber: "min 10 characters"
            identifyNumber: "123123123", # max 12 characters
            dateOfBirth: "2002-01-02T15:04:05Z",
        })
    }
}
```

### Query Customer

Using for Admin

- queryCustomer need authentication or API will return `access denied`
- Add below section in header

```json
{
  "Authorization": "token_admin_here"
}
```

```graphql
query queryCustomer {
    Customer {
        customers(page: 1, limit: 5) {
            data {
                id
                name
                email
                address
                phoneNumber
                identifyNumber
                dateOfBirth
                memberCode
                createdAt
                updatedAt
            }
            page
            limit
            total
        }
    }
}
```

### Create FLight

For Admin create flight, they need to provide `name`, `from`, `to`, `departureDate`, `arrivalDate`, `available_first_slot`, `available_economy_slot`, `status`

- createFlight need authentication or API will return `access denied`
- Add below section in header

```json
{
  "Authorization": "token_admin_here"
}
```

```graphql
mutation createFlight {
    Flight {
        createFlight(input: {
            name: "112",
            from: "HN"
            to: "SGN"
            departureDate:"2023-05-20T13:50:20+07:00",
            arrivalDate: "2023-05-22T06:00:20.720717Z",
            available_first_slot: 30,
            available_economy_slot: 70,
            status: "Available", # Cancel, Available, Arrived
        })
    }
}

```

### Update Flight

For Admin update flight, they need to provide `name`, `departureDate`, `arrivalDate`, `available_first_slot`, `available_economy_slot`, `status`

- For status: `Cancel`, `Available`, `Arrived`

- If the flight was marked as `Cancel` or `Arrived` before, it can not be updated

- updateFlight need authentication or API will return `access denied`

- Add below section in header

```json
{
  "Authorization": "token_admin_here"
}
```

```graphql
mutation updateFlight {
    Flight {
        updateFlight(id: 442944984, input: {
            name:"113"
            departureDate: "2023-06-08T03:31:20.720717Z"
            arrivalDate: "2023-06-08T03:31:20.720717Z"
            available_first_slot: 30,
            available_economy_slot: 70,
            status: "Cancel" # Cancel, Available, Arrived
        })
    }
}
```

### Cancel Booking

For Customer cancel booking, they need to provide `bookingCode`

For Guest booking, it will return `call airline for cancel booking`

- cancelBooking need authentication or API will return `access denied`

- Add below section in header

```json
{
  "Authorization": "token_user_here"
}
```

```graphql
mutation cancelBooking {
    Booking {
        cancelBooking(input: {bookingCode: "4283XS"}) {
            code
            data {
                bookingCode
                cancelDate
                customerId
                flightId
                status
            }
        }
    }
}
```

### Create Booking

- If guest booking with no login, they need to provide `flight name`, `email`, `phone number`, `name`, `address`, `date of birth`, `identify number`, `ticket type`

```graphql
mutation createBooking {
    Booking {
        createBooking(input:{
            flightName: "1345"
            email: "john.techvify@gmail.com"
            phoneNumber: "012345677"
            name: "nvl"
            address: "Ha Noi"
            dateOfBirth: "2023-04-08T03:31:20+07:00"
            identifyNumber: "u2c22c2ccc"
            ticketType: 1
        }) {
            code
            data {
                bookingCode
                bookingDate
                customerId
                flightId
                status
            }
        }
    }
}
```

- If user booking, only `flight name` and `ticket type` needed

```json
{
  "Authorization": "token_user_here"
}
```

```graphql
mutation createBooking {
    Booking {
        createBooking(input:{
            flightName: "1345"
            #    email: "john.techvify@gmail.com"
            #    phoneNumber: "012345677"
            #    name: "nvl"
            #    address: "Ha Noi"
            #    dateOfBirth: "2023-04-08T03:31:20+07:00"
            #    identifyNumber: "u2c22c2ccc"
            ticketType: 1
        }) {
            code
            data {
                bookingCode
                bookingDate
                customerId
                flightId
                status
            }
        }
    }
}
```

### View Booking

- viewBooking show booking which booked by guest

- `booking code` and `identify number` is required

- if identify number not match with booking, it will return `identify number not match`

```graphql
query viewBooking {
    Booking {
        viewBooking(bookingCode: "R87Q26", identifyNumber: "13512331"){
            code
            data {
                bookingCode
                bookingDate
                bookingStatus
                customer {
                    id
                    name
                    email
                    address
                    phoneNumber
                    memberCode
                }
                flight {
                    id
                    name
                    from
                    to
                    departureDate
                    arrivalDate
                }
            }
        }
    }

}
```

### Booking History

- bookingHistory show booking history which booked by user

- bookingHistory need authentication or API will return `access denied`

```json
{
  "Authorization": "token_user_here"
}
```

```graphql
query bookingHistory {
    Booking {
        bookingHistory(page: 1, limit: 5) {
            code
            data {
                bookingCode
                bookingDate
                bookingStatus
                customer {
                    id
                    name
                    email
                    address
                    phoneNumber
                    memberCode
                }
                flight {
                    id
                    name
                    from
                    to
                    departureDate
                    arrivalDate
                    status
                }
            }
            page
            limit
            total
        }
    }
}
```

### Search Flight
For everyone search flight, they need to provide `from`, `to`, `departureDate`, `ticketType`

- searchFlight show flight which match with `from`, `to`, and between `departureDate`, `ticketType`

```graphql
query searchFlight {
    Flight {
        searchFlights(
            page: 1
            limit: 5
            input: {
                from: "HN"
                to: "SGN"
                arrivalDate: "2023-05-20T13:50:20+07:00"
                departureDate: "2023-05-22T13:50:20+07:00"
            }) {
            data {
                id
                name
                from
                to
                departureDate
                arrivalDate
                available_first_slot
                available_economy_slot
                status
            }
            page
            limit
            total
        }
    }
}
```

### Query Flight
For admin query flight

- queryFlight show all flight

- queryFlight need authentication or API will return `access denied`

```graphql
query queryFlights {
    Flight {
        flights(page: 1, limit: 20) {
            data {
                id
                name
                from
                to
                departureDate
                arrivalDate
                available_first_slot
                available_economy_slot
                status
                createdAt
                updatedAt
            }
            page
            limit
            total
        }
    }
}
```

### Update User Password
For Admin change admin account password

- updatePassword need authentication of admin or API will return `access denied`

```json
{
  "Authorization": "token_admin_here"
}
```

```graphql
mutation updateUserPassword {
    User {
        updateUserPassword(input: {
            oldPassword: "123123123"
            password: "321321321"
        })
    }
}
```
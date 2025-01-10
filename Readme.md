# Employee REST API Documentation

## Description
A simple employee management system built using Go and MySQL as the database. This API allows you to manage employee data, including adding, updating, deleting, and retrieving employee information.

## System Requirements
- **Go** >= 1.22.6
- **MariaDB**

## How to Run the Application

1. **Clone the repository:**
    ```bash
    $ git clone https://github.com/boombaw/go-employee.git
    $ cd go-employee
    ```

2. **Set up database:**
    - Create database with database name **db_employee**
    - Restore **db_employee.sql** on folder **pkg/v1/mock-data**.
       
3. **Set up the environment file:**
    - Create a `.env` file in the root directory.
    - Copy the content from `.env.example` and adjust the values as needed.

4. **Run the application:**
    ```bash
    $ go run main.go
    ```
    Or use the Makefile if available:
    ```bash
    $ make run
    ```


## Base URL
`http://localhost:3000`

## Endpoints

### Employee

1. **Get All Employees**
   - **URL**: `/api/v1/employees`
   - **Method**: GET
   - **Headers**:
     - `x-api-key`: `XXZiqjZgoLXTO23VPJEjxtc7ZXX`
   - **Description**: Retrieve all employee data.

2. **Get Employee by ID**
   - **URL**: `/api/v1/employee/{id}`
   - **Method**: GET
   - **Headers**:
     - `x-api-key`: `XXZiqjZgoLXTO23VPJEjxtc7ZXX`
   - **Description**: Retrieve employee data by ID.

3. **Get Employees by Status**
   - **URL**: `/api/v1/employee?status=active`
   - **Method**: GET
   - **Headers**:
     - `x-api-key`: `XXZiqjZgoLXTO23VPJEjxtc7ZXX`
   - **Description**: Retrieve employee data with active status.

4. **Create Employee**
   - **URL**: `/api/v1/employee`
   - **Method**: POST
   - **Headers**:
     - `Content-Type`: `application/json`
   - **Body**:
     ```json
     {
       "name": "adi",
       "department": "IT",
       "status": "active",
       "salary": 15000
     }
     ```
   - **Description**: Add a new employee.

5. **Update Employee**
   - **URL**: `/api/v1/employee/{id}`
   - **Method**: PUT
   - **Headers**:
     - `Content-Type`: `application/json`
   - **Body**:
     ```json
     {
       "department": "IT",
       "status": "active",
       "salary": 15000
     }
     ```
   - **Description**: Update employee data by ID.

### Department

1. **Get Average Salary by Department**
   - **URL**: `/api/v1/department?name=IT`
   - **Method**: GET
   - **Headers**:
     - `x-api-key`: `XXZiqjZgoLXTO23VPJEjxtc7ZXX`
   - **Description**: Retrieve the average salary by department.


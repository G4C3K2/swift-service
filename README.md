# Swift Service - Recruitment Task

## Requirements

* **Docker:** Docker must be installed and running.
* **Docker Compose:** Docker Compose must be installed.
* **Go:** A working Go environment is required.
* **Gin Web Framework:** The application utilizes the Gin web framework for routing and handling HTTP requests.
* **MongoDB Driver:** The application interacts with MongoDB using the official Go driver.

## Running the Application

1.  **Clone the Repository:**
    ```bash
    git clone <repository_address>
    cd <repository_name>
    ```

2.  **Database Configuration:**
    The application uses the `MONGO_URI` environment variable to connect to the MongoDB database. Before running the application, you need to set this variable in file docker-compose.yaml to point to a running MongoDB instance.

    In docker-compose.yaml you should see example `MONGO_URI` format:
    ```
    MONGO_URI=mongodb+srv://<user>:<password>.1@cluster0.wpmnnvs.mongodb.net/
    ```

    You should be provided with connection strin after creating new cluster on MongoDB.

3.  **Run the Application using Docker Compose:**
    ```bash
    docker-compose up --build -d
    ```
    This command will build (if necessary) and run the containers defined in `docker-compose.yaml` in detached mode (in the background).

## Accessing the Application

Once the containers are running, the application will be accessible at:

http://localhost:8080

After successfuly building project you should have file Data.csv in your root folder. It is swift spreadsheet provided for recrutation.

To load data from the `Data.csv` file (located in the root directory of the project) into the database, send a `POST` request to the endpoint: /swift-codes/load

To retrieve the details of a bank based on its SWIFT code, send a GET request to the endpoint: /swift-codes/{swiftCode} 
Replace {swiftCode} with the specific SWIFT code (e.g., INGBPW5SXXX).

To retrieve a list of banks for a given country ISO 2 code, send a GET request to the endpoint: /countries/{countryISO2code}
Replace {countryISO2code} with the specific ISO 2 code (e.g., PL).

To add a new SWIFT code, send a POST request to the /swift-codes endpoint with a JSON body in the models.AddSwiftCodeRequest format:

{
  "SwiftCode": "XXXXYYZZ123",
  "BankName": "New Bank Name",
  "Address": "Bank Address",
  "CountryISO2": "XY",
  "CountryName": "Country Name",
  "IsHeadquarter": false
}

To delete an existing SWIFT code, send a DELETE request to the endpoint: /swift-codes/{swiftCode}
Replace {swiftCode} with the SWIFT code to delete (e.g., INGBPW5SXXX).

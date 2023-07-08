# IPFS Cluster File Upload and Download

This is a simple web application that allows users to upload files, store them in an IPFS cluster, and download files from the cluster.

## Prerequisites

- Go programming language (v1.16 or higher)
- IPFS (go-ipfs)
- IPFS Cluster
- PostgreSQL (optional if you want to store metadata)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/rishitashaw/ipfs-cluster.git
   cd ipfs-cluster
   ```

2. Install the dependencies:

   ```shell
   go mod download
   ```

3. Start the IPFS daemon:

   ```shell
   docker-compose up -d
   ```

4. Build and run the application:

   ```shell
   go run main.go
   ```

   The application will be available at `http://localhost:8888`.

## Usage

- Open your web browser and visit `http://localhost:8888`.
- Use the file upload form to select a file and upload it to the IPFS Cluster.
- Once the upload is complete, the file will be stored in the IPFS Cluster, and you will see its hash displayed on the page.
- You can then click the "Download" button to retrieve the file from the IPFS Cluster.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

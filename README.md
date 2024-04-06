# Distributed File System

Distributed File System Project

How To Run in Local Network

```
Add IP for Master in the Config file

// Run Master
go run .\master_tracker\main.go

// Run Data Node
go run .\data_keeper\main.go

// Run Client
go run ./client/main.go
```

#### Run Locally

1. **Update Config File**:
   Update the configuration file and add local host as ip for the master

2. **Run Master Node**:

   ```sh
   go run ./master_tracker/main.go

   ```

3. **Run Data Node**:

   ```sh
   go run .\data_keeper\main.go 127.0.0.1

   ```

4. **Run Client**:
   ```sh
   go run ./client/main.go 127.0.0.1
   ```

#### Run In Network

1. **Update Config File**:
   Update the configuration file to include the IP address of the Master node.

2. **Run Master Node**:

   ```sh
   go run ./master_tracker/main.go

   ```

3. **Run Data Node**:

   ```sh
   go run .\data_keeper\main.go

   ```

4. **Run Client**:
   ```sh
   go run ./client/main.go
   ```

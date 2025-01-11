# Distributed File System 🗃️🖥️🌐

<p align="center">
  <img width="100%" src="https://png.pngtree.com/thumb_back/fh260/background/20230616/pngtree-isometric-illustration-3d-rendering-of-non-fungible-tokens-nfts-image_3623483.jpg" alt="Distributed system Image" />
</p>


## <img  align= center width=80px src="giphy.gif">  Table of Content
<!--  Overview  -->
## <img  align= center width =60px src="https://cdn-icons-png.flaticon.com/512/8632/8632710.png"> Overview <a id="overview"></a>
With the ever-growing technological expansion, distributed systems are becoming increasingly prevalent in the modern world. These systems represent a critical and complex area of study in computer science. At their core, a distributed system is a group of computers that collaborate to appear as a single machine to the end user.

Key characteristics of distributed systems include:
- **Shared State**: The machines share data and maintain synchronization.
- **Concurrent Operations**: They operate simultaneously, working together efficiently.
- **Fault Tolerance**: Components can fail independently without disrupting the overall system’s uptime.
  
In this project, we aim to build a simple distributed file system capable of:
- Reading and writing .mp4 files for user interaction.
- File replication to ensure fault tolerance and reliability across the system.
This project offers a hands-on experience with the fundamentals of distributed systems while emphasizing scalability and resilience.

<!-- Built Using -->
## <img  align= center width =60px  height =70px src="https://media4.giphy.com/media/ux6vPam8BubuCxbW20/giphy.gif?cid=6c09b952gi267xsujaqufpqwuzeqhbi88q0ohj83jwv6dpls&ep=v1_stickers_related&rid=giphy.gif&ct=s"> Built Using <a id="tools"></a>
<table>
  <tr>
        <td align="center"><img height="100" src ="https://img.icons8.com/?size=512&id=44442&format=png"/></td>
        <td align="center"><img height="100" src ="https://miro.medium.com/v2/resize:fit:1400/1*xZXmBNa-o0P5YYsKmsKO0Q.png"/></td>
  </tr>
</table>

<!-- Getting Started -->
## <img align="center" width="60px" height="60px" src="https://media3.giphy.com/media/wuZWV7keWqi2jJGzdB/giphy.gif?cid=6c09b952wp4ev7jtywg3j6tt7ec7vr3piiwql2vhrlsgydyz&ep=v1_internal_gif_by_id&rid=giphy.gif&ct=s"> Getting Started <a id="started"></a>

#### Clone the Repository
   ```sh
    git clone https://github.com/BasmaElhoseny01/Distributed-File-System.git
   ```
    
#### Run Locally
1. **Update Config File**:
   Update the configuration file and add local host(127.0.0.1) as IP for the master

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


<!-- Contributors -->
## <img  align= center width=50px height=50px src="https://media1.giphy.com/media/WFZvB7VIXBgiz3oDXE/giphy.gif?cid=6c09b952tmewuarqtlyfot8t8i0kh6ov6vrypnwdrihlsshb&rid=giphy.gif&ct=s"> Contributors <a id = "contributors"></a>

<!-- Contributors list -->
<table align="center" >
  <tr>
    <td align="center"><a href="https://github.com/Ahmedsabry11"><img src="https://avatars.githubusercontent.com/u/75908511?v=4" width="150px;" alt=""/><br /><sub><b>Ahmed Sabry</b></sub></a></td>
    <td align="center"><a href="https://github.com/AhmedHosny2024"><img src="https://avatars.githubusercontent.com/u/76389601?v=4" width="150px;" alt=""/><br /><sub><b>Ahmed Hosny</b></sub></a></td>
    <td align="center"><a href="https://github.com/zeinabmoawad"><img src="https://avatars.githubusercontent.com/u/92188433?v=4" width="150px;" alt=""/><br /><sub><b>Zeinab Moawed</b></sub></a></td>
    <td align="center"><a href="https://github.com/BasmaElhoseny01"><img src="https://avatars.githubusercontent.com/u/72309546?v=4" width="150px;" alt=""/><br /><sub><b>Basma Elhoseny</b></sub></a></td>
  </tr>
</table>

## <img  align= center height=50px src="https://alhidaayahacademy.sch.ng/wp-content/uploads/2016/04/LE_certificate_gr.gif"> License <a id="license"></a>
This software is licensed under the [MIT License](https://github.com/BasmaElhoseny01/Distributed-File-System/blob/main/LICENSE). © Basma Elhoseny.

# jianji-server

## 配置数据库

### 使用docker

1. 新建容器
    ```bash
    docker run --name my-postgres -e POSTGRES_PASSWORD=333333 -d postgres  
    ```
   
2. docker exec
    ```bash
    psql -U postgres
    ```
    
    ```postgresql
    CREATE DATABASE jianji;
    CREATE USER jianji WITH PASSWORD '333333';
    ```

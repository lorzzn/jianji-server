# jianji-server

一个使用gorm，postgres，gin搭建的笔记项目的后端，前端项目：[jianji-web](https://github.com/lorzzn/jianji-web)

## 配置数据库

1. 新建容器
    ```bash
    docker run --name postgres-for-jianji -p 5434:5432 -e POSTGRES_PASSWORD=333333 -d postgres  
    ```
   
2. docker exec
    ```bash
    psql -U postgres
    ```
    
    ```postgresql
   CREATE ROLE jianji WITH LOGIN PASSWORD '333333';
    CREATE DATABASE jianji OWNER jianji;
    GRANT ALL PRIVILEGES ON DATABASE jianji TO jianji;
    ```

## 配置Redis

1. 新建容器
   ```bash
   docker run --name redis-for-jianji -d -p 6379:6379 redis
   ```

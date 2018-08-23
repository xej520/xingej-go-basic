## 1、驱动包
 https://github.com/Go-SQL-Driver/MySQL
### 1.1、 下载驱动包
 go get github.com/Go-SQL-Driver/MySQL

### 1.2、创建表user_info
    create table user_info(id INT(11), username varchar(20), departname varchar(20), create_time TIMESTAMP default "0000-00-00 00:00:00", primary KEY(id))

## 2、插入数据
```
//Prepare是预编译
 stmt, err := db.Prepare("INSERT user_info  SET username=?,departname=?,create_time=?")

  res, err := stmt.Exec("zhangsan", "技术部", "2012-12-09")
  id, err := res.LastInsertId()
```
## 3、更新数据
```
stmt, err = db.Prepare("update user_info set username=? where uid=?")

res, err = stmt.Exec("lisi", id)
affect, err := res.RowsAffected()

```
## 4、查询数据
```
rows, err := db.Query("SELECT * FROM user_info")
for rows.Next() {
        var id int
        var username string
        var department string
        var create_time string
        err = rows.Scan(&uid, &username, &department, &create_time)
}

```
## 5、删除数据
```
stmt, err = db.Prepare("delete from user_info where uid=?")
res, err = stmt.Exec(id)

```


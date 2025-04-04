### 核心API
#### 用户登录
```
POST http://118.178.120.11:8888/v1/user/login
数据：
{
"phone":"17344995006",
"password":"admin"
}
//或
// {
//     "phone":"22233334444",
//     "password":"123456"
// }
```
注：返回 jwt token

### 获得好友列表
```
GET http://118.178.120.11:8888/v1/social/friends
无数据
```
注： 需要携带 token

### 获得群组列表
```
GET http://118.178.120.11:8888/v1/social/groups
{
    "user_id":"1843306319776321536"
}
```
注： 需要携带 token



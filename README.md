# InstagramAPIClone
Instagram Api Clone using Go.
The api has following routes<br>
POST ‘/users' - create new users<br>
GET ‘/users/<id here>’ - get users from their id<br>
POST ‘/posts' - create new post<br>
GET ‘/posts/<id here>’ - get post from post id<br>
GET ‘/posts/users/<Id here>' - get all posts by a user.The response is paginated<br>

Tech Stack:Go,MongoDB

The entities are present in models package and handler functions in controller package. main.go is the entry point.
Tests in all_test.go file using testing standard library.

Task given by Appointy.

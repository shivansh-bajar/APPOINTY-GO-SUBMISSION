ALL APIs ARE IMPLMENTED IN THE main.go FILE

TO CALL DIFFERENT APIs USE THE FOLLOWING URL AFTER RUNNING THE main.go FILE

For addUser:
localhost:8082/users

For addPosts:
localhost:8082/posts

getUser:
localhost:8082/users/?id=<object_id>
Here object_id is the ObjectID given by MongoDB

getPost:
localhost:8082/posts/?id=<object_id>
Here object_id is the ObjectID given by MongoDB

getPostList:
localhost:8082/posts/users/?id=<Name>
Here Name is the Name value of the object stored in MongoDB,
This Name value refers to the name of the user who made the post


TO RUN THE main.go FILE:
2 COMMANDS IN THE TERMINAL
go env -w GO111MODULE=off
go run main.go


THANK YOU!!!
// Install dependencies
go get -u github.com/gin-gonic/gin
go get -u github.com/globalsign/mgo
go get -u github.com/globalsign/mgo/bson
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go // Requred to generate proto.



// Generate proto file.
Install protoc - brew install protobuf
Run: protoc --go_out=. --go-grpc_out=. path/to/your/protofile.proto
     Replace path/to/your/protofile.proto with the path to your .proto file.


ERROR:
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.

To fix this error, make sure that you have installed the Go protocol buffer plugin using the following command:
go get -u github.com/golang/protobuf/protoc-gen-go

After installing the plugin, make sure that the bin directory of your GOPATH is included in the system's PATH variable. You can check this by running the echo $PATH command in your terminal and verifying that the directory containing the protoc-gen-go binary is included in the output.

If the bin directory is not included in the PATH variable, you can add it by running the following command:
export PATH=$PATH:$(go env GOPATH)/bin

If the protoc-gen-go binary is not present in your bin directory:
Make sure that your GOPATH and GOBIN environment variables are set correctly. You can check them by running the following commands:
echo $GOPATH
echo $GOBIN

If they are not set, you can set them using the following commands:
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin

Try reinstalling the Go protocol buffer plugin using the following command:
go install google.golang.org/protobuf/cmd/protoc-gen-go

After this you might also need this:
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc


Usaage:
protoc --go_out=. --go-grpc_out=. api/proto/contact.proto

.................

To import generated grpc classes:
import (
	pb "your-project-path/your-proto-package" // import the generated protobuf package
)
You should replace your-project-path and your-proto-package with the actual values for your project and protobuf package. Then, you can use the generated classes in your service implementation.







For error:
no required module provides package go.mongodb.org/mongo-driver/bson/primitive; to add it:
go get go.mongodb.org/mongo-driver/bson/primitive

And use import:
import "go.mongodb.org/mongo-driver/bson/primitive"



If you are getting error on import of these:
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

    do following:
        go get go.mongodb.org/mongo-driver/mongo@v1.11.4
        go get go.mongodb.org/mongo-driver/x/mongo/driver/auth@v1.11.4



Sample data:

{
    "first_name" : "Ranjit", 
    "last_name" : "Kaliraj", 
    "email" : "ranjit.kaliraj@rew3.com",
    "phone" : "+977 122 333 2323",
    "address" : {
        "street" : "Avenue Street",
        "city" : "Edmond",
        "state" : "OKC",
        "zip" : "900032"
    }
}



// Deploy to server. 
cd to cmd
go build cmd/main.go -o app-name

For cross architecture:
env GOOS=linux GOARCH=amd64 go build -o demo_go_crm_server
https://blog.devgenius.io/how-to-handle-cannot-execute-binary-file-exec-format-error-when-running-golang-executable-b5b77110b820



Enable Env read:
go mod init goenv
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
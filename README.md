![my image][def]

[def]: /:D/MYFILE/go-workspace/src/task/ERD-MyTaskAPP.jpg "ERD-MyTaskAPP"

documentation API: https://app.swaggerhub.com/apis-docs/farihatulilmiyya2625/My-Task-APP/1.0.0

mockery --dir=features/user --name=UserDataInterface --filename=UserData.go --structname=UserData
go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

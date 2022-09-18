gen-orm:
	go generate ./ent
gen-mocks:
	moq -out mocks_test.go -rm internal/module/todo/infrastructure/ TodoRespository

gen-api:
	buf generate
	buf build

gen-orm:
	go generate ./ent

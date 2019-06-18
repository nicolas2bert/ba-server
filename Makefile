generate: swagger.yaml
		# $< the first prerequisite
		# TODO: -P, --principal= the model to use for the security principal
		# TODO: add --with-flatten=full ?
	 	swagger generate server -f $< -A ba-api -t gen --principal=auth.PrincipalBA

build:
		GO111MODULE=on go build -v -o ba-server gen/cmd/ba-server/main.go

run:
		HOST=0.0.0.0 PORT=8383 ./ba-server

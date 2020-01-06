# These environment variables must be set for deploymeny to work.
S3_BUCKET := nickthecloudguy-email
STACK_NAME := nickthecloudguy-email

# Common values used throughout the Makefile, not intended to be configured.
TEMPLATE = template.yaml
PACKAGED_TEMPLATE = packaged.yaml

.PHONY: lambda
lambda:
	dep ensure -v
	GOOS=linux GOARCH=amd64 go build -o bin/nickthecloudguy-email ./nickthecloudguy-email

.PHONY: clean
clean: 
	rm -rf ./bin ./vendor Gopkg.lock
	
.PHONY: build
build: clean lambda

.PHONY: run
run: build
	sam local start-api

.PHONY: package
package: build
	sam package --template-file $(TEMPLATE) --s3-bucket $(S3_BUCKET) --output-template-file $(PACKAGED_TEMPLATE)

.PHONY: deploy
deploy: package
	sam deploy --stack-name $(STACK_NAME) --template-file $(PACKAGED_TEMPLATE) --capabilities CAPABILITY_IAM

.PHONY: teardown
teardown:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
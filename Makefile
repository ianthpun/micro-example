export GO111MODULE := on
export CGO_ENABLED := 0
export DOCKER_BUILDKIT := 1

.PHONY: containerize 
containerize:
	@docker build \
	--build-arg SERVICE_DIR=${SERVICE_DIR} \
       	-f ${SERVICE_DIR}/Dockerfile \
        -t ${IMAGE} \
	. 1>/dev/null

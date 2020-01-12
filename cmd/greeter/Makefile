MONOREPO_ROOT=`pwd`/../../
SERVICE_DIR=cmd/greeter

export GO111MODULE := on
export CGO_ENABLED := 0
export DOCKER_BUILDKIT := 1

.PHONY: containerize
containerize:
	@cd ${MONOREPO_ROOT};docker build \
       	-f ${SERVICE_DIR}/Dockerfile \
        -t ${IMAGE} \
	. 1>/dev/null

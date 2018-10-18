FROM golang:1.11-alpine

ADD . $GOPATH/src/github.com/EconomistDigitalSolutions/aws-secret-manager-utility

WORKDIR $GOPATH/src/github.com/EconomistDigitalSolutions/aws-secret-manager-utility

RUN apk add --no-cache git curl gcc musl-dev && \
    export GO111MODULE=on && \
    go get -u github.com/c4milo/github-release && \
    go get -u github.com/mitchellh/gox

CMD export GO111MODULE=on && \
    rm -rf build && \
    cd cmd/awssecrets && \
	$GOPATH/bin/gox -ldflags "-X main.Version=$VERSION" \
	-osarch="darwin/amd64" \
	-osarch="linux/amd64" \
	-osarch="windows/amd64" \
	-output "../../build/{{.Dir}}_$VERSION_{{.OS}}_{{.Arch}}/$NAME" && \
    cd ../.. && \
    rm -rf dist && \
    mkdir dist && \
    ls build && \ 
    files=$(ls build) && \
    for f in $files; do \
		(cd $PWD/build/$f && ls && pwd && tar -cvzf ../../dist/$f.tar.gz *); \
		echo $f; \
	done && \
    cd $GOPATH/src/github.com/EconomistDigitalSolutions/aws-secret-manager-utility && \
    latest_tag=$(git describe --tags `git rev-list --tags --max-count=1`) && \
    comparison="$latest_tag..HEAD" && \
    if [ -z "$latest_tag" ]; then comparison=""; fi && \
    changelog=$(git log $comparison --oneline --no-merges --reverse) && \
    $GOPATH/bin/github-release EconomistDigitalSolutions/$NAME $VERSION "$(git rev-parse --abbrev-ref HEAD)" "**Changelog**<br/>$changelog" 'dist/*'; 

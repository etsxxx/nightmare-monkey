NAME=nightmare-monkey
dockerbuild:
	docker build -t $(NAME) .

build:
	docker run --rm \
		-v "$(PWD)/..:/root/go/src/github.com/etsxxx/nightmare-monkey" \
		-v "$(PWD)/out/bin:/root/go/src/github.com/etsxxx/nightmare-monkey/bin" \
		-w /root/go/src/github.com/etsxxx/nightmare-monkey \
		-v "$(PWD)/out/rpms:/root/rpmbuild/RPMS" \
		$(NAME) \
		rpmbuild -bb rpmbuild/spec/nightmare-monkey.spec
clean:
	docker image rm -f $(NAME)

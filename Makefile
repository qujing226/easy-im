user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test

release-test: user-rpc-dev

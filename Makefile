.PHONY: bump
bump:
	@echo "🚀 Bumping Version"
	git tag $(shell svu patch)
	git push --tags

.PHONY: build
build:
	@echo "🚀 Building Version $(shell svu current)"
	go build -o bindiff ./cmd/bindiff/main.go
VERSION_FILE=version.txt
MAJOR=$(shell cat $(VERSION_FILE) | cut -d'.' -f1)
MINOR=$(shell cat $(VERSION_FILE) | cut -d'.' -f2)
PATCH=$(shell cat $(VERSION_FILE) | cut -d'.' -f3)
OLD_VERSION=$(MAJOR).$(MINOR).$(PATCH)

define change-version =
	@echo Old Version: $(OLD_VERSION)
	@echo New Version: $(NEW_VERSION)
	@echo $(NEW_VERSION) > $(VERSION_FILE)
	
	sed -i -e 's/${OLD_VERSION}/${NEW_VERSION}/g' README.md
endef

dependencies:
	dep ensure

test:
	go test -v ./...

install:
	go clean
	make dependencies
	make test
	go install

releasePatch: NEW_VERSION=$(shell echo $(MAJOR).$(MINOR).$$(($(PATCH) + 1)))
releasePatch:
	$(change-version)

releaseMinor: NEW_VERSION=$(shell echo $(MAJOR).$$(($(MINOR) +1)).$(PATCH))
releaseMinor:
	$(change-version)

releaseMajor: NEW_VERSION=$(shell echo $$(($(MAJOR) +1)).$(MINOR).$(PATCH))
releaseMajor:
	$(change-version)


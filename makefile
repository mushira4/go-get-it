VERSION_FILE=version.txt
MAJOR=$(shell cat $(VERSION_FILE) | cut -d'.' -f1)
MINOR=$(shell cat $(VERSION_FILE) | cut -d'.' -f2)
PATCH=$(shell cat $(VERSION_FILE) | cut -d'.' -f3)

OLD_VERSION=$(MAJOR).$(MINOR).$(PATCH)
NEW_PATCH_VERSION=$(shell echo $(MAJOR).$(MINOR).$$(($(PATCH) + 1)))
NEW_MINOR_VERSION=$(shell echo $(MAJOR).$$(($(MINOR) +1)).$(PATCH))
NEW_MAJOR_VERSION=$(shell echo $$(($(MAJOR) +1)).$(MINOR).$(PATCH))


dependencies:
	dep ensure

test:
	go test -v ./...

install:
	go clean
	make dependencies
	make test
	go install

releasePatch:
	@echo Old Version: $(OLD_VERSION)
	@echo New Version: $(NEW_PATCH_VERSION)
	@echo $(NEW_PATCH_VERSION) > $(VERSION_FILE)
	
	sed -i -e 's/${OLD_VERSION}/${NEW_PATCH_VERSION}/g' README.md


releaseMinor:
	@echo Old Version: $(OLD_VERSION)
	@echo New Version: $(NEW_MINOR_VERSION)
	@echo $(NEW_MINOR_VERSION) > $(VERSION_FILE)
	
	sed -i -e 's/${OLD_VERSION}/${NEW_MINOR_VERSION}/g' README.md
	
releaseMajor:
	@echo Old Version: $(OLD_VERSION)
	@echo New Version: $(NEW_MAJOR_VERSION)
	@echo $(NEW_MAJOR_VERSION) > $(VERSION_FILE)
	
	sed -i -e 's/${OLD_VERSION}/${NEW_MAJOR_VERSION}/g' README.md


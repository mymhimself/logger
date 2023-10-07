
VERSION ?= $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT ?= $(shell git log -1 --format='%H')
BRANCH ?= $(shell git for-each-ref --format='%(objectname) %(refname:short)' refs/heads | awk "/^$$(git rev-parse HEAD)/ {print \$$2}")

#
# ─── GLOBAL VARIABLES ───────────────────────────────────────────────────────────
#

REPORTS_PATH ?= ./reports
COVERAGE_REPORT ?= $(REPORTS_PATH)/coverage.out

PROJECT_KEY ?= krm_venture_logger

#
# ────────────────────────────────────────────────  ──────────
#   :::::: H E L P : :  :   :    :     :        :          :
# ──────────────────────────────────────────────────────────
#
.PHONY: help test clean build

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

#
# ──────────────────────────────────────────────────────── I ──────────
#   :::::: C O M M A N D S : :  :   :    :     :        :          :
# ──────────────────────────────────────────────────────────────────
#

version: ## Output the current version
	@echo $(VERSION)

commit: ## Output the current commit
	@echo $(COMMIT)

branch: ## Output the current branch
	@echo $(BRANCH)

build: ## Build the repository - nothing to build
	@-echo "This repo is a Library and is not buildable."

test:  ## Run the unit tests
	@mkdir -p $(REPORTS_PATH)
	@docker run \
		--rm \
		-v "${CURDIR}:/usr/src/app" \
		-w "/usr/src/app" \
		golang:latest \
		go test ./... -v -coverprofile=$(COVERAGE_REPORT)
	

clean: ## Clean up the repoisoty
	@echo "Nothing to clean"

analyze: ## analyze the project
	$(MAKE) .sonar-scanner
#
# ────────────────────────────────────────────────── II ──────────
#   :::::: U T I L S : :  :   :    :     :        :          :
# ────────────────────────────────────────────────────────────
#

#
# ─── SONAR-SCANNER ──────────────────────────────────────────────────────────────
#

SONAR_TOKEN ?= "SonarScannerToken"
SONAR_HOST_URL := "http://127.0.0.1:9000"
SONAR_PROJECT_KEY ?= $(PROJECT_KEY)_$(BRANCH)

SONAR_WDR ?= /usr/src
SONAR_PROJECT_KEY := $(subst $(subst ,, ),:,$(subst /,--,$(SONAR_PROJECT_KEY)))
.sonar-scanner:
	@docker run \
		--rm \
		-e SONAR_HOST_URL="$(SONAR_HOST_URL)" \
		-e SONAR_LOGIN="$(SONAR_TOKEN)" \
		-e SONAR_PROJECTKEY="$(SONAR_PROJECT_KEY)" \
		-v "${CURDIR}:$(SONAR_WDR)" \
		sonarsource/sonar-scanner-cli \
		-D"sonar.projectKey=$(SONAR_PROJECT_KEY)" \
		-D"sonar.go.coverage.reportPaths=$(COVERAGE_REPORT)"
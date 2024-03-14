function printLine() {
	printf "$1"
	printf "\n"
}

REPO_NAME="sandbox-go-webserver"

printLine "ℹ️  Initializing $REPO_NAME..."

printLine "1️⃣  Configure repo to use the source-controlled scripts in the .githooks/ directory..."
git config core.hooksPath ./.githooks/

printLine "2️⃣  Install dependencies..."
go mod tidy

printLine "3️⃣  Generate the initial build..."
make build

printLine "✅ $REPO_NAME has been initialized."
printLine 'ℹ️  Run either `make dev` or `./main` to start the server.'

echo "Building for Linux (amd64)..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
if [ $? -eq 0 ]; then
    echo "Build successful! Binary created at: ./main"
else
    echo "Build failed!"
    exit 1
fi
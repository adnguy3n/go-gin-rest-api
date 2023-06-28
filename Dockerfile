FROM golang:1.20.5

# Set the current working directory.
WORKDIR /usr/src/app

# Get Air for hot reloads.
RUN go install github.com/cosmtrek/air@latest

# Copy the source from the current directory to the workspace.
COPY . .

# Run go mod tidy to download any missing dependencies and remove unused dependencies.
#RUN go mod tidy

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable
CMD ["air", "src/main.go -b 0.0.0.0:8080"]
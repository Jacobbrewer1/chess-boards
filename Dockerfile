FROM golang:1.20

# Create and change to the app directory.
WORKDIR /chess-boards

# Copy local code into docker image
COPY . .

# Runnign the executable
CMD ["./chess.exe"]
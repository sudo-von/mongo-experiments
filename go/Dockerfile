FROM golang:1.15

# Copy files.
COPY . /mongo-experiments/
WORKDIR /mongo-experiments/cmd/
# Make binary.
RUN go build -o $pwd main
# Run binary.
CMD ./main
# Expose port.
EXPOSE 3000
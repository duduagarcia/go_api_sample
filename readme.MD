
docker build --platform="linux/arm64" -t csw24-grupog-ticket-gin-we .

Run

docker run -d -p 8080:8080 --name ticket-gin-container csw24-grupog-ticket-gin-we

 

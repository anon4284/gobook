# A guestbook sample written in go and dart


# Running the this app
- Install docker
- copy ca-certificates from /etc/ssl/certs/ca-certificates to the root dir
- run docker build -t gobook .
- run docker run -d -p 7016:7016 gobook
- open your browser on localhost:7016

# Frameworks used

- Bootstrap for styling
- Angular2-dart frontend framework
- jade for less verbose html
- aws-sdk-go for database saving
- negroni, gorilla router for better http request handling
- govalidator for email validation

# Structure

- lib folder contains frontend source code
- server folder contains backend source code

# URL:
- http://goadstack.com:7016

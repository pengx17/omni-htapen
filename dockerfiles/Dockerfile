FROM ubuntu

COPY ./dist ./omnihtapen 
WORKDIR /omnihtapen

EXPOSE 8080
ENTRYPOINT [ "./server", "--port", "8080", "--secure" ]

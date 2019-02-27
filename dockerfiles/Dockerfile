FROM ubuntu

COPY ./dist ./omnihtapen 
WORKDIR /omnihtapen

EXPOSE 80
ENTRYPOINT [ "./server", "--port", "80", "--secure" ]

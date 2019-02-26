FROM ubuntu

RUN mkdir omnihtapen
COPY ./server ./omnihtapen 
WORKDIR /omnihtapen

EXPOSE 8080
ENTRYPOINT [ "./server", "--port", "8080", "--secure" ]

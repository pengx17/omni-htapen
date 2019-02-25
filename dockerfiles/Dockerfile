FROM ubuntu

COPY ./server ./ominihtapen 
WORKDIR /ominihtapen

EXPOSE 8080
ENTRYPOINT [ "./server", "--port", "8080", "--assetPath", "./" ]

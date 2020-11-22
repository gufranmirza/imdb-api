FROM mongo:latest

WORKDIR /data
COPY . .
RUN ["chmod", "+x", "init.sh"]
CMD ./init.sh
FROM alpine:3.7
COPY app .
EXPOSE 8000
CMD [ "./app" ]
FROM alpine:3.7
COPY workspace/app .
EXPOSE 8000
CMD [ "./app" ]
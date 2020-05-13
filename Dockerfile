FROM alpine:3.7
RUN ls
COPY app .
EXPOSE 8000
CMD [ "./app" ]
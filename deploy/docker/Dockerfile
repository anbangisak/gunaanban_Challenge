FROM golang:latest
WORKDIR /app
EXPOSE 8200
EXPOSE 8080
#RUN git clone https://github.com/anbangisak/gunaanban_Challenge.git;
COPY . /app/.
RUN cd ./gunaanban_Challenge; go build
CMD [ "./gunaanban_Challenge", "-e", "development" ]
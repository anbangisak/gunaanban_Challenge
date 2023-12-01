FROM golang:latest

WORKDIR /app
RUN git clone https://github.com/anbangisak/gunaanban_Challenge.git;
RUN cd ./gunaanban_Challenge; go build
CMD [ "./gunaanban_Challenge", "-e", "development" ]
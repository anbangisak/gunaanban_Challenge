FROM golang:latest

WORKDIR /app
RUN git clone https://github.com/anbangisak/gunaanban_Challenge.git; ./gunaanban_Challenge/build.sh
CMD [ "./gunaanban_Challenge", "-e", "development" ]
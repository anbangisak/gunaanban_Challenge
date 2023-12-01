FROM golang:latest

WORKDIR /app
RUN git clone https://github.com/anbangisak/gunaanban_Challenge.git;
RUN chmod 777 ./gunaanban_Challenge/build.sh
RUN ./gunaanban_Challenge/build.sh
CMD [ "./gunaanban_Challenge", "-e", "development" ]
FROM golang:latest
LABEL maintainer="Vinogradov Nikita  <navinogradov_2@edu.hse.ru>"
WORKDIR /app
RUN git clone https://gitlab.com/quark_lts/quark-node app && cd app
RUN go mod download
RUN make
RUN go install
EXPOSE 9999
CMD ["./quark_node"]
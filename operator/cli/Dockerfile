FROM ubuntu

EXPOSE 5001

WORKDIR /app
RUN apt update -y
RUN apt install -y socat python3-pip

ADD . ./
RUN pip install -r requirements.txt
ENTRYPOINT [ "./entrypoint.sh" ]
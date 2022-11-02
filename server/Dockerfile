FROM python

WORKDIR  /app
EXPOSE 5000
ADD . ./
RUN pip install -r ./requirements.txt

ENTRYPOINT [ "./entrypoint.sh" ]
FROM alpine:3.8
RUN apk add --no-cache py3-pip python3 && \
    pip3 install flask
COPY . /usr/src/app/
ENV FLASK_APP app.py
WORKDIR /usr/src/app
CMD flask run --host=0.0.0.0 --port=6000

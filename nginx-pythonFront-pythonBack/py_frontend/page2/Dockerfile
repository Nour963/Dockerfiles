FROM alpine:3.8
RUN apk add --no-cache py3-pip python3 && \
    pip3 install flask requests Flask-Opentracing jaeger-client
COPY app.py /usr/src/page/
COPY go.html /usr/src/page/templates/
ENV FLASK_APP app.py
WORKDIR /usr/src/page
CMD flask run --host=0.0.0.0 --port=5000
